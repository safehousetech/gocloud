```
package aliloadbalancer
    import "github.com/safehousetech/gocloud/loadbalancer/aliloadbalancer"
```

### CONSTANTS

```
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
```

### TYPES

```
type Aliloadbalancer struct {
}
    Aliloadbalancer represents Aliloadbalancer struct.

func (aliloadbalancer *Aliloadbalancer) AttachNodeWithLoadBalancer(request interface{}) (resp interface{}, err error)
    AttachNodeWithLoadBalancer attach ali ecs instance to ali loadbalancer

func (aliloadbalancer *Aliloadbalancer) CreateLoadBalancer(request interface{}) (resp interface{}, err error)
    CreateLoadBalancer creates ali loadbalancer

func (aliloadbalancer *Aliloadbalancer) DeleteLoadBalancer(request interface{}) (resp interface{}, err error)
    DeleteLoadBalancer deletes ali loadbalancer

func (aliloadbalancer *Aliloadbalancer) DetachNodeWithLoadBalancer(request interface{}) (resp interface{}, err error)
    DetachNodeWithLoadBalancer detach ali ecs instance from ali loadbalancer

func (aliloadbalancer *Aliloadbalancer) ListLoadBalancer(request interface{}) (resp interface{}, err error)
    ListLoadBalancer lists ali loadbalancer

type AttachLoadBalancer struct {
    LoadBalancerID string
    BackendServers string
}
    AttachLoadBalancer represents Attach node with loadbalancer attribute.

type AttachLoadBalancerBuilder struct {
    // contains filtered or unexported fields
}
    AttachLoadBalancer builder pattern code

func NewAttachLoadBalancerBuilder() *AttachLoadBalancerBuilder

func (b *AttachLoadBalancerBuilder) BackendServers(backendServers string) *AttachLoadBalancerBuilder

func (b *AttachLoadBalancerBuilder) Build() (map[string]interface{}, error)

func (b *AttachLoadBalancerBuilder) LoadBalancerID(loadBalancerID string) *AttachLoadBalancerBuilder

type AttachLoadBalancerResp struct {
    StatusCode     int
    LoadBalancerId string
    BackendServers struct {
        BackendServer []BackendServerInfo
    }
}

func ParseAttachLoadBalancerResp(resp interface{}) (attachLoadBalancerResp AttachLoadBalancerResp, err error)

type BackendServerInfo struct {
    ServerId string
    Weight   int
    Type     string
}

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
    CreateLoadBalancer struct represents attribute of create LoadBalancer.

type CreateLoadBalancerBuilder struct {
    // contains filtered or unexported fields
}
    CreateLoadBalancer builder pattern code

func NewCreateLoadBalancerBuilder() *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) AddressType(addressType string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) Autopay(autopay bool) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) Bandwidth(bandwidth int) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) Build() (map[string]interface{}, error)

func (b *CreateLoadBalancerBuilder) ClientToken(clientToken string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) Duration(duration string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) InternetChargeType(internetChargeType string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) LoadBalancerName(loadBalancerName string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) LoadBalancerSpec(loadBalancerSpec string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) MasterZoneID(masterZoneID string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) PayType(payType string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) PricingCycle(pricingCycle string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) RegionID(regionID string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) ResourceGroupID(resourceGroupID string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) SlaveZoneID(slaveZoneID string) *CreateLoadBalancerBuilder

func (b *CreateLoadBalancerBuilder) VSwitchID(vSwitchID string) *CreateLoadBalancerBuilder

type CreateLoadBalancerResp struct {
    StatusCode       int
    LoadBalancerId   string
    Address          string
    NetworkType      string
    MasterZoneId     string
    SlaveZoneId      string
    LoadBalancerName string
}

func ParseCreateLoadBalancerResp(resp interface{}) (createLoadBalancerResp CreateLoadBalancerResp, err error)

type DeleteLoadBalancer struct {
    RegionID       string
    LoadBalancerID string
}
    DeleteLoadBalancer struct represents attribute of delete LoadBalancer.

type DeleteLoadBalancerBuilder struct {
    // contains filtered or unexported fields
}
    DeleteLoadBalancer builder pattern code

func NewDeleteLoadBalancerBuilder() *DeleteLoadBalancerBuilder

func (b *DeleteLoadBalancerBuilder) Build() (map[string]interface{}, error)

func (b *DeleteLoadBalancerBuilder) LoadBalancerID(loadBalancerID string) *DeleteLoadBalancerBuilder

func (b *DeleteLoadBalancerBuilder) RegionID(regionID string) *DeleteLoadBalancerBuilder

type DetachLoadBalancer struct {
    RegionID       string
    LoadBalancerID string
    BackendServers string
}
    DetachLoadBalancer represents Detach node with loadbalancer attribute.

type DetachLoadBalancerBuilder struct {
    // contains filtered or unexported fields
}
    DetachLoadBalancer builder pattern code

func NewDetachLoadBalancerBuilder() *DetachLoadBalancerBuilder

func (b *DetachLoadBalancerBuilder) BackendServers(backendServers string) *DetachLoadBalancerBuilder

func (b *DetachLoadBalancerBuilder) Build() (map[string]interface{}, error)

func (b *DetachLoadBalancerBuilder) LoadBalancerID(loadBalancerID string) *DetachLoadBalancerBuilder

func (b *DetachLoadBalancerBuilder) RegionID(regionID string) *DetachLoadBalancerBuilder

type DetachLoadBalancerResp struct {
    StatusCode     int
    LoadBalancerId string
    BackendServers struct {
        BackendServer []BackendServerInfo
    }
}

func ParseDetachLoadBalancerResp(resp interface{}) (detachLoadBalancerResp DetachLoadBalancerResp, err error)

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
    ListLoadBalancer struct represents attribute of list LoadBalancer.

type ListLoadBalancerBuilder struct {
    // contains filtered or unexported fields
}
    ListLoadBalancer builder pattern code

func NewListLoadBalancerBuilder() *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) Address(address string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) AddressType(addressType string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) Build() (map[string]interface{}, error)

func (b *ListLoadBalancerBuilder) InternetChargeType(internetChargeType string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) LoadBalancerID(loadBalancerID string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) LoadBalancerName(loadBalancerName string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) MasterZoneID(masterZoneID string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) NetworkType(networkType string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) RegionID(regionID string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) ServerID(serverID string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) ServerIntranetAddress(serverIntranetAddress int) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) SlaveZoneID(slaveZoneID string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) VpcID(vpcID string) *ListLoadBalancerBuilder

func (b *ListLoadBalancerBuilder) VswitchID(vswitchID string) *ListLoadBalancerBuilder
```

