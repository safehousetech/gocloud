package digiocean

import (
	droplet "github.com/safehousetech/gocloud/compute/droplet"
	digioceancontainer "github.com/safehousetech/gocloud/container/digioceancontainer"
	digioceannosql "github.com/safehousetech/gocloud/database/digioceannosql"
	digioceandns "github.com/safehousetech/gocloud/dns/digioceandns"
	"github.com/safehousetech/gocloud/gocloudinterface"
	digioceanloadbalancer "github.com/safehousetech/gocloud/loadbalancer/digioceanloadbalancer"
	digioceanmachinelearning "github.com/safehousetech/gocloud/machinelearning/digioceanmachinelearning"
	digioceanserverless "github.com/safehousetech/gocloud/serverless/digioceanserverless"
	digioceanstorage "github.com/safehousetech/gocloud/storage/digioceanstorage"
)

// DigitalOcean struct represents DigitalOcean cloud provider.
type DigitalOcean struct {
	droplet.Droplet
	digioceandns.Digioceandns
	digioceanloadbalancer.DigioceanLoadBalancer
	digioceanstorage.Digioceanstorage
	digioceancontainer.Digioceancontainer
	digioceanserverless.Digioceanserverless
	digioceannosql.Digioceannosql
	digioceanmachinelearning.Digioceanmachinelearning
}

func (*DigitalOcean) Compute() gocloudinterface.Compute {
	return &droplet.Droplet{}
}

func (*DigitalOcean) Storage() gocloudinterface.Storage {
	return &digioceanstorage.Digioceanstorage{}
}

func (*DigitalOcean) LoadBalancer() gocloudinterface.LoadBalancer {
	return &digioceanloadbalancer.DigioceanLoadBalancer{}
}

func (*DigitalOcean) DNS() gocloudinterface.DNS {
	return &digioceandns.Digioceandns{}
}
