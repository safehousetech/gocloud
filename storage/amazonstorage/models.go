package amazonstorage

import (
	"encoding/xml"
)

//AmazonStorage ...
type AmazonStorage struct {
}

const (
	debug = false

	legacyAPIVersion = "2011-12-15"

	vpcAPIVersion = "2013-10-15"
)

//CreateVolume ...
type CreateVolume struct {
	AvailZone  string
	SnapshotID string
	VolumeType string
	VolumeSize int
	Encrypted  bool
	IOPS       int64
}

//Volume ...
type Volume struct {
	ID          string             `xml:"volumeId"`
	Size        int                `xml:"size"`
	SnapshotID  string             `xml:"snapshotId"`
	Status      string             `xml:"status"`
	IOPS        int64              `xml:"iops"`
	AvailZone   string             `xml:"availabilityZone"`
	CreateTime  string             `xml:"createTime"`
	VolumeType  string             `xml:"volumeType"`
	Encrypted   bool               `xml:"encrypted"`
	Tags        []Tag              `xml:"tagSet>item"`
	Attachments []VolumeAttachment `xml:"attachmentSet>item"`
}

//Tag ..
type Tag struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

//VolumeAttachment ...
type VolumeAttachment struct {
	VolumeID            string `xml:"volumeId"`
	Device              string `xml:"device"`
	InstanceID          string `xml:"instanceId"`
	Status              string `xml:"status"`
	DeleteOnTermination bool   `xml:"deleteOnTermination"`
}

//CreateVolumeResp ...
type CreateVolumeResp struct {
	RequestID string `xml:"requestId"`
	Volume
}

//SimpleResp ...
type SimpleResp struct {
	XMLName   xml.Name
	RequestID string `xml:"requestId"`
}

//VolumeAttachmentResp ...
type VolumeAttachmentResp struct {
	RequestID  string `xml:"requestId"`
	VolumeID   string `xml:"volumeId"`
	Device     string `xml:"device"`
	InstanceID string `xml:"instanceId"`
	Status     string `xml:"status"`
	AttachTime string `xml:"attachTime"`
}

//SnapshotsResp ...
type SnapshotsResp struct {
	RequestID string     `xml:"requestId"`
	Snapshots []Snapshot `xml:"snapshotSet>item"`
}

//Snapshot ...
type Snapshot struct {
	ID          string `xml:"snapshotId"`
	VolumeID    string `xml:"volumeId"`
	VolumeSize  string `xml:"volumeSize"`
	Status      string `xml:"status"`
	StartTime   string `xml:"startTime"`
	Description string `xml:"description"`
	Progress    string `xml:"progress"`
	OwnerID     string `xml:"ownerId"`
	OwnerAlias  string `xml:"ownerAlias"`
	Tags        []Tag  `xml:"tagSet>item"`
}

//CreateSnapshotResp ...
type CreateSnapshotResp struct {
	RequestID string `xml:"requestId"`
	Snapshot
}
