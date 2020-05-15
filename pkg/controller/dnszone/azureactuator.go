package dnszone

import (
	"context"
	"errors"

	"github.com/Azure/azure-sdk-for-go/services/dns/mgmt/2018-05-01/dns"
	hivev1 "github.com/openshift/hive/pkg/apis/hive/v1"
	"github.com/openshift/hive/pkg/azureclient"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
)

type AzureActuator struct {
	logger log.FieldLogger

	azureClient azureclient.Client

	dnsZone *hivev1.DNSZone

	managedZone *dns.Zone
}

type azureClientBuilderType func(secret *corev1.Secret) (azureclient.Client, error)

func NewAzureActuator(
	logger log.FieldLogger,
	secret *corev1.Secret,
	dnsZone *hivev1.DNSZone,
	azureClientBuilder azureClientBuilderType,
) (*AzureActuator, error) {
	azureClient, err := azureClientBuilder(secret)
	if err != nil {
		logger.WithError(err).Error("Error creating AzureClient")
		return nil, err
	}

	azureActuator := &AzureActuator{
		logger:      logger,
		azureClient: azureClient,
		dnsZone:     dnsZone,
	}

	return azureActuator, nil
}

var _ Actuator = &AzureActuator{}

func (a *AzureActuator) Create() error {
	logger := a.logger.WithField("zone", a.dnsZone.Spec.Zone)
	logger.Info("Creating managed zone")

	resourceGroupName := a.dnsZone.Spec.Azure.BaseDomainResourceGroupName

	zone := a.dnsZone.Spec.Zone
	managedZone, err := a.azureClient.CreateOrUpdateZone(context.TODO(), resourceGroupName, zone)
	if err != nil {
		logger.WithError(err).Error("Error creating managed zone")
		return err
	}

	logger.Debug("Managed zone successfully created")
	a.managedZone = &managedZone
	return nil
}

func (a *AzureActuator) Delete() error {
	if a.managedZone == nil {
		return errors.New("managedZone is unpopulated")
	}

	resourceGroupName := a.dnsZone.Spec.Azure.BaseDomainResourceGroupName
	logger := a.logger.WithField("zone", a.dnsZone.Spec.Zone).WithField("zoneName", a.managedZone.Name)
	logger.Info("Deleting managed zone")
	err := a.azureClient.DeleteZone(context.TODO(), resourceGroupName, *a.managedZone.Name)
	if err != nil {
		log.WithError(err).Log(log.ErrorLevel, "Cannot delete managed zone")
	}

	return err
}

func (a *AzureActuator) Exists() (bool, error) {
	return a.managedZone != nil, nil
}

func (a *AzureActuator) GetNameServers() ([]string, error) {
	if a.managedZone == nil {
		return nil, errors.New("managedZone is unpopulated")
	}

	logger := a.logger.WithField("zone", a.dnsZone.Spec.Zone)
	result := a.managedZone.NameServers
	logger.WithField("nameservers", result).Debug("found managed zone name servers")
	return *result, nil
}

func (a *AzureActuator) ModifyStatus() error {
	if a.managedZone == nil {
		return errors.New("managedZone is unpopulated")
	}

	a.dnsZone.Status.Azure = &hivev1.AzureDNSZoneStatus{
		ZoneName: a.managedZone.Name,
	}

	return nil
}

func (a *AzureActuator) Refresh() error {
	var zoneName string
	if a.dnsZone.Status.Azure != nil && a.dnsZone.Status.Azure.ZoneName != nil {
		a.logger.Debug("ZoneName is set in status, will retrieve by that name")
		zoneName = *a.dnsZone.Status.Azure.ZoneName
	}

	resourceGroupName := a.dnsZone.Spec.Azure.BaseDomainResourceGroupName

	// Fetch the managed zone
	logger := a.logger.WithField("zoneName", zoneName)
	logger.Debug("Fetching managed zone by zone name")
	resp, err := a.azureClient.GetZone(context.TODO(), resourceGroupName, zoneName)
	if err != nil {
		logger.WithError(err).Error("Cannot get managed zone")
		return err
	}

	logger.Debug("Found managed zone")
	a.managedZone = &resp
	return nil
}

// UpdateMetadata implements the UpdateMetadata call of the actuator interface
func (a *AzureActuator) UpdateMetadata() error {
	return nil
}
