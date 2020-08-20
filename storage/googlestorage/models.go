package googlestorage

//GoogleStorage ...
type GoogleStorage struct {
	Sname  string `json:"sname"`
	SType  string `json:"stype"`
	Szone  string `json:"Szone"`
	SizeGB string `json:"sizeGB"`
}

//Snapshot ...
type Snapshot struct {
	Name                     string                  `json:"name"`
	CreationTimestamp        string                  `json:"creationTimestamp"`
	Description              string                  `json:"description"`
	DiskSizeGB               string                  `json:"diskSizeGB"`
	ID                       string                  `json:"id"`
	Kind                     string                  `json:"kind"`
	LabelFingerprint         string                  `json:"labelFingerprint"`
	SelfLink                 string                  `json:"selfLink"`
	SourceDisk               string                  `json:"sourceDisk"`
	SourceDiskID             string                  `json:"sourceDiskId"`
	Status                   string                  `json:"status"`
	StorageBytes             string                  `json:"storageBytes"`
	StorageBytesStatus       string                  `json:"storageBytesStatus"`
	Licenses                 []string                `json:"licenses"`
	SourceDiskEncryptionKeys SourceDiskEncryptionKey `json:"sourceDiskEncryptionKey"`
	SnapshotEncryptionKeys   SnapshotEncryptionKey   `json:"snapshotEncryptionKey"`
}

//SourceDiskEncryptionKey ...
type SourceDiskEncryptionKey struct {
	RawKey string `json:"rawKey"`
	Sha256 string `json:"sha256"`
}

//SnapshotEncryptionKey ...
type SnapshotEncryptionKey struct {
	RawKey string `json:"rawKey"`
	Sha256 string `json:"sha256"`
}

//AttachDisk ...
type AttachDisk struct {
	Source             string            `json:"source"`
	DeviceName         string            `json:"deviceName"`
	AutoDelete         bool              `json:"autoDelete"`
	Boot               bool              `json:"boot"`
	DiskEncryptionKeys DiskEncryptionKey `json:"diskEncryptionKey"`
	Index              int               `json:"index"`
	Interface          string            `json:"interface"`
	Kind               string            `json:"kind"`
	Licenses           []string          `json:"licenses"`
	Mode               string            `json:"mode"`
	Type               string            `json:"type"`
	InitializeParam    InitializeParams  `json:"initializeParams"`
}

//InitializeParams ...
type InitializeParams struct {
	DiskName                  string                   `json:"diskName"`
	DiskType                  string                   `json:"diskType"`
	DiskSizeGb                string                   `json:"diskSizeGb"`
	SourceImage               string                   `json:"sourceImage"`
	SourceImageEncryptionKeys SourceImageEncryptionKey `json:"sourceImageEncryptionKey"`
}

//CreateDisk ...
type CreateDisk struct {
	Name                         string                      `json:"name"`
	Type                         string                      `json:"type"`
	Zone                         string                      `json:"zone"`
	SizeGb                       string                      `json:"sizeGb"`
	SourceImageEncryptionKeys    SourceImageEncryptionKey    `json:"sourceImageEncryptionKey"`
	DiskEncryptionKeys           DiskEncryptionKey           `json:"diskEncryptionKey"`
	SourceSnapshotEncryptionKeys SourceSnapshotEncryptionKey `json:"sourceSnapshotEncryptionKey"`
	Licenses                     []string                    `json:"licenses"`
	Users                        []string                    `json:"users"`

	CreationTimestamp   string `json:"creationTimestamp"`
	Description         string `json:"description"`
	ID                  string `json:"id"`
	Kind                string `json:"kind"`
	LabelFingerprint    string `json:"labelFingerprint"`
	SourceSnapshotID    string `json:"sourceSnapshotId"`
	Status              string `json:"status"`
	LastAttachTimestamp string `json:"lastAttachTimestamp"`
	LastDetachTimestamp string `json:"lastDetachTimestamp"`
	Options             string `json:"options"`
	SelfLink            string `json:"selfLink"`
	SourceImage         string `json:"sourceImage"`
	SourceImageID       string `json:"sourceImageId"`
	SourceSnapshot      string `json:"sourceSnapshot"`
}

//SourceImageEncryptionKey ...
type SourceImageEncryptionKey struct {
	RawKey string `json:"rawKey"`
	Sha256 string `json:"sha256"`
}

//DiskEncryptionKey ...
type DiskEncryptionKey struct {
	RawKey string `json:"rawKey"`
	Sha256 string `json:"sha256"`
}

//SourceSnapshotEncryptionKey ...
type SourceSnapshotEncryptionKey struct {
	RawKey string `json:"rawKey"`
	Sha256 string `json:"sha256"`
}
