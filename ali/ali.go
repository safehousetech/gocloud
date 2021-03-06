package ali

import (
	"github.com/safehousetech/gocloud/compute/ecs"
	"github.com/safehousetech/gocloud/container/alicontainer"
	"github.com/safehousetech/gocloud/database/alinosql"
	"github.com/safehousetech/gocloud/dns/alidns"
	"github.com/safehousetech/gocloud/gocloudinterface"
	"github.com/safehousetech/gocloud/loadbalancer/aliloadbalancer"
	alimachinelearning "github.com/safehousetech/gocloud/machinelearning/alimachinelearning"
	"github.com/safehousetech/gocloud/serverless/aliserverless"
	"github.com/safehousetech/gocloud/storage/alistorage"
)

//Ali struct represents Ali-cloud provider.
type Ali struct {
	ecs.ECS
	alistorage.Alistorage
	aliloadbalancer.Aliloadbalancer
	alicontainer.Alicontainer
	alidns.Alidns
	aliserverless.Aliserverless
	alinosql.Alinosql
	alimachinelearning.Alimachinelearning
}

//Compute .
func (*Ali) Compute() gocloudinterface.Compute {
	return &ecs.ECS{}
}

//Storage .
func (*Ali) Storage() gocloudinterface.Storage {
	return &alistorage.Alistorage{}
}

//LoadBalancer .
func (*Ali) LoadBalancer() gocloudinterface.LoadBalancer {
	return &aliloadbalancer.Aliloadbalancer{}
}

//Container .
func (*Ali) Container() gocloudinterface.Container {
	return &alicontainer.Alicontainer{}
}

//DNS .
func (*Ali) DNS() gocloudinterface.DNS {
	return &alidns.Alidns{}
}
