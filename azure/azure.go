package azure

import (
	azurecompute "github.com/safehousetech/gocloud/compute/azurecompute"
	azurecontainer "github.com/safehousetech/gocloud/container/azurecontainer"
	azurenosql "github.com/safehousetech/gocloud/database/azurenosql"
	azuredns "github.com/safehousetech/gocloud/dns/azuredns"
	azureloadbalancer "github.com/safehousetech/gocloud/loadbalancer/azureloadbalancer"
	azuremachinelearning "github.com/safehousetech/gocloud/machinelearning/azuremachinelearning"
	azureserverless "github.com/safehousetech/gocloud/serverless/azureserverless"
	azurestorage "github.com/safehousetech/gocloud/storage/azurestorage"
)

// Azure  struct represents Microsoft Azure cloud provider.
type Azure struct {
	azurecompute.Azurecompute
	azurestorage.Azurestorage
	azureloadbalancer.Azureloadbalancer
	azurecontainer.Azurecontainer
	azuredns.Azuredns
	azureserverless.Azureserverless
	azurenosql.Azurenosql
	azuremachinelearning.Azuremachinelearning
}
