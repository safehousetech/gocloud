package rackspace

import (
	rackspacecompute "github.com/safehousetech/gocloud/compute/rackspacecompute"
	rackspacecontainer "github.com/safehousetech/gocloud/container/rackspacecontainer"
	rackspacenosql "github.com/safehousetech/gocloud/database/rackspacenosql"
	rackspacedns "github.com/safehousetech/gocloud/dns/rackspacedns"
	rackspaceloadbalancer "github.com/safehousetech/gocloud/loadbalancer/rackspaceloadbalancer"
	rackspacemachinelearning "github.com/safehousetech/gocloud/machinelearning/rackspacemachinelearning"
	rackspaceserverless "github.com/safehousetech/gocloud/serverless/rackspaceserverless"
	rackspacestorage "github.com/safehousetech/gocloud/storage/rackspacestorage"
)

// Rackspace  struct represents Rackspace cloud provider.
type Rackspace struct {
	rackspacecompute.Rackspacecompute
	rackspacestorage.Rackspacestorage
	rackspaceloadbalancer.Rackspaceloadbalancer
	rackspacecontainer.Rackspacecontainer
	rackspacedns.Rackspacedns
	rackspaceserverless.Rackspaceserverless
	rackspacenosql.Rackspacenosql
	rackspacemachinelearning.Rackspacemachinelearning
}
