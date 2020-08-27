package googlestorage

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	googleauth "github.com/safehousetech/gocloud/googleauth"
)

//CreateDisk ...
func (googlestorage *GoogleStorage) CreateDisk(request interface{}) (resp interface{}, err error) {

	var option CreateDisk

	var projectID string

	var zone string

	var _type string

	param := request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "projectid":
			projectID, _ = value.(string)

		case "Name":
			name, _ := value.(string)
			option.Name = name

		case "Zone":
			zoneV, _ := value.(string)
			zone = zoneV

		case "Type":
			typeV, _ := value.(string)
			_type = typeV

		case "SizeGb":
			sizeGbV, _ := value.(string)
			option.SizeGb = sizeGbV

		case "SourceImageEncryptionKeys":
			sourceImageEncryptionKeysV, _ := value.(map[string]string)
			option.SourceImageEncryptionKeys.RawKey = sourceImageEncryptionKeysV["RawKey"]
			option.SourceImageEncryptionKeys.Sha256 = sourceImageEncryptionKeysV["Sha256"]

		case "DiskEncryptionKeys":
			diskEncryptionKeysV, _ := value.(map[string]string)
			option.DiskEncryptionKeys.RawKey = diskEncryptionKeysV["RawKey"]
			option.DiskEncryptionKeys.Sha256 = diskEncryptionKeysV["Sha256"]

		case "SourceSnapshotEncryptionKeys":
			sourceSnapshotEncryptionKeysV, _ := value.(map[string]string)
			option.SourceSnapshotEncryptionKeys.RawKey = sourceSnapshotEncryptionKeysV["RawKey"]
			option.SourceSnapshotEncryptionKeys.Sha256 = sourceSnapshotEncryptionKeysV["Sha256"]

		case "Licenses":
			licensesV, _ := value.([]string)
			option.Licenses = licensesV

		case "Users":
			usersV, _ := value.([]string)
			option.Users = usersV

		case "CreationTimestamp":
			creationTimestampV, _ := value.(string)
			option.CreationTimestamp = creationTimestampV

		case "Description":
			descriptionV, _ := value.(string)
			option.Description = descriptionV

		case "ID":
			idV, _ := value.(string)
			option.ID = idV

		case "Kind":
			kindV, _ := value.(string)
			option.Kind = kindV

		case "LabelFingerprint":
			labelFingerprintV, _ := value.(string)
			option.LabelFingerprint = labelFingerprintV

		case "SourceSnapshotID":
			sourceSnapshotIDV, _ := value.(string)
			option.SourceSnapshotID = sourceSnapshotIDV

		case "Status":
			statusV, _ := value.(string)
			option.Status = statusV

		case "LastAttachTimestamp":
			lastAttachTimestampV, _ := value.(string)
			option.LastAttachTimestamp = lastAttachTimestampV

		case "LastDetachTimestamp":
			lastDetachTimestampV, _ := value.(string)
			option.LastDetachTimestamp = lastDetachTimestampV

		case "Options":
			optionsV, _ := value.(string)
			option.Options = optionsV

		case "SelfLink":
			selfLinkV, _ := value.(string)
			option.SelfLink = selfLinkV

		case "SourceImage":
			sourceImage, _ := value.(string)
			option.SourceImage = sourceImage

		case "SourceImageID":
			sourceImageIDV, _ := value.(string)
			option.SourceImageID = sourceImageIDV

		case "SourceSnapshot":
			sourceSnapshotV, _ := value.(string)
			option.SourceSnapshot = sourceSnapshotV

		}
	}

	zonevalue := "projects/" + projectID + "/zones/" + zone
	option.Zone = zonevalue

	typeValue := "projects/" + projectID + "/zones/" + zone + "/diskTypes/" + _type
	option.Type = typeValue

	creatDiskJSONmap := make(map[string]interface{})

	CreateDiskDictnoaryConvert(option, creatDiskJSONmap)

	creatDiskJSON, _ := json.Marshal(creatDiskJSONmap)

	creatDiskJSONstring := string(creatDiskJSON)

	var creatDiskJSONstringByte = []byte(creatDiskJSONstring)

	url := "https://www.googleapis.com/compute/v1/projects/" + projectID + "/zones/" + zone + "/disks"

	client := googleauth.Client()

	createDiskRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(creatDiskJSONstringByte))

	createDiskRequest.Header.Set("Content-Type", "application/json")

	createDiskResp, err := client.Do(createDiskRequest)

	if err != nil {
		return nil, err
	}

	defer createDiskResp.Body.Close()

	body, err := ioutil.ReadAll(createDiskResp.Body)

	createDiskResponse := make(map[string]interface{})
	createDiskResponse["status"] = createDiskResp.StatusCode
	createDiskResponse["body"] = string(body)
	resp = createDiskResponse
	return resp, err
}

//DeleteDisk ...
func (googlestorage *GoogleStorage) DeleteDisk(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/" + options["Zone"] + "/disks/" + options["disk"]

	client := googleauth.Client()

	deleteDiskRequest, err := http.NewRequest("DELETE", url, nil)
	deleteDiskRequest.Header.Set("Content-Type", "application/json")

	deleteDiskResp, err := client.Do(deleteDiskRequest)

	if err != nil {
		return nil, err
	}

	defer deleteDiskResp.Body.Close()

	body, err := ioutil.ReadAll(deleteDiskResp.Body)

	deleteDiskResponse := make(map[string]interface{})
	deleteDiskResponse["status"] = deleteDiskResp.StatusCode
	deleteDiskResponse["body"] = string(body)
	resp = deleteDiskResponse
	return resp, err
}

//CreateSnapshot ...
func (googlestorage *GoogleStorage) CreateSnapshot(request interface{}) (resp interface{}, err error) {

	var snapshot Snapshot
	var projectID string
	var Zone string
	var Disk string

	param := request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "projectid":
			projectID, _ = value.(string)

		case "Name":
			nameV, _ := value.(string)
			snapshot.Name = nameV

		case "Zone":
			zoneV, _ := value.(string)
			Zone = zoneV

		case "disk":
			diskV, _ := value.(string)
			Disk = diskV

		case "CreationTimestamp":
			creationTimestampV, _ := value.(string)
			snapshot.CreationTimestamp = creationTimestampV

		case "Description":
			descriptionV, _ := value.(string)
			snapshot.Description = descriptionV

		case "DiskSizeGb":
			diskSizeGBV, _ := value.(string)
			snapshot.DiskSizeGB = diskSizeGBV

		case "ID":
			idV, _ := value.(string)
			snapshot.ID = idV

		case "Kind":
			kindV, _ := value.(string)
			snapshot.Kind = kindV

		case "LabelFingerprint":
			labelFingerprintV, _ := value.(string)
			snapshot.LabelFingerprint = labelFingerprintV

		case "SelfLink":
			selfLinkV, _ := value.(string)
			snapshot.SelfLink = selfLinkV

		case "StorageBytes":
			storageBytesV, _ := value.(string)
			snapshot.StorageBytes = storageBytesV

		case "Status":
			statusV, _ := value.(string)
			snapshot.Status = statusV

		case "SourceDiskID":
			sourceDiskIDV, _ := value.(string)
			snapshot.SourceDiskID = sourceDiskIDV

		case "SourceDisk":
			sourceDiskV, _ := value.(string)
			snapshot.SourceDisk = sourceDiskV

		case "StorageBytesStatus":
			storageBytesStatusV, _ := value.(string)
			snapshot.StorageBytesStatus = storageBytesStatusV

		case "Licenses":
			licensesV, _ := value.([]string)
			snapshot.Licenses = licensesV

		case "SourceDiskEncryptionKeys":
			sourceDiskEncryptionKeysV, _ := value.(map[string]string)
			snapshot.SourceDiskEncryptionKeys.RawKey = sourceDiskEncryptionKeysV["RawKey"]
			snapshot.SourceDiskEncryptionKeys.Sha256 = sourceDiskEncryptionKeysV["Sha256"]

		case "SnapshotEncryptionKeys":
			snapshotEncryptionKeysV, _ := value.(map[string]string)
			snapshot.SnapshotEncryptionKeys.RawKey = snapshotEncryptionKeysV["RawKey"]
			snapshot.SnapshotEncryptionKeys.Sha256 = snapshotEncryptionKeysV["Sha256"]

		}
	}

	createSnapshotJSONmap := make(map[string]interface{})

	CreateSnapshotdictnoaryconvert(snapshot, createSnapshotJSONmap)

	createSnapshotJSON, _ := json.Marshal(createSnapshotJSONmap)
	createSnapshotString := string(createSnapshotJSON)

	var createSnapshotStringByte = []byte(createSnapshotString)

	url := "https://www.googleapis.com/compute/v1/projects/" + projectID + "/zones/" + Zone + "/disks/" + Disk + "/createSnapshot"

	client := googleauth.Client()

	createSnapshotRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(createSnapshotStringByte))
	createSnapshotRequest.Header.Set("Content-Type", "application/json")

	createSnapshotResp, err := client.Do(createSnapshotRequest)

	if err != nil {
		return nil, err
	}

	defer createSnapshotResp.Body.Close()

	body, err := ioutil.ReadAll(createSnapshotResp.Body)

	createSnapshotResponse := make(map[string]interface{})
	createSnapshotResponse["status"] = createSnapshotResp.StatusCode
	createSnapshotResponse["body"] = string(body)
	resp = createSnapshotResponse
	return resp, err
}

//DeleteSnapshot ...
func (googlestorage *GoogleStorage) DeleteSnapshot(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/global/snapshots/" + options["snapshot"]

	client := googleauth.Client()

	deleteSnapshotRequest, err := http.NewRequest("DELETE", url, nil)
	deleteSnapshotRequest.Header.Set("Content-Type", "application/json")

	deleteSnapshotResp, err := client.Do(deleteSnapshotRequest)

	if err != nil {
		return nil, err
	}

	defer deleteSnapshotResp.Body.Close()
	body, err := ioutil.ReadAll(deleteSnapshotResp.Body)

	deleteSnapshotResponse := make(map[string]interface{})
	deleteSnapshotResponse["status"] = deleteSnapshotResp.StatusCode
	deleteSnapshotResponse["body"] = string(body)
	resp = deleteSnapshotResponse
	return resp, err
}

//AttachDisk ...
func (googlestorage *GoogleStorage) AttachDisk(request interface{}) (resp interface{}, err error) {

	var attachdisk AttachDisk
	var projectID string
	var zone string
	var instance string

	param := request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "projectid":
			projectID, _ = value.(string)

		case "Zone":
			zoneV, _ := value.(string)
			zone = zoneV

		case "Source":
			source, _ := value.(string)
			attachdisk.Source = source

		case "instance":
			instanceV, _ := value.(string)
			instance = instanceV

		case "Licenses":
			licensesV, _ := value.([]string)
			attachdisk.Licenses = licensesV

		case "DiskEncryptionKeys":
			diskEncryptionKeysV, _ := value.(map[string]string)
			attachdisk.DiskEncryptionKeys.RawKey = diskEncryptionKeysV["RawKey"]
			attachdisk.DiskEncryptionKeys.Sha256 = diskEncryptionKeysV["Sha256"]

		case "Mode":
			modeV, _ := value.(string)
			attachdisk.Mode = modeV

		case "Type":
			typeV, _ := value.(string)
			attachdisk.Type = typeV

		case "Kind":
			kindV, _ := value.(string)
			attachdisk.Kind = kindV

		case "Interface":
			interfaceV, _ := value.(string)
			attachdisk.Interface = interfaceV

		case "Index":
			indexV, _ := value.(int)
			attachdisk.Index = indexV

		case "Boot":
			bootV, _ := value.(bool)
			attachdisk.Boot = bootV

		case "AutoDelete":
			autoDeleteV, _ := value.(bool)
			attachdisk.AutoDelete = autoDeleteV

		case "DeviceName":
			deviceNameV, _ := value.(string)
			attachdisk.DeviceName = deviceNameV

		case "InitializeParam":
			initializeParamV, _ := value.(map[string]interface{})
			for key, value := range initializeParamV {
				switch key {
				case "DiskName":
					diskNameV, _ := value.(string)
					attachdisk.InitializeParam.DiskName = diskNameV

				case "DiskType":
					diskTypeV, _ := value.(string)
					attachdisk.InitializeParam.DiskType = diskTypeV

				case "DiskSizeGb":
					diskSizeGbV, _ := value.(string)
					attachdisk.InitializeParam.DiskSizeGb = diskSizeGbV

				case "SourceImage":
					sourceImageV, _ := value.(string)
					attachdisk.InitializeParam.SourceImage = sourceImageV

				case "SourceImageEncryptionKeys":
					SourceImageEncryptionKeysV, _ := value.(map[string]string)
					attachdisk.InitializeParam.SourceImageEncryptionKeys.RawKey = SourceImageEncryptionKeysV["RawKey"]
					attachdisk.InitializeParam.SourceImageEncryptionKeys.Sha256 = SourceImageEncryptionKeysV["Sha256"]
				}
			}
		}
	}

	attachDiskJSONmap := make(map[string]interface{})

	AttachDiskdictnoaryconvert(attachdisk, attachDiskJSONmap)

	attachDiskJSON, _ := json.Marshal(attachDiskJSONmap)

	attachDiskJSONstring := string(attachDiskJSON)

	var attachDiskJSONstringByte = []byte(attachDiskJSONstring)

	url := "https://www.googleapis.com/compute/v1/projects/" + projectID + "/zones/" + zone + "/instances/" + instance + "/attachDisk"

	client := googleauth.Client()

	attachDiskRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(attachDiskJSONstringByte))
	attachDiskRequest.Header.Set("Content-Type", "application/json")

	attachDiskResp, err := client.Do(attachDiskRequest)

	if err != nil {
		return nil, err
	}

	defer attachDiskResp.Body.Close()

	body, err := ioutil.ReadAll(attachDiskResp.Body)

	attachDiskResponse := make(map[string]interface{})
	attachDiskResponse["status"] = attachDiskResp.StatusCode
	attachDiskResponse["body"] = string(body)
	resp = attachDiskResponse
	return resp, err
}

//DetachDisk ..
func (googlestorage *GoogleStorage) DetachDisk(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/compute/v1/projects/" + options["projectid"] + "/zones/" + options["Zone"] + "/instances/" + options["instance"] + "/detachDisk"

	client := googleauth.Client()

	detachDiskRequest, err := http.NewRequest("POST", url, nil)

	detachDiskRequestParam := detachDiskRequest.URL.Query()

	detachDiskRequestParam.Add("deviceName", options["deviceName"])

	detachDiskRequest.URL.RawQuery = detachDiskRequestParam.Encode()

	detachDiskRequest.Header.Set("Content-Type", "application/json")

	detachDiskResp, err := client.Do(detachDiskRequest)

	if err != nil {
		return nil, err
	}

	defer detachDiskResp.Body.Close()

	body, err := ioutil.ReadAll(detachDiskResp.Body)

	detachDiskResponse := make(map[string]interface{})
	detachDiskResponse["status"] = detachDiskResp.StatusCode
	detachDiskResponse["body"] = string(body)
	resp = detachDiskResponse
	return resp, err
}
