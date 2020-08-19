package vultr

import (
	"github.com/safehousetech/gocloud/baremetal/vultrbaremetal"
	"github.com/safehousetech/gocloud/compute/vultrcompute"
	"github.com/safehousetech/gocloud/container/vultrcontainer"
	"github.com/safehousetech/gocloud/database/vultrnosql"
	"github.com/safehousetech/gocloud/dns/vultrdns"
	"github.com/safehousetech/gocloud/gocloudinterface"
	"github.com/safehousetech/gocloud/loadbalancer/vultrloadbalancer"
	"github.com/safehousetech/gocloud/machinelearning/vultrmachinelearning"
	"github.com/safehousetech/gocloud/serverless/vultrserverless"
	"github.com/safehousetech/gocloud/storage/vultrstorage"
)

// Vultr struct represents Vultr cloud provider.
type Vultr struct {
	vultrcompute.VultrCompute
	vultrstorage.VultrStorage
	vultrloadbalancer.VultrLoadBalancer
	vultrcontainer.VultrContainer
	vultrdns.VultrDNS
	vultrserverless.Vultrserverless
	vultrnosql.Vultrnosql
	vultrmachinelearning.Vultrmachinelearning
}

func (*Vultr) Compute() gocloudinterface.Compute {
	return &vultrcompute.VultrCompute{}
}

func (*Vultr) Storage() gocloudinterface.Storage {
	return &vultrstorage.VultrStorage{}
}

func (*Vultr) DNS() gocloudinterface.DNS {
	return &vultrdns.VultrDNS{}
}

func (*Vultr) BareMetal() gocloudinterface.BareMetal {
	return &vultrbaremetal.VultrBareMetal{}
}
