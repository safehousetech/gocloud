package ecs

import (
	"errors"

	"github.com/safehousetech/gocloud/aliauth"
)

// CreateNode to store all attribute to create Ali-cloud ECS instance
type CreateNode struct {
	RegionID                string
	ZoneID                  string
	ImageID                 string
	InstanceType            string
	SecurityGroupID         string
	InstanceName            string
	Description             string
	InternetChargeType      string
	InternetMaxBandwidthIn  int
	InternetMaxBandwidthOut int
	HostName                string
	Password                string
	IoOptimized             string
	SystemDiskCategory      string
	SystemDiskSize          string
	SystemDiskName          string
	SystemDiskDescription   string
}

// StartNode to store all attribute to start Ali-cloud ECS instance
type StartNode struct {
	InstanceID    string
	InitLocalDisk bool
}

// StopNode to store all attribute to Stop Ali-cloud ECS instance
type StopNode struct {
	InstanceID  string
	ForceStop   bool
	ConfirmStop bool
	StoppedMode string
}

// RebootNode to store all attribute to Reboot Ali-cloud ECS instance
type RebootNode struct {
	InstanceID string
	ForceStop  bool
}

// DeleteNode to store all attribute to Delete Ali-cloud ECS instance
type DeleteNode struct {
	InstanceID string
}

// CreateNodeBuilder pattern code
type CreateNodeBuilder struct {
	createNode *CreateNode
}

//NewCreateNodeBuilder ...
func NewCreateNodeBuilder() *CreateNodeBuilder {
	createNode := &CreateNode{}
	b := &CreateNodeBuilder{createNode: createNode}
	return b
}

//RegionID ...
func (b *CreateNodeBuilder) RegionID(regionID string) *CreateNodeBuilder {
	b.createNode.RegionID = regionID
	return b
}

//ZoneID ...
func (b *CreateNodeBuilder) ZoneID(zoneID string) *CreateNodeBuilder {
	b.createNode.ZoneID = zoneID
	return b
}

//ImageID ...
func (b *CreateNodeBuilder) ImageID(imageID string) *CreateNodeBuilder {
	b.createNode.ImageID = imageID
	return b
}

//InstanceType ...
func (b *CreateNodeBuilder) InstanceType(instanceType string) *CreateNodeBuilder {
	b.createNode.InstanceType = instanceType
	return b
}

//SecurityGroupID ...
func (b *CreateNodeBuilder) SecurityGroupID(securityGroupID string) *CreateNodeBuilder {
	b.createNode.SecurityGroupID = securityGroupID
	return b
}

//InstanceName ...
func (b *CreateNodeBuilder) InstanceName(instanceName string) *CreateNodeBuilder {
	b.createNode.InstanceName = instanceName
	return b
}

//Description ...
func (b *CreateNodeBuilder) Description(description string) *CreateNodeBuilder {
	b.createNode.Description = description
	return b
}

//InternetChargeType ...
func (b *CreateNodeBuilder) InternetChargeType(internetChargeType string) *CreateNodeBuilder {
	b.createNode.InternetChargeType = internetChargeType
	return b
}

//InternetMaxBandwidthIn ...
func (b *CreateNodeBuilder) InternetMaxBandwidthIn(internetMaxBandwidthIn int) *CreateNodeBuilder {
	b.createNode.InternetMaxBandwidthIn = internetMaxBandwidthIn
	return b
}

//InternetMaxBandwidthOut ...
func (b *CreateNodeBuilder) InternetMaxBandwidthOut(internetMaxBandwidthOut int) *CreateNodeBuilder {
	b.createNode.InternetMaxBandwidthOut = internetMaxBandwidthOut
	return b
}

//HostName ...
func (b *CreateNodeBuilder) HostName(hostName string) *CreateNodeBuilder {
	b.createNode.HostName = hostName
	return b
}

//Password ...
func (b *CreateNodeBuilder) Password(password string) *CreateNodeBuilder {
	b.createNode.Password = password
	return b
}

//IoOptimized ...
func (b *CreateNodeBuilder) IoOptimized(ioOptimized string) *CreateNodeBuilder {
	b.createNode.IoOptimized = ioOptimized
	return b
}

//SystemDiskCategory ...
func (b *CreateNodeBuilder) SystemDiskCategory(systemDiskCategory string) *CreateNodeBuilder {
	b.createNode.SystemDiskCategory = systemDiskCategory
	return b
}

//SystemDiskSize ...
func (b *CreateNodeBuilder) SystemDiskSize(systemDiskSize string) *CreateNodeBuilder {
	b.createNode.SystemDiskSize = systemDiskSize
	return b
}

//SystemDiskName ..
func (b *CreateNodeBuilder) SystemDiskName(systemDiskName string) *CreateNodeBuilder {
	b.createNode.SystemDiskName = systemDiskName
	return b
}

//SystemDiskDescription ...
func (b *CreateNodeBuilder) SystemDiskDescription(systemDiskDescription string) *CreateNodeBuilder {
	b.createNode.SystemDiskDescription = systemDiskDescription
	return b
}

//Build ...
func (b *CreateNodeBuilder) Build() (map[string]interface{}, error) {
	if b.createNode.RegionID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "RegionID")
	}
	if b.createNode.ImageID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "ImageID")
	}
	if b.createNode.InstanceType == "" {
		return nil, errors.New(aliauth.StrMissRequired + "InstanceType")
	}
	if b.createNode.SecurityGroupID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "SecurityGroupID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.createNode)
	return params, nil
}

// StartNodeBuilder pattern code
type StartNodeBuilder struct {
	startNode *StartNode
}

//NewStartNodeBuilder ...
func NewStartNodeBuilder() *StartNodeBuilder {
	startNode := &StartNode{}
	b := &StartNodeBuilder{startNode: startNode}
	return b
}

//InstanceID ...
func (b *StartNodeBuilder) InstanceID(instanceID string) *StartNodeBuilder {
	b.startNode.InstanceID = instanceID
	return b
}

//InitLocalDisk ...
func (b *StartNodeBuilder) InitLocalDisk(initLocalDisk bool) *StartNodeBuilder {
	b.startNode.InitLocalDisk = initLocalDisk
	return b
}

//Build ...
func (b *StartNodeBuilder) Build() (map[string]interface{}, error) {
	if b.startNode.InstanceID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "InstanceID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.startNode)
	return params, nil
}

// StopNodeBuilder pattern code
type StopNodeBuilder struct {
	stopNode *StopNode
}

//NewStopNodeBuilder ...
func NewStopNodeBuilder() *StopNodeBuilder {
	stopNode := &StopNode{}
	b := &StopNodeBuilder{stopNode: stopNode}
	return b
}

//InstanceID ...
func (b *StopNodeBuilder) InstanceID(instanceID string) *StopNodeBuilder {
	b.stopNode.InstanceID = instanceID
	return b
}

//ForceStop ...
func (b *StopNodeBuilder) ForceStop(forceStop bool) *StopNodeBuilder {
	b.stopNode.ForceStop = forceStop
	return b
}

//ConfirmStop ...
func (b *StopNodeBuilder) ConfirmStop(confirmStop bool) *StopNodeBuilder {
	b.stopNode.ConfirmStop = confirmStop
	return b
}

//StoppedMode ...
func (b *StopNodeBuilder) StoppedMode(stoppedMode string) *StopNodeBuilder {
	b.stopNode.StoppedMode = stoppedMode
	return b
}

//Build ...
func (b *StopNodeBuilder) Build() (map[string]interface{}, error) {
	if b.stopNode.InstanceID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "InstanceID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.stopNode)
	return params, nil
}

// RebootNodeBuilder pattern code
type RebootNodeBuilder struct {
	rebootNode *RebootNode
}

//NewRebootNodeBuilder ...
func NewRebootNodeBuilder() *RebootNodeBuilder {
	rebootNode := &RebootNode{}
	b := &RebootNodeBuilder{rebootNode: rebootNode}
	return b
}

//InstanceID ...
func (b *RebootNodeBuilder) InstanceID(instanceID string) *RebootNodeBuilder {
	b.rebootNode.InstanceID = instanceID
	return b
}

//ForceStop ...
func (b *RebootNodeBuilder) ForceStop(forceStop bool) *RebootNodeBuilder {
	b.rebootNode.ForceStop = forceStop
	return b
}

//Build ...
func (b *RebootNodeBuilder) Build() (map[string]interface{}, error) {
	if b.rebootNode.InstanceID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "InstanceID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.rebootNode)
	return params, nil
}

// DeleteNodeBuilder pattern code
type DeleteNodeBuilder struct {
	deleteNode *DeleteNode
}

//NewDeleteNodeBuilder ...
func NewDeleteNodeBuilder() *DeleteNodeBuilder {
	deleteNode := &DeleteNode{}
	b := &DeleteNodeBuilder{deleteNode: deleteNode}
	return b
}

// InstanceID ...
func (b *DeleteNodeBuilder) InstanceID(instanceID string) *DeleteNodeBuilder {
	b.deleteNode.InstanceID = instanceID
	return b
}

//Build ...
func (b *DeleteNodeBuilder) Build() (map[string]interface{}, error) {
	if b.deleteNode.InstanceID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "InstanceID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.deleteNode)
	return params, nil
}
