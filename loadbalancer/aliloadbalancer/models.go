package aliloadbalancer

import (
	"errors"

	"github.com/safehousetech/gocloud/aliauth"
)

//Aliloadbalancer represents Aliloadbalancer struct.
type Aliloadbalancer struct {
}

// CreateLoadBalancer struct represents attribute of create LoadBalancer.
type CreateLoadBalancer struct {
	RegionID           string
	MasterZoneID       string
	SlaveZoneID        string
	LoadBalancerSpec   string
	LoadBalancerName   string
	AddressType        string
	VSwitchID          string
	PayType            string
	PricingCycle       string
	Duration           string
	Autopay            bool
	InternetChargeType string
	Bandwidth          int
	ClientToken        string
	ResourceGroupID    string
}

// DeleteLoadBalancer struct represents attribute of delete LoadBalancer.
type DeleteLoadBalancer struct {
	RegionID       string
	LoadBalancerID string
}

// ListLoadBalancer struct represents attribute of list LoadBalancer.
type ListLoadBalancer struct {
	RegionID              string
	LoadBalancerID        string
	LoadBalancerName      string
	AddressType           string
	NetworkType           string
	VpcID                 string
	VswitchID             string
	Address               string
	ServerIntranetAddress int
	InternetChargeType    string
	ServerID              string
	MasterZoneID          string
	SlaveZoneID           string
}

// AttachLoadBalancer represents Attach node with loadbalancer attribute.
type AttachLoadBalancer struct {
	LoadBalancerID string
	BackendServers string
}

// DetachLoadBalancer represents Detach node with loadbalancer attribute.
type DetachLoadBalancer struct {
	RegionID       string
	LoadBalancerID string
	BackendServers string
}

// ...
const (
	DefaultRegion = "slb.aliyuncs.com"
	Zhangjiakou   = "slb.cn-zhangjiakou.aliyuncs.com"
	Hohhot        = "slb.cn-huhehaote.aliyuncs.com"
	Tokyo         = "slb.ap-northeast-1.aliyuncs.com"
	Sydney        = "slb.ap-southeast-2.aliyuncs.com"
	KualaLumpur   = "slb.ap-southeast-3.aliyuncs.com"
	Jakarta       = "slb.ap-southeast-5.aliyuncs.com"
	Mumbai        = "slb.ap-south-1.aliyuncs.com"
	Dubai         = "slb.me-east-1.aliyuncs.com"
	Frankfurt     = "slb.eu-central-1.aliyuncs.com"
)

// CreateLoadBalancerBuilder pattern code
type CreateLoadBalancerBuilder struct {
	createLoadBalancer *CreateLoadBalancer
}

//NewCreateLoadBalancerBuilder ...
func NewCreateLoadBalancerBuilder() *CreateLoadBalancerBuilder {
	createLoadBalancer := &CreateLoadBalancer{}
	b := &CreateLoadBalancerBuilder{createLoadBalancer: createLoadBalancer}
	return b
}

//RegionID ...
func (b *CreateLoadBalancerBuilder) RegionID(regionID string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.RegionID = regionID
	return b
}

//MasterZoneID ...
func (b *CreateLoadBalancerBuilder) MasterZoneID(masterZoneID string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.MasterZoneID = masterZoneID
	return b
}

//SlaveZoneID ...
func (b *CreateLoadBalancerBuilder) SlaveZoneID(slaveZoneID string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.SlaveZoneID = slaveZoneID
	return b
}

//LoadBalancerSpec ...
func (b *CreateLoadBalancerBuilder) LoadBalancerSpec(loadBalancerSpec string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.LoadBalancerSpec = loadBalancerSpec
	return b
}

//LoadBalancerName ...
func (b *CreateLoadBalancerBuilder) LoadBalancerName(loadBalancerName string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.LoadBalancerName = loadBalancerName
	return b
}

//AddressType ...
func (b *CreateLoadBalancerBuilder) AddressType(addressType string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.AddressType = addressType
	return b
}

//VSwitchID ...
func (b *CreateLoadBalancerBuilder) VSwitchID(vSwitchID string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.VSwitchID = vSwitchID
	return b
}

//PayType ...
func (b *CreateLoadBalancerBuilder) PayType(payType string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.PayType = payType
	return b
}

//PricingCycle ...
func (b *CreateLoadBalancerBuilder) PricingCycle(pricingCycle string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.PricingCycle = pricingCycle
	return b
}

//Duration ...
func (b *CreateLoadBalancerBuilder) Duration(duration string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.Duration = duration
	return b
}

//Autopay ...
func (b *CreateLoadBalancerBuilder) Autopay(autopay bool) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.Autopay = autopay
	return b
}

//InternetChargeType ...
func (b *CreateLoadBalancerBuilder) InternetChargeType(internetChargeType string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.InternetChargeType = internetChargeType
	return b
}

//Bandwidth ...
func (b *CreateLoadBalancerBuilder) Bandwidth(bandwidth int) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.Bandwidth = bandwidth
	return b
}

//ClientToken ...
func (b *CreateLoadBalancerBuilder) ClientToken(clientToken string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.ClientToken = clientToken
	return b
}

//ResourceGroupID ...
func (b *CreateLoadBalancerBuilder) ResourceGroupID(resourceGroupID string) *CreateLoadBalancerBuilder {
	b.createLoadBalancer.ResourceGroupID = resourceGroupID
	return b
}

//Build ...
func (b *CreateLoadBalancerBuilder) Build() (map[string]interface{}, error) {
	if b.createLoadBalancer.RegionID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "RegionID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.createLoadBalancer)
	return params, nil
}

// DeleteLoadBalancerBuilder pattern code
type DeleteLoadBalancerBuilder struct {
	deleteLoadBalancer *DeleteLoadBalancer
}

//NewDeleteLoadBalancerBuilder ...
func NewDeleteLoadBalancerBuilder() *DeleteLoadBalancerBuilder {
	deleteLoadBalancer := &DeleteLoadBalancer{}
	b := &DeleteLoadBalancerBuilder{deleteLoadBalancer: deleteLoadBalancer}
	return b
}

//RegionID ...
func (b *DeleteLoadBalancerBuilder) RegionID(regionID string) *DeleteLoadBalancerBuilder {
	b.deleteLoadBalancer.RegionID = regionID
	return b
}

//LoadBalancerID ...
func (b *DeleteLoadBalancerBuilder) LoadBalancerID(loadBalancerID string) *DeleteLoadBalancerBuilder {
	b.deleteLoadBalancer.LoadBalancerID = loadBalancerID
	return b
}

//Build ...
func (b *DeleteLoadBalancerBuilder) Build() (map[string]interface{}, error) {
	if b.deleteLoadBalancer.RegionID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "RegionID")
	}
	if b.deleteLoadBalancer.LoadBalancerID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "LoadBalancerID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.deleteLoadBalancer)
	return params, nil
}

// ListLoadBalancerBuilder pattern code
type ListLoadBalancerBuilder struct {
	listLoadBalancer *ListLoadBalancer
}

//NewListLoadBalancerBuilder ...
func NewListLoadBalancerBuilder() *ListLoadBalancerBuilder {
	listLoadBalancer := &ListLoadBalancer{}
	b := &ListLoadBalancerBuilder{listLoadBalancer: listLoadBalancer}
	return b
}

//RegionID ...
func (b *ListLoadBalancerBuilder) RegionID(regionID string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.RegionID = regionID
	return b
}

//LoadBalancerID ...
func (b *ListLoadBalancerBuilder) LoadBalancerID(loadBalancerID string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.LoadBalancerID = loadBalancerID
	return b
}

//LoadBalancerName ...
func (b *ListLoadBalancerBuilder) LoadBalancerName(loadBalancerName string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.LoadBalancerName = loadBalancerName
	return b
}

//AddressType ...
func (b *ListLoadBalancerBuilder) AddressType(addressType string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.AddressType = addressType
	return b
}

//NetworkType ...
func (b *ListLoadBalancerBuilder) NetworkType(networkType string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.NetworkType = networkType
	return b
}

//VpcID ...
func (b *ListLoadBalancerBuilder) VpcID(vpcID string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.VpcID = vpcID
	return b
}

//VswitchID ...
func (b *ListLoadBalancerBuilder) VswitchID(vswitchID string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.VswitchID = vswitchID
	return b
}

//Address ..
func (b *ListLoadBalancerBuilder) Address(address string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.Address = address
	return b
}

//ServerIntranetAddress ...
func (b *ListLoadBalancerBuilder) ServerIntranetAddress(serverIntranetAddress int) *ListLoadBalancerBuilder {
	b.listLoadBalancer.ServerIntranetAddress = serverIntranetAddress
	return b
}

//InternetChargeType ...
func (b *ListLoadBalancerBuilder) InternetChargeType(internetChargeType string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.InternetChargeType = internetChargeType
	return b
}

//ServerID ...
func (b *ListLoadBalancerBuilder) ServerID(serverID string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.ServerID = serverID
	return b
}

//MasterZoneID ...
func (b *ListLoadBalancerBuilder) MasterZoneID(masterZoneID string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.MasterZoneID = masterZoneID
	return b
}

//SlaveZoneID ...
func (b *ListLoadBalancerBuilder) SlaveZoneID(slaveZoneID string) *ListLoadBalancerBuilder {
	b.listLoadBalancer.SlaveZoneID = slaveZoneID
	return b
}

//Build ...
func (b *ListLoadBalancerBuilder) Build() (map[string]interface{}, error) {
	if b.listLoadBalancer.RegionID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "RegionID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.listLoadBalancer)
	return params, nil
}

// AttachLoadBalancerBuilder pattern code
type AttachLoadBalancerBuilder struct {
	attachLoadBalancer *AttachLoadBalancer
}

//NewAttachLoadBalancerBuilder ...
func NewAttachLoadBalancerBuilder() *AttachLoadBalancerBuilder {
	attachLoadBalancer := &AttachLoadBalancer{}
	b := &AttachLoadBalancerBuilder{attachLoadBalancer: attachLoadBalancer}
	return b
}

//LoadBalancerID ...
func (b *AttachLoadBalancerBuilder) LoadBalancerID(loadBalancerID string) *AttachLoadBalancerBuilder {
	b.attachLoadBalancer.LoadBalancerID = loadBalancerID
	return b
}

//BackendServers ...
func (b *AttachLoadBalancerBuilder) BackendServers(backendServers string) *AttachLoadBalancerBuilder {
	b.attachLoadBalancer.BackendServers = backendServers
	return b
}

//Build ..
func (b *AttachLoadBalancerBuilder) Build() (map[string]interface{}, error) {
	if b.attachLoadBalancer.LoadBalancerID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "LoadBalancerID")
	}
	if b.attachLoadBalancer.BackendServers == "" {
		return nil, errors.New(aliauth.StrMissRequired + "BackendServers")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.attachLoadBalancer)
	return params, nil
}

// DetachLoadBalancerBuilder pattern code
type DetachLoadBalancerBuilder struct {
	detachLoadBalancer *DetachLoadBalancer
}

//NewDetachLoadBalancerBuilder ...
func NewDetachLoadBalancerBuilder() *DetachLoadBalancerBuilder {
	detachLoadBalancer := &DetachLoadBalancer{}
	b := &DetachLoadBalancerBuilder{detachLoadBalancer: detachLoadBalancer}
	return b
}

//RegionID ..
func (b *DetachLoadBalancerBuilder) RegionID(regionID string) *DetachLoadBalancerBuilder {
	b.detachLoadBalancer.RegionID = regionID
	return b
}

//LoadBalancerID ...
func (b *DetachLoadBalancerBuilder) LoadBalancerID(loadBalancerID string) *DetachLoadBalancerBuilder {
	b.detachLoadBalancer.LoadBalancerID = loadBalancerID
	return b
}

//BackendServers ...
func (b *DetachLoadBalancerBuilder) BackendServers(backendServers string) *DetachLoadBalancerBuilder {
	b.detachLoadBalancer.BackendServers = backendServers
	return b
}

//Build ..Build
func (b *DetachLoadBalancerBuilder) Build() (map[string]interface{}, error) {
	if b.detachLoadBalancer.RegionID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "RegionID")
	}
	if b.detachLoadBalancer.LoadBalancerID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "LoadBalancerID")
	}
	if b.detachLoadBalancer.BackendServers == "" {
		return nil, errors.New(aliauth.StrMissRequired + "BackendServers")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.detachLoadBalancer)
	return params, nil
}
