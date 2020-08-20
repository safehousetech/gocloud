package googlestorage

//CreateDiskDictnoaryConvert ...
func CreateDiskDictnoaryConvert(option CreateDisk, Creatdiskjsonmap map[string]interface{}) {

	if option.Name != "" {
		Creatdiskjsonmap["name"] = option.Name
	}

	if option.Type != "" {
		Creatdiskjsonmap["type"] = option.Type
	}

	if option.Zone != "" {
		Creatdiskjsonmap["zone"] = option.Zone
	}

	if option.SizeGb != "" {
		Creatdiskjsonmap["sizeGb"] = option.SizeGb
	}

	if option.SourceImageEncryptionKeys != (SourceImageEncryptionKey{}) {
		Creatdiskjsonmap["sourceImageEncryptionKey"] = option.SourceImageEncryptionKeys
	}
	if option.DiskEncryptionKeys != (DiskEncryptionKey{}) {
		Creatdiskjsonmap["diskEncryptionKey"] = option.DiskEncryptionKeys
	}
	if option.SourceSnapshotEncryptionKeys != (SourceSnapshotEncryptionKey{}) {
		Creatdiskjsonmap["sourceSnapshotEncryptionKey"] = option.SourceSnapshotEncryptionKeys
	}

	if len(option.Licenses) != 0 {
		Creatdiskjsonmap["licenses"] = option.Licenses
	}

	if len(option.Users) != 0 {
		Creatdiskjsonmap["users"] = option.Users
	}

	if option.LastAttachTimestamp != "" {
		Creatdiskjsonmap["lastAttachTimestamp"] = option.LastAttachTimestamp
	}

	if option.SourceSnapshot != "" {
		Creatdiskjsonmap["SourceSnapshot"] = option.SourceSnapshot
	}

	if option.LastDetachTimestamp != "" {
		Creatdiskjsonmap["lastDetachTimestamp"] = option.LastDetachTimestamp
	}

	if option.Options != "" {
		Creatdiskjsonmap["options"] = option.Options
	}

	if option.SelfLink != "" {
		Creatdiskjsonmap["selfLink"] = option.SelfLink
	}

	if option.SourceImage != "" {
		Creatdiskjsonmap["sourceImage"] = option.SourceImage
	}

	if option.SourceImageID != "" {
		Creatdiskjsonmap["sourceImageID"] = option.SourceImageID
	}

	if option.Status != "" {
		Creatdiskjsonmap["status"] = option.Status
	}

	if option.SourceSnapshotID != "" {
		Creatdiskjsonmap["sourceSnapshotID"] = option.SourceSnapshotID
	}

	if option.LabelFingerprint != "" {
		Creatdiskjsonmap["labelFingerprint"] = option.LabelFingerprint
	}

	if option.Kind != "" {
		Creatdiskjsonmap["kind"] = option.Kind
	}

	if option.ID != "" {
		Creatdiskjsonmap["id"] = option.ID
	}

	if option.Description != "" {
		Creatdiskjsonmap["description"] = option.Description
	}

	if option.CreationTimestamp != "" {
		Creatdiskjsonmap["creationTimestamp"] = option.CreationTimestamp
	}

}

//CreateSnapshotdictnoaryconvert ..
func CreateSnapshotdictnoaryconvert(option Snapshot, CreateSnapshotJSONmap map[string]interface{}) {

	if len(option.Licenses) != 0 {
		CreateSnapshotJSONmap["licenses"] = option.Licenses
	}

	if option.Name != "" {
		CreateSnapshotJSONmap["name"] = option.Name
	}

	if option.CreationTimestamp != "" {
		CreateSnapshotJSONmap["creationTimestamp"] = option.CreationTimestamp
	}

	if option.Description != "" {
		CreateSnapshotJSONmap["description"] = option.Description
	}

	if option.DiskSizeGB != "" {
		CreateSnapshotJSONmap["diskSizeGB"] = option.DiskSizeGB
	}

	if option.ID != "" {
		CreateSnapshotJSONmap["id"] = option.ID
	}

	if option.Kind != "" {
		CreateSnapshotJSONmap["kind"] = option.Kind
	}

	if option.Status != "" {
		CreateSnapshotJSONmap["status"] = option.Status
	}

	if option.SelfLink != "" {
		CreateSnapshotJSONmap["selfLink"] = option.SelfLink
	}

	if option.LabelFingerprint != "" {
		CreateSnapshotJSONmap["labelFingerprint"] = option.LabelFingerprint
	}

	if option.StorageBytes != "" {
		CreateSnapshotJSONmap["storageBytes"] = option.StorageBytes
	}

	if option.SourceDisk != "" {
		CreateSnapshotJSONmap["sourceDisk"] = option.SourceDisk
	}

	if option.SourceDiskID != "" {
		CreateSnapshotJSONmap["sourceDiskID"] = option.SourceDiskID
	}

	if option.StorageBytesStatus != "" {
		CreateSnapshotJSONmap["storageBytesStatus"] = option.StorageBytesStatus
	}

	if option.SourceDiskEncryptionKeys != (SourceDiskEncryptionKey{}) {
		CreateSnapshotJSONmap["sourceDiskEncryptionKey"] = option.SourceDiskEncryptionKeys
	}

	if option.SnapshotEncryptionKeys != (SnapshotEncryptionKey{}) {
		CreateSnapshotJSONmap["snapshotEncryptionKey"] = option.SnapshotEncryptionKeys
	}

}

//AttachDiskdictnoaryconvert ...
func AttachDiskdictnoaryconvert(option AttachDisk, AttachDiskjsonmap map[string]interface{}) {

	if len(option.Licenses) != 0 {
		AttachDiskjsonmap["licenses"] = option.Licenses
	}

	if option.DiskEncryptionKeys != (DiskEncryptionKey{}) {
		AttachDiskjsonmap["diskEncryptionKey"] = option.DiskEncryptionKeys
	}

	if option.Source != "" {
		AttachDiskjsonmap["source"] = option.Source
	}

	if option.DeviceName != "" {
		AttachDiskjsonmap["deviceName"] = option.DeviceName
	}

	if option.AutoDelete {
		AttachDiskjsonmap["autoDelete"] = option.AutoDelete
	}

	if option.Boot {
		AttachDiskjsonmap["boot"] = option.Boot
	}

	if option.Index != 0 {
		AttachDiskjsonmap["index"] = option.Index
	}

	if option.Interface != "" {
		AttachDiskjsonmap["interface"] = option.Interface
	}

	if option.Kind != "" {
		AttachDiskjsonmap["kind"] = option.Kind
	}

	if option.Mode != "" {
		AttachDiskjsonmap["mode"] = option.Mode
	}

	if option.Type != "" {
		AttachDiskjsonmap["type"] = option.Type
	}

	if option.Type != "" {
		AttachDiskjsonmap["type"] = option.Type
	}

	if option.InitializeParam != (InitializeParams{}) {

		InitializeParamMap := make(map[string]interface{})

		if option.InitializeParam.DiskName != "" {
			InitializeParamMap["diskName"] = option.InitializeParam.DiskName
		}

		if option.InitializeParam.DiskType != "" {
			InitializeParamMap["diskType"] = option.InitializeParam.DiskType
		}

		if option.InitializeParam.DiskSizeGb != "" {
			InitializeParamMap["diskSizeGb"] = option.InitializeParam.DiskSizeGb
		}

		if option.InitializeParam.SourceImage != "" {
			InitializeParamMap["sourceImage"] = option.InitializeParam.SourceImage
		}

		if (option.InitializeParam.SourceImageEncryptionKeys != SourceImageEncryptionKey{}) {

			SourceImageEncryptionKeysMap := make(map[string]interface{})

			if option.InitializeParam.SourceImageEncryptionKeys.RawKey != "" {
				SourceImageEncryptionKeysMap["rawKey"] = option.InitializeParam.SourceImageEncryptionKeys.RawKey
			}
			if option.InitializeParam.SourceImageEncryptionKeys.RawKey != "" {
				SourceImageEncryptionKeysMap["sha256"] = option.InitializeParam.SourceImageEncryptionKeys.Sha256
			}

			InitializeParamMap["sourceImage"] = SourceImageEncryptionKeysMap
		}

		AttachDiskjsonmap["initializeParams"] = InitializeParamMap
	}

}
