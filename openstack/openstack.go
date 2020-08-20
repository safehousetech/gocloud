package openstack

import (
	nova "github.com/safehousetech/gocloud/compute/nova"
	magnum "github.com/safehousetech/gocloud/container/magnum"
	openstacknosql "github.com/safehousetech/gocloud/database/openstacknosql"
	designate "github.com/safehousetech/gocloud/dns/designate"
	neutron "github.com/safehousetech/gocloud/loadbalancer/neutron"
	openstackmachinelearning "github.com/safehousetech/gocloud/machinelearning/openstackmachinelearning"
	openstackserverless "github.com/safehousetech/gocloud/serverless/openstackserverless"
	cinder "github.com/safehousetech/gocloud/storage/cinder"
)

//Openstack  struct represents openstack cloud provider.
type Openstack struct {
	nova.Nova
	cinder.Cinder
	designate.Designate
	magnum.Magnum
	neutron.Neutron
	openstackserverless.Openstackserverless
	openstacknosql.Openstacknosql
	openstackmachinelearning.Openstackmachinelearning
}
