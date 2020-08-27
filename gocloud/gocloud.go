package gocloud

import (
	"fmt"

	"github.com/safehousetech/gocloud/ali"
	aliAuth "github.com/safehousetech/gocloud/aliauth"
	awsAuth "github.com/safehousetech/gocloud/auth"
	"github.com/safehousetech/gocloud/aws"
	"github.com/safehousetech/gocloud/azure"
	"github.com/safehousetech/gocloud/digiocean"
	digioceanAuth "github.com/safehousetech/gocloud/digioceanauth"
	"github.com/safehousetech/gocloud/google"
	googleAuth "github.com/safehousetech/gocloud/googleauth"
	"github.com/safehousetech/gocloud/openstack"
	"github.com/safehousetech/gocloud/rackspace"
	"github.com/safehousetech/gocloud/rackspaceauth"
	"github.com/safehousetech/gocloud/vultr"
	"github.com/safehousetech/gocloud/vultrauth"
)

// Gocloud is a interface which hides the difference between different cloud providers.
type Gocloud interface {
	CreateNode(request interface{}) (resp interface{}, err error)
	StartNode(request interface{}) (resp interface{}, err error)
	StopNode(request interface{}) (resp interface{}, err error)
	DeleteNode(request interface{}) (resp interface{}, err error)
	RebootNode(request interface{}) (resp interface{}, err error)
	CreateDisk(request interface{}) (resp interface{}, err error)
	DeleteDisk(request interface{}) (resp interface{}, err error)
	CreateSnapshot(request interface{}) (resp interface{}, err error)
	DeleteSnapshot(request interface{}) (resp interface{}, err error)
	AttachDisk(request interface{}) (resp interface{}, err error)
	DetachDisk(request interface{}) (resp interface{}, err error)
	CreateLoadBalancer(request interface{}) (resp interface{}, err error)
	DeleteLoadBalancer(request interface{}) (resp interface{}, err error)
	ListLoadBalancer(request interface{}) (resp interface{}, err error)
	AttachNodeWithLoadBalancer(request interface{}) (resp interface{}, err error)
	DetachNodeWithLoadBalancer(request interface{}) (resp interface{}, err error)
	CreateCluster(request interface{}) (resp interface{}, err error)
	DeleteCluster(request interface{}) (resp interface{}, err error)
	CreateService(request interface{}) (resp interface{}, err error)
	RunTask(request interface{}) (resp interface{}, err error)
	DeleteService(request interface{}) (resp interface{}, err error)
	StopTask(request interface{}) (resp interface{}, err error)
	StartTask(request interface{}) (resp interface{}, err error)
	ListDNS(request interface{}) (resp interface{}, err error)
	CreateDNS(request interface{}) (resp interface{}, err error)
	DeleteDNS(request interface{}) (resp interface{}, err error)
	ListResourceDNSRecordSets(request interface{}) (resp interface{}, err error)
	GetFunction(request interface{}) (resp interface{}, err error)
	CreateFunction(request interface{}) (resp interface{}, err error)
	CallFunction(request interface{}) (resp interface{}, err error)
	ListFunction(request interface{}) (resp interface{}, err error)
	DeleteFunction(request interface{}) (resp interface{}, err error)
	ListTables(request interface{}) (resp interface{}, err error)
	DeleteTables(request interface{}) (resp interface{}, err error)
	DescribeTables(request interface{}) (resp interface{}, err error)
	CreateTables(request interface{}) (resp interface{}, err error)
}

const (
	// Amazonprovider represents Amazon cloud.
	Amazonprovider = "aws"

	// Googleprovider represents Google cloud.
	Googleprovider = "google"

	// Openstackprovider represents Openstack cloud.
	Openstackprovider = "openstack"

	// Azureprovider represents Openstack cloud.
	Azureprovider = "azure"

	// Digioceanprovider represents Digital Ocean cloud.
	Digioceanprovider = "digiocean"

	// Aliprovider reperents Google cloud.
	Aliprovider = "ali"

	// Rackspaceprovider reperents rackspace cloud.
	Rackspaceprovider = "rackspace"

	// Vultrprovider reperents rackspace cloud.
	Vultrprovider = "vultr"
)

// CloudProvider returns the instance of respective cloud and maps it to
// Gocloud so that we can call the method like CreateNode on CloudProvider instance.
// This is a delegation of CloudProvider.
func CloudProvider(provider string) (Gocloud, error) {

	switch provider {
	case Amazonprovider:
		// Calls authentication procedure for AWS
		awsAuth.LoadConfig()
		return new(aws.AWS), nil

	case Googleprovider:
		googleAuth.LoadConfig()
		return new(google.Google), nil

	case Openstackprovider:
		return new(openstack.Openstack), nil

	case Digioceanprovider:
		// Calls authentication procedure for Digital Ocean.
		digioceanAuth.LoadConfig()
		return new(digiocean.DigitalOcean), nil

	case Azureprovider:
		return new(azure.Azure), nil

	case Aliprovider:
		aliAuth.LoadConfig()
		return new(ali.Ali), nil

	case Rackspaceprovider:
		rackspaceauth.LoadConfigAndAuthenticate()
		return new(rackspace.Rackspace), nil

	case Vultrprovider:
		vultrauth.LoadConfig()
		return new(vultr.Vultr), nil

	default:
		return nil, fmt.Errorf("provider %s not recognized", provider)
	}
}
