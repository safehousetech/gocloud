package ec2

import (
	"encoding/xml"
)

//RunInstances runinstance to store all attribute to create EC2 instance
type RunInstances struct {
	ImageID               string
	MinCount              int
	MaxCount              int
	KeyName               string
	InstanceType          string
	SecurityGroups        []SecurityGroup
	KernelID              string
	RamdiskID             string
	UserData              []byte
	AvailZone             string
	PlacementGroupName    string
	Monitoring            bool
	SubnetID              string
	DisableAPITermination bool
	ShutdownBehavior      string
	PrivateIPAddress      string
	BlockDeviceMappings   []BlockDeviceMapping
	NetworkInterfaces     []RunNetworkInterface
}

//SecurityGroup struct
type SecurityGroup struct {
	ID   string `xml:"groupId"`
	Name string `xml:"groupName"`
}

// BlockDeviceMapping struct to attach device
type BlockDeviceMapping struct {
	DeviceName          string `xml:"deviceName"`
	VirtualName         string `xml:"virtualName"`
	SnapshotID          string `xml:"ebs>snapshotId"`
	VolumeType          string `xml:"ebs>volumeType"`
	VolumeSize          int64  `xml:"ebs>volumeSize"`
	DeleteOnTermination bool   `xml:"ebs>deleteOnTermination"`
	IOPS                int64  `xml:"ebs>iops"`
}

//RunNetworkInterface struct for Ec2
type RunNetworkInterface struct {
	ID                      string
	DeviceIndex             int
	SubnetID                string
	Description             string
	PrivateIPs              []PrivateIP
	SecurityGroupIds        []string
	DeleteOnTermination     bool
	SecondaryPrivateIPCount int
}

//PrivateIP Private ip to assign PrivateIP
type PrivateIP struct {
	Address   string `xml:"privateIpAddress"`
	DNSName   string `xml:"privateDnsName"`
	IsPrimary bool   `xml:"primary"`
}

//RunInstancesResp run instance response
type RunInstancesResp struct {
	RequestID      string          `xml:"requestId"`
	ReservationID  string          `xml:"reservationId"`
	OwnerID        string          `xml:"ownerId"`
	SecurityGroups []SecurityGroup `xml:"groupSet>item"`
	Instances      []Instance      `xml:"instancesSet>item"`
}

// Instance represents running instance
type Instance struct {
	InstanceID         string             `xml:"instanceId"`
	InstanceType       string             `xml:"instanceType"`
	ImageID            string             `xml:"imageId"`
	PrivateDNSName     string             `xml:"privateDnsName"`
	DNSName            string             `xml:"dnsName"`
	IPAddress          string             `xml:"ipAddress"`
	PrivateIPAddress   string             `xml:"privateIpAddress"`
	SubnetID           string             `xml:"subnetId"`
	VPCId              string             `xml:"vpcId"`
	SourceDestCheck    bool               `xml:"sourceDestCheck"`
	KeyName            string             `xml:"keyName"`
	AMILaunchIndex     int                `xml:"amiLaunchIndex"`
	Hypervisor         string             `xml:"hypervisor"`
	VirtType           string             `xml:"virtualizationType"`
	Monitoring         string             `xml:"monitoring>state"`
	AvailZone          string             `xml:"placement>availabilityZone"`
	PlacementGroupName string             `xml:"placement>groupName"`
	State              InstanceState      `xml:"instanceState"`
	Tags               []Tag              `xml:"tagSet>item"`
	SecurityGroups     []SecurityGroup    `xml:"groupSet>item"`
	NetworkInterfaces  []NetworkInterface `xml:"networkInterfaceSet>item"`
}

//InstanceStateChange represents instance state change
type InstanceStateChange struct {
	InstanceID    string        `xml:"instanceId"`
	CurrentState  InstanceState `xml:"currentState"`
	PreviousState InstanceState `xml:"previousState"`
}

//SimpleResp ..
type SimpleResp struct {
	XMLName   xml.Name
	RequestID string `xml:"requestId"`
}

//TerminateInstancesResp ...
type TerminateInstancesResp struct {
	RequestID    string                `xml:"requestId"`
	StateChanges []InstanceStateChange `xml:"instancesSet>item"`
}

// InstanceState struct
type InstanceState struct {
	Code int    `xml:"code"`
	Name string `xml:"name"`
}

//NetworkInterface ...
type NetworkInterface struct {
	ID               string                     `xml:"networkInterfaceId"`
	SubnetID         string                     `xml:"subnetId"`
	VPCId            string                     `xml:"vpcId"`
	AvailZone        string                     `xml:"availabilityZone"`
	Description      string                     `xml:"description"`
	OwnerID          string                     `xml:"ownerId"`
	RequesterID      string                     `xml:"requesterId"`
	RequesterManaged bool                       `xml:"requesterManaged"`
	Status           string                     `xml:"status"`
	MACAddress       string                     `xml:"macAddress"`
	PrivateIPAddress string                     `xml:"privateIpAddress"`
	PrivateDNSName   string                     `xml:"privateDnsName"`
	SourceDestCheck  bool                       `xml:"sourceDestCheck"`
	Groups           []SecurityGroup            `xml:"groupSet>item"`
	Attachment       NetworkInterfaceAttachment `xml:"attachment"`
	Tags             []Tag                      `xml:"tagSet>item"`
	PrivateIPs       []PrivateIP                `xml:"privateIpAddressesSet>item"`
}

//NetworkInterfaceAttachment ...
type NetworkInterfaceAttachment struct {
	ID                  string `xml:"attachmentId"`
	InstanceID          string `xml:"instanceId"`
	InstanceOwnerID     string `xml:"instanceOwnerId"`
	DeviceIndex         int    `xml:"deviceIndex"`
	Status              string `xml:"status"`
	AttachTime          string `xml:"attachTime"`
	DeleteOnTermination bool   `xml:"deleteOnTermination"`
}

//Tag reperent tag assgin to instance
type Tag struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

//StartInstanceResp start instance response
type StartInstanceResp struct {
	RequestID    string                `xml:"requestId"`
	StateChanges []InstanceStateChange `xml:"instancesSet>item"`
}

//StopInstanceResp stop instances response
type StopInstanceResp struct {
	RequestID    string                `xml:"requestId"`
	StateChanges []InstanceStateChange `xml:"instancesSet>item"`
}
