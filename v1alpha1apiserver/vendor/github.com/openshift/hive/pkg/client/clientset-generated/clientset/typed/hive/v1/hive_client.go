// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/hive/pkg/apis/hive/v1"
	"github.com/openshift/hive/pkg/client/clientset-generated/clientset/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type HiveV1Interface interface {
	RESTClient() rest.Interface
	CheckpointsGetter
	ClusterDeploymentsGetter
	ClusterDeprovisionsGetter
	ClusterImageSetsGetter
	ClusterProvisionsGetter
	ClusterStatesGetter
	DNSZonesGetter
	HiveConfigsGetter
	MachinePoolsGetter
	MachinePoolNameLeasesGetter
	SelectorSyncIdentityProvidersGetter
	SelectorSyncSetsGetter
	SyncIdentityProvidersGetter
	SyncSetsGetter
	SyncSetInstancesGetter
}

// HiveV1Client is used to interact with features provided by the hive.openshift.io group.
type HiveV1Client struct {
	restClient rest.Interface
}

func (c *HiveV1Client) Checkpoints(namespace string) CheckpointInterface {
	return newCheckpoints(c, namespace)
}

func (c *HiveV1Client) ClusterDeployments(namespace string) ClusterDeploymentInterface {
	return newClusterDeployments(c, namespace)
}

func (c *HiveV1Client) ClusterDeprovisions(namespace string) ClusterDeprovisionInterface {
	return newClusterDeprovisions(c, namespace)
}

func (c *HiveV1Client) ClusterImageSets() ClusterImageSetInterface {
	return newClusterImageSets(c)
}

func (c *HiveV1Client) ClusterProvisions(namespace string) ClusterProvisionInterface {
	return newClusterProvisions(c, namespace)
}

func (c *HiveV1Client) ClusterStates(namespace string) ClusterStateInterface {
	return newClusterStates(c, namespace)
}

func (c *HiveV1Client) DNSZones(namespace string) DNSZoneInterface {
	return newDNSZones(c, namespace)
}

func (c *HiveV1Client) HiveConfigs() HiveConfigInterface {
	return newHiveConfigs(c)
}

func (c *HiveV1Client) MachinePools(namespace string) MachinePoolInterface {
	return newMachinePools(c, namespace)
}

func (c *HiveV1Client) MachinePoolNameLeases(namespace string) MachinePoolNameLeaseInterface {
	return newMachinePoolNameLeases(c, namespace)
}

func (c *HiveV1Client) SelectorSyncIdentityProviders() SelectorSyncIdentityProviderInterface {
	return newSelectorSyncIdentityProviders(c)
}

func (c *HiveV1Client) SelectorSyncSets() SelectorSyncSetInterface {
	return newSelectorSyncSets(c)
}

func (c *HiveV1Client) SyncIdentityProviders(namespace string) SyncIdentityProviderInterface {
	return newSyncIdentityProviders(c, namespace)
}

func (c *HiveV1Client) SyncSets(namespace string) SyncSetInterface {
	return newSyncSets(c, namespace)
}

func (c *HiveV1Client) SyncSetInstances(namespace string) SyncSetInstanceInterface {
	return newSyncSetInstances(c, namespace)
}

// NewForConfig creates a new HiveV1Client for the given config.
func NewForConfig(c *rest.Config) (*HiveV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &HiveV1Client{client}, nil
}

// NewForConfigOrDie creates a new HiveV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *HiveV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new HiveV1Client for the given RESTClient.
func New(c rest.Interface) *HiveV1Client {
	return &HiveV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *HiveV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
