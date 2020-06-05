package nameserver

import (
	"context"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/dns/mgmt/2018-05-01/dns"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/openshift/hive/pkg/azureclient"
	controllerutils "github.com/openshift/hive/pkg/controller/utils"
)

const (
	defaultCallTimeout = 30 * time.Second
)

func contextWithTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, defaultCallTimeout)
}

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
	return currentNameServers, errors.Wrap(err, "error querying name servers")
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
	ctx, cancel := contextWithTimeout(context.TODO())
	defer cancel()

	return azureClient.DeleteRecordSet(ctx, q.resourceGroupName, rootDomain, q.getRelativeDomain(rootDomain, domain), dns.NS)
}

// createNameServers creates the name servers for the specified domain in the specified managed zone.
func (q *azureQuery) createNameServers(azureClient azureclient.Client, rootDomain string, domain string, values sets.String) error {
	ctx, cancel := contextWithTimeout(context.TODO())
	defer cancel()

	_, err := azureClient.CreateOrUpdateRecordSet(ctx, q.resourceGroupName, rootDomain, q.getRelativeDomain(rootDomain, domain), dns.NS, q.recordSet(values))

	return errors.Wrap(err, "something went wrong when creating name servers")
}

func (q *azureQuery) recordSet(values sets.String) dns.RecordSet {
	nsRecords := make([]dns.NsRecord, len(values))

	for i, v := range values.List() {
		nsRecords[i] = dns.NsRecord{
			Nsdname: to.StringPtr(v),
		}
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
	ctx, cancel := contextWithTimeout(context.TODO())
	defer cancel()

	relativeDomain := q.getRelativeDomain(rootDomain, domain)

	recordSetsPage, err := azureClient.ListRecordSetsByZone(ctx, q.resourceGroupName, rootDomain, relativeDomain)
	if err != nil {
		return nil, err
	}

	for recordSetsPage.NotDone() {
		for _, recordSet := range recordSetsPage.Values() {
			if recordSet.RecordSetProperties.NsRecords != nil {
				continue
			}
			if *recordSet.Name == relativeDomain {
				values := sets.NewString()
				for _, v := range *recordSet.NsRecords {
					values.Insert(*v.Nsdname)
				}
				return values, nil
			}
		}
		if err := recordSetsPage.NextWithContext(ctx); err != nil {
			return nil, err
		}
	}

	// No records found
	return nil, nil
}

// queryNameServers queries Azure for the name servers in the specified managed zone.
func (q *azureQuery) queryNameServers(azureClient azureclient.Client, rootDomain string) (map[string]sets.String, error) {
	ctx, cancel := contextWithTimeout(context.TODO())
	defer cancel()

	nameServers := map[string]sets.String{}
	recordSetsPage, err := azureClient.ListRecordSetsByZone(ctx, q.resourceGroupName, rootDomain, "")
	if err != nil {
		return nil, err
	}

	for recordSetsPage.NotDone() {
		for _, recordSet := range recordSetsPage.Values() {
			if recordSet.RecordSetProperties.NsRecords == nil {
				continue
			}
			values := sets.NewString()
			for _, v := range *recordSet.NsRecords {
				values.Insert(*v.Nsdname)
			}
			if *recordSet.Name == "@" {
				nameServers[rootDomain] = values
			} else {
				nameServers[q.getFQDN(*recordSet.Name, rootDomain)] = values
			}
		}

		if err := recordSetsPage.NextWithContext(ctx); err != nil {
			return nil, err
		}
	}

	return nameServers, nil
}

func (q *azureQuery) getRelativeDomain(rootDomain string, domain string) string {
	return controllerutils.Undotted(strings.TrimSuffix(domain, rootDomain))
}

func (q *azureQuery) getFQDN(relativeDomain string, rootDomain string) string {
	return relativeDomain + "." + rootDomain
}
