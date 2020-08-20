package vultrstorage

import (
	"errors"

	"github.com/safehousetech/gocloud/vultrauth"
)

//VultrStorage ...
type VultrStorage struct {
}

//CreateSnapshot ...
type CreateSnapshot struct {
	SUBID       int
	description string // (optional) Description of snapshot contents
}

//DeleteSnapshot ...
type DeleteSnapshot struct {
	SNAPSHOTID string
}

//CreateDisk ...
type CreateDisk struct {
	DCID   int    // DCID of the location to create this subscription in.  See /v1/regions/list
	sizeGB int    // Size (in GB) of this subscription.
	label  string // (optional) Text label that will be associated with the subscription
}

//AttachDisk ...
type AttachDisk struct {
	SUBID         int // ID of the block storage subscription to attach
	attachToSUBID int // ID of the VPS subscription to mount the block storage subscription to
}

//DetachDisk ...
type DetachDisk struct {
	SUBID int // ID of the block storage subscription to detach
}

//DeleteDisk ...
type DeleteDisk struct {
	SUBID int // ID of the block storage subscription to delete
}

//CreateSnapshotBuilder builder pattern code
type CreateSnapshotBuilder struct {
	createSnapshot *CreateSnapshot
}

//NewCreateSnapshotBuilder ...
func NewCreateSnapshotBuilder() *CreateSnapshotBuilder {
	createSnapshot := &CreateSnapshot{}
	b := &CreateSnapshotBuilder{createSnapshot: createSnapshot}
	return b
}

//SUBID ...
func (b *CreateSnapshotBuilder) SUBID(sUBID int) *CreateSnapshotBuilder {
	b.createSnapshot.SUBID = sUBID
	return b
}

//Description ...
func (b *CreateSnapshotBuilder) Description(description string) *CreateSnapshotBuilder {
	b.createSnapshot.description = description
	return b
}

//Build ...
func (b *CreateSnapshotBuilder) Build() (map[string]interface{}, error) {
	if b.createSnapshot.SUBID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "SUBID")
	}
	params := map[string]interface{}{
		"SUBID": b.createSnapshot.SUBID,
	}
	if b.createSnapshot.description != "" {
		params["description"] = b.createSnapshot.description
	}
	return params, nil
}

//DeleteSnapshotBuilder builder pattern code
type DeleteSnapshotBuilder struct {
	deleteSnapshot *DeleteSnapshot
}

//NewDeleteSnapshotBuilder ...
func NewDeleteSnapshotBuilder() *DeleteSnapshotBuilder {
	deleteSnapshot := &DeleteSnapshot{}
	b := &DeleteSnapshotBuilder{deleteSnapshot: deleteSnapshot}
	return b
}

//SNAPSHOTID ...
func (b *DeleteSnapshotBuilder) SNAPSHOTID(sNAPSHOTID string) *DeleteSnapshotBuilder {
	b.deleteSnapshot.SNAPSHOTID = sNAPSHOTID
	return b
}

//Build ...
func (b *DeleteSnapshotBuilder) Build() (map[string]interface{}, error) {
	if b.deleteSnapshot.SNAPSHOTID == "" {
		return nil, errors.New(vultrauth.StrMissRequired + "SNAPSHOTID")
	}
	params := map[string]interface{}{
		"SNAPSHOTID": b.deleteSnapshot.SNAPSHOTID,
	}
	return params, nil
}

//CreateDiskBuilder pattern code
type CreateDiskBuilder struct {
	createDisk *CreateDisk
}

//NewCreateDiskBuilder ...
func NewCreateDiskBuilder() *CreateDiskBuilder {
	createDisk := &CreateDisk{}
	b := &CreateDiskBuilder{createDisk: createDisk}
	return b
}

//DCID ...
func (b *CreateDiskBuilder) DCID(dCID int) *CreateDiskBuilder {
	b.createDisk.DCID = dCID
	return b
}

//SizeGB ...
func (b *CreateDiskBuilder) SizeGB(sizeGB int) *CreateDiskBuilder {
	b.createDisk.sizeGB = sizeGB
	return b
}

//Label ..
func (b *CreateDiskBuilder) Label(label string) *CreateDiskBuilder {
	b.createDisk.label = label
	return b
}

//Build ...
func (b *CreateDiskBuilder) Build() (map[string]interface{}, error) {
	if b.createDisk.DCID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "DCID")
	}
	if b.createDisk.sizeGB == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "size_gb")
	}
	params := make(map[string]interface{})
	params["DCID"] = b.createDisk.DCID
	params["size_gb"] = b.createDisk.sizeGB
	if b.createDisk.label != "" {
		params["label"] = b.createDisk.label
	}
	return params, nil
}

//AttachDiskBuilder pattern code
type AttachDiskBuilder struct {
	attachDisk *AttachDisk
}

//NewAttachDiskBuilder ...
func NewAttachDiskBuilder() *AttachDiskBuilder {
	attachDisk := &AttachDisk{}
	b := &AttachDiskBuilder{attachDisk: attachDisk}
	return b
}

//SUBID ...
func (b *AttachDiskBuilder) SUBID(subID int) *AttachDiskBuilder {
	b.attachDisk.SUBID = subID
	return b
}

//AttachToSUBID ...
func (b *AttachDiskBuilder) AttachToSUBID(attachTOSUBID int) *AttachDiskBuilder {
	b.attachDisk.attachToSUBID = attachTOSUBID
	return b
}

//Build ...
func (b *AttachDiskBuilder) Build() (map[string]interface{}, error) {
	if b.attachDisk.SUBID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "SUBID")
	}
	if b.attachDisk.attachToSUBID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "attach_to_SUBID")
	}
	params := make(map[string]interface{})
	params["SUBID"] = b.attachDisk.SUBID
	params["attach_to_SUBID"] = b.attachDisk.attachToSUBID
	return params, nil
}

//DetachDiskBuilder pattern code
type DetachDiskBuilder struct {
	detachDisk *DetachDisk
}

//NewDetachDiskBuilder ...
func NewDetachDiskBuilder() *DetachDiskBuilder {
	detachDisk := &DetachDisk{}
	b := &DetachDiskBuilder{detachDisk: detachDisk}
	return b
}

//SUBID ...
func (b *DetachDiskBuilder) SUBID(subID int) *DetachDiskBuilder {
	b.detachDisk.SUBID = subID
	return b
}

//Build ...
func (b *DetachDiskBuilder) Build() (map[string]interface{}, error) {
	if b.detachDisk.SUBID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "SUBID")
	}
	params := make(map[string]interface{})
	params["SUBID"] = b.detachDisk.SUBID
	return params, nil
}

//DeleteDiskBuilder pattern code
type DeleteDiskBuilder struct {
	deleteDisk *DeleteDisk
}

//NewDeleteDiskBuilder ...
func NewDeleteDiskBuilder() *DeleteDiskBuilder {
	deleteDisk := &DeleteDisk{}
	b := &DeleteDiskBuilder{deleteDisk: deleteDisk}
	return b
}

//SUBID ..
func (b *DeleteDiskBuilder) SUBID(sUBID int) *DeleteDiskBuilder {
	b.deleteDisk.SUBID = sUBID
	return b
}

//Build ...
func (b *DeleteDiskBuilder) Build() (map[string]interface{}, error) {
	if b.deleteDisk.SUBID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "SUBID")
	}
	params := make(map[string]interface{})
	params["SUBID"] = b.deleteDisk.SUBID
	return params, nil
}
