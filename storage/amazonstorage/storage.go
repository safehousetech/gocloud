package amazonstorage

import (
	"strconv"
)

//CreateDisk ...
func (amazonstorage *AmazonStorage) CreateDisk(request interface{}) (resp interface{}, err error) {

	param, _ := request.(map[string]interface{})

	var createVolume CreateVolume
	var region string
	for key, value := range param {
		switch key {

		case "Region":
			regionV, _ := value.(string)
			region = regionV

		case "AvailZone":
			availZoneV, _ := value.(string)
			createVolume.AvailZone = availZoneV

		case "VolumeType":
			VolumeTypeV, _ := value.(string)
			createVolume.VolumeType = VolumeTypeV

		case "VolumeSize":
			VolumeSizeV, _ := value.(int)
			createVolume.VolumeSize = VolumeSizeV

		case "IOPS":
			IOPSV, _ := value.(int64)
			createVolume.IOPS = IOPSV

		case "Encrypted":
			EncryptedV, _ := value.(bool)
			createVolume.Encrypted = EncryptedV

		case "SnapshotId":
			SnapshotIDV, _ := value.(string)
			createVolume.SnapshotID = SnapshotIDV
		}
	}

	params := makeParams("CreateVolume")
	prepareVolume(params, createVolume)

	response := make(map[string]interface{})

	err = amazonstorage.PrepareSignatureV2query(params, region, response)
	if err != nil {
		return nil, err
	}
	resp = response
	return resp, nil

}

//DeleteDisk ...DeleteDisk
func (amazonstorage *AmazonStorage) DeleteDisk(request interface{}) (resp interface{}, err error) {
	param, _ := request.(map[string]string)
	params := makeParams("DeleteVolume")
	params["VolumeId"] = param["VolumeId"]
	region := param["Region"]

	response := make(map[string]interface{})

	err = amazonstorage.PrepareSignatureV2query(params, region, response)
	if err != nil {
		return nil, err
	}
	resp = response
	return resp, nil

}

//CreateSnapshot ...
func (amazonstorage *AmazonStorage) CreateSnapshot(request interface{}) (resp interface{}, err error) {

	param, _ := request.(map[string]string)
	params := makeParams("CreateSnapshot")

	params["VolumeId"] = param["VolumeId"]
	params["Description"] = param["Description"]
	Region := param["Region"]
	response := make(map[string]interface{})

	err = amazonstorage.PrepareSignatureV2query(params, Region, response)
	if err != nil {
		return nil, err
	}
	resp = response
	return resp, nil
}

//DeleteSnapshot ...
func (amazonstorage *AmazonStorage) DeleteSnapshot(request interface{}) (resp interface{}, err error) {
	ids := []string{}
	param, _ := request.(map[string]string)
	ids = append(ids, param["SnapshotId"])
	params := makeParams("DeleteSnapshot")
	Region := param["Region"]
	for i, id := range ids {
		params["SnapshotId."+strconv.Itoa(i+1)] = id
	}

	response := make(map[string]interface{})

	err = amazonstorage.PrepareSignatureV2query(params, Region, response)
	if err != nil {
		return nil, err
	}
	resp = response
	return resp, nil

}

//AttachDisk ...
func (amazonstorage *AmazonStorage) AttachDisk(request interface{}) (resp interface{}, err error) {
	param, _ := request.(map[string]string)
	params := makeParams("AttachVolume")
	params["VolumeId"] = param["VolumeId"]
	params["InstanceId"] = param["InstanceId"]
	params["Device"] = param["Device"]
	Region := param["Region"]
	response := make(map[string]interface{})

	err = amazonstorage.PrepareSignatureV2query(params, Region, response)
	if err != nil {
		return nil, err
	}
	resp = response
	return resp, nil
}

//DetachDisk ...
func (amazonstorage *AmazonStorage) DetachDisk(request interface{}) (resp interface{}, err error) {
	param, _ := request.(map[string]string)
	params := makeParams("DetachVolume")
	params["VolumeId"] = param["VolumeId"]
	params["InstanceId"] = param["InstanceId"]
	params["Device"] = param["Device"]
	if param["Force"] == "true" {
		params["Force"] = "true"
	}
	Region := param["Region"]
	response := make(map[string]interface{})

	err = amazonstorage.PrepareSignatureV2query(params, Region, response)
	if err != nil {
		return nil, err
	}
	resp = response
	return resp, nil
}
