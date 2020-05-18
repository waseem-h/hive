package nameserver

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/dns/mgmt/2018-05-01/dns"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/openshift/hive/pkg/azureclient"
	controllerutils "github.com/openshift/hive/pkg/controller/utils"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewAzureQuery creates a new name server query for Azure.
func NewAzureQuery(c client.Client, credsSecretName string, resourceGroupName string) Query {
	return &azureQuery{
		getAzureClient: func() (azureclient.Client, error) {
			credsSecret := &corev1.Secret{}
			if err := c.Get(
				context.Background(),
				client.ObjectKey{Namespace: controllerutils.GetHiveNamespace(), Name: credsSecretName},
				credsSecret,
			); err != nil {
				return nil, errors.Wrap(err, "could not get the creds secret")
			}
			azureClient, err := azureclient.NewClientFromSecret(credsSecret)
			return azureClient, errors.Wrap(err, "error creating Azure client")
		},
		resourceGroupName: resourceGroupName,
	}
}

type azureQuery struct {
	getAzureClient    func() (azureclient.Client, error)
	resourceGroupName string
}

var _ Query = (*azureQuery)(nil)

// Get implements Query.Get.
func (q *azureQuery) Get(domain string) (map[string]sets.String, error) {
	azureClient, err := q.getAzureClient()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get azure client")
	}
	currentNameServers, err := q.queryNameServers(azureClient, domain)
	return currentNameServers, errors.Wrap(err, "error quering name servers")
}

// Create implements Query.Create.
func (q *azureQuery) Create(rootDomain string, domain string, values sets.String) error {
	azureClient, err := q.getAzureClient()
	if err != nil {
		return errors.Wrap(err, "failed to get Azure client")
	}

	return errors.Wrap(q.createNameServers(azureClient, rootDomain, domain, values), "error creating the name server")
}

// Delete implements Query.Delete.
func (q *azureQuery) Delete(rootDomain string, domain string, values sets.String) error {
	azureClient, err := q.getAzureClient()
	if err != nil {
		return errors.Wrap(err, "failed to get Azure client")
	}

	if len(values) != 0 {
		// If values were provided for the name servers, attempt to perform a
		// delete using those values.
		return errors.Wrap(q.deleteNameServers(azureClient, rootDomain, domain, values), "error deleting the name server")
	}

	// Since we do not have up-to-date values for the name servers, we need
	// to query Azure for the current values to use them in the delete.
	values, err = q.queryNameServer(azureClient, rootDomain, domain)
	if err != nil {
		return errors.Wrap(err, "error querying the current values of the name server")
	}
	if len(values) == 0 {
		return nil
	}

	return errors.Wrap(
		q.deleteNameServers(azureClient, rootDomain, domain, values),
		"error deleting the name server with recently read values",
	)
}

// deleteNameServers deletes the name servers for the specified domain in the specified managed zone.
func (q *azureQuery) deleteNameServers(azureClient azureclient.Client, rootDomain string, domain string, values sets.String) error {
	return azureClient.DeleteRecordSet(context.TODO(), q.resourceGroupName, rootDomain, domain, dns.NS)
}

// createNameServers creates the name servers for the specified domain in the specified managed zone.
func (q *azureQuery) createNameServers(azureClient azureclient.Client, rootDomain string, domain string, values sets.String) error {
	_, err := azureClient.CreateOrUpdateRecordSet(context.TODO(), q.resourceGroupName, rootDomain, domain, dns.NS, q.recordSet(domain, values))

	return errors.Wrap(err, "something went wrong when creating name servers")
}

func (q *azureQuery) recordSet(domain string, values sets.String) dns.RecordSet {
	nsRecords := make([]dns.NsRecord, len(values))

	for _, v := range values.List() {
		nsRecords = append(nsRecords, dns.NsRecord{
			Nsdname: to.StringPtr(v),
		})
	}

	return dns.RecordSet{
		RecordSetProperties: &dns.RecordSetProperties{
			NsRecords: &nsRecords,
			TTL:       to.Int64Ptr(60),
		},
	}
}

// queryNameServer queries Azure for the name servers for the specified domain in the specified managed zone.
func (q *azureQuery) queryNameServer(azureClient azureclient.Client, rootDomain string, domain string) (sets.String, error) {
	recordSets, err := azureClient.ListRecordSetsByZone(context.TODO(), q.resourceGroupName, rootDomain)
	if err != nil {
		return nil, err
	}

	for _, recordSet := range *recordSets {
		if *recordSet.Type != string(dns.NS) {
			continue
		}
		if *recordSet.Name == domain {
			values := sets.NewString()
			for _, v := range *recordSet.NsRecords {
				values.Insert(*v.Nsdname)
			}
			return values, nil
		}
	}

	return nil, errors.New("Cannot find name servers for domain")
}

// queryNameServers queries Azure for the name servers in the specified managed zone.
func (q *azureQuery) queryNameServers(azureClient azureclient.Client, rootDomain string) (map[string]sets.String, error) {
	nameServers := map[string]sets.String{}
	recordSets, err := azureClient.ListRecordSetsByZone(context.TODO(), q.resourceGroupName, rootDomain)
	if err != nil {
		return nil, err
	}

	for _, recordSet := range *recordSets {
		if *recordSet.Type != string(dns.NS) {
			continue
		}
		values := sets.NewString()
		for _, v := range *recordSet.NsRecords {
			values.Insert(*v.Nsdname)
		}
		nameServers[*recordSet.Name] = values
	}

	return nameServers, nil
}
