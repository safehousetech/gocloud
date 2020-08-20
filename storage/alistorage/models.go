package alistorage

import (
	"errors"

	"github.com/safehousetech/gocloud/aliauth"
)

// Alistorage struct represents Alistorage attribute and method associates with it.
type Alistorage struct {
}

// CreateDisk to store all attribute to create Ali-cloud ECS-Disk
type CreateDisk struct {
	RegionID     string
	ZoneID       string
	DiskName     string
	Description  string
	Encrypted    bool
	DiskCategory string
	Size         int
	SnapshotID   string
	ClientToken  string
}

// DeleteDisk to store all attribute to delete Ali-cloud ECS-Disk
type DeleteDisk struct {
	DiskID string
}

// AttachDisk to store all attribute of attach Ali-cloud ECS-Disk to ECS
type AttachDisk struct {
	InstanceID         string
	DiskID             string
	DeleteWithInstance bool
}

// DetachDisk to store all attribute of detach Ali-cloud ECS-Disk from ECS
type DetachDisk struct {
	InstanceID string
	DiskID     string
}

// CreateSnapshot to store all attribute of create Ali-cloud ECS-Disk 's Snapshot
type CreateSnapshot struct {
	DiskID       string
	SnapshotName string
	Description  string
	ClientToken  string
}

// DeleteSnapshot to store all attribute of delete Ali-cloud ECS-Disk 's Snapshot
type DeleteSnapshot struct {
	SnapshotID string
}

// CreateDiskBuilder pattern code
type CreateDiskBuilder struct {
	createDisk *CreateDisk
}

//NewCreateDiskBuilder ...
func NewCreateDiskBuilder() *CreateDiskBuilder {
	createDisk := &CreateDisk{}
	b := &CreateDiskBuilder{createDisk: createDisk}
	return b
}

//RegionID ...
func (b *CreateDiskBuilder) RegionID(regionID string) *CreateDiskBuilder {
	b.createDisk.RegionID = regionID
	return b
}

//ZoneID ...
func (b *CreateDiskBuilder) ZoneID(zoneID string) *CreateDiskBuilder {
	b.createDisk.ZoneID = zoneID
	return b
}

//DiskName ...
func (b *CreateDiskBuilder) DiskName(diskName string) *CreateDiskBuilder {
	b.createDisk.DiskName = diskName
	return b
}

//Description ...
func (b *CreateDiskBuilder) Description(description string) *CreateDiskBuilder {
	b.createDisk.Description = description
	return b
}

//Encrypted ...
func (b *CreateDiskBuilder) Encrypted(encrypted bool) *CreateDiskBuilder {
	b.createDisk.Encrypted = encrypted
	return b
}

//DiskCategory ...
func (b *CreateDiskBuilder) DiskCategory(diskCategory string) *CreateDiskBuilder {
	b.createDisk.DiskCategory = diskCategory
	return b
}

//Size ...
func (b *CreateDiskBuilder) Size(size int) *CreateDiskBuilder {
	b.createDisk.Size = size
	return b
}

//SnapshotID ..
func (b *CreateDiskBuilder) SnapshotID(snapshotID string) *CreateDiskBuilder {
	b.createDisk.SnapshotID = snapshotID
	return b
}

//ClientToken ...
func (b *CreateDiskBuilder) ClientToken(clientToken string) *CreateDiskBuilder {
	b.createDisk.ClientToken = clientToken
	return b
}

//Build ...
func (b *CreateDiskBuilder) Build() (map[string]interface{}, error) {
	if b.createDisk.RegionID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "RegionID")
	}
	if b.createDisk.ZoneID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "ZoneID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.createDisk)
	return params, nil
}

// DeleteDiskBuilder pattern code
type DeleteDiskBuilder struct {
	deleteDisk *DeleteDisk
}

//NewDeleteDiskBuilder ...
func NewDeleteDiskBuilder() *DeleteDiskBuilder {
	deleteDisk := &DeleteDisk{}
	b := &DeleteDiskBuilder{deleteDisk: deleteDisk}
	return b
}

//DiskID ...
func (b *DeleteDiskBuilder) DiskID(diskID string) *DeleteDiskBuilder {
	b.deleteDisk.DiskID = diskID
	return b
}

//Build ...
func (b *DeleteDiskBuilder) Build() (map[string]interface{}, error) {
	if b.deleteDisk.DiskID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "DiskID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.deleteDisk)
	return params, nil
}

// AttachDiskBuilder pattern code
type AttachDiskBuilder struct {
	attachDisk *AttachDisk
}

//NewAttachDiskBuilder ...
func NewAttachDiskBuilder() *AttachDiskBuilder {
	attachDisk := &AttachDisk{}
	b := &AttachDiskBuilder{attachDisk: attachDisk}
	return b
}

//InstanceID ...
func (b *AttachDiskBuilder) InstanceID(instanceID string) *AttachDiskBuilder {
	b.attachDisk.InstanceID = instanceID
	return b
}

//DiskID ...
func (b *AttachDiskBuilder) DiskID(diskID string) *AttachDiskBuilder {
	b.attachDisk.DiskID = diskID
	return b
}

//DeleteWithInstance ...
func (b *AttachDiskBuilder) DeleteWithInstance(deleteWithInstance bool) *AttachDiskBuilder {
	b.attachDisk.DeleteWithInstance = deleteWithInstance
	return b
}

//Build ...
func (b *AttachDiskBuilder) Build() (map[string]interface{}, error) {
	if b.attachDisk.InstanceID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "InstanceID")
	}
	if b.attachDisk.DiskID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "DiskID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.attachDisk)
	return params, nil
}

// DetachDiskBuilder pattern code
type DetachDiskBuilder struct {
	detachDisk *DetachDisk
}

//NewDetachDiskBuilder ...
func NewDetachDiskBuilder() *DetachDiskBuilder {
	detachDisk := &DetachDisk{}
	b := &DetachDiskBuilder{detachDisk: detachDisk}
	return b
}

//InstanceID ...
func (b *DetachDiskBuilder) InstanceID(instanceID string) *DetachDiskBuilder {
	b.detachDisk.InstanceID = instanceID
	return b
}

//DiskID ...
func (b *DetachDiskBuilder) DiskID(diskID string) *DetachDiskBuilder {
	b.detachDisk.DiskID = diskID
	return b
}

//Build ...
func (b *DetachDiskBuilder) Build() (map[string]interface{}, error) {
	if b.detachDisk.InstanceID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "InstanceID")
	}
	if b.detachDisk.DiskID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "DiskID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.detachDisk)
	return params, nil
}

// CreateSnapshotBuilder pattern code
type CreateSnapshotBuilder struct {
	createSnapshot *CreateSnapshot
}

//NewCreateSnapshotBuilder ...
func NewCreateSnapshotBuilder() *CreateSnapshotBuilder {
	createSnapshot := &CreateSnapshot{}
	b := &CreateSnapshotBuilder{createSnapshot: createSnapshot}
	return b
}

//DiskID ...
func (b *CreateSnapshotBuilder) DiskID(diskID string) *CreateSnapshotBuilder {
	b.createSnapshot.DiskID = diskID
	return b
}

//SnapshotName ...
func (b *CreateSnapshotBuilder) SnapshotName(snapshotName string) *CreateSnapshotBuilder {
	b.createSnapshot.SnapshotName = snapshotName
	return b
}

//Description ...
func (b *CreateSnapshotBuilder) Description(description string) *CreateSnapshotBuilder {
	b.createSnapshot.Description = description
	return b
}

// ClientToken ...
func (b *CreateSnapshotBuilder) ClientToken(clientToken string) *CreateSnapshotBuilder {
	b.createSnapshot.ClientToken = clientToken
	return b
}

//Build ...
func (b *CreateSnapshotBuilder) Build() (map[string]interface{}, error) {
	if b.createSnapshot.DiskID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "DiskID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.createSnapshot)
	return params, nil
}

// DeleteSnapshotBuilder pattern code
type DeleteSnapshotBuilder struct {
	deleteSnapshot *DeleteSnapshot
}

//NewDeleteSnapshotBuilder ...
func NewDeleteSnapshotBuilder() *DeleteSnapshotBuilder {
	deleteSnapshot := &DeleteSnapshot{}
	b := &DeleteSnapshotBuilder{deleteSnapshot: deleteSnapshot}
	return b
}

//SnapshotID ...
func (b *DeleteSnapshotBuilder) SnapshotID(snapshotID string) *DeleteSnapshotBuilder {
	b.deleteSnapshot.SnapshotID = snapshotID
	return b
}

//Build ...
func (b *DeleteSnapshotBuilder) Build() (map[string]interface{}, error) {
	if b.deleteSnapshot.SnapshotID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "SnapshotID")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.deleteSnapshot)
	return params, nil
}
