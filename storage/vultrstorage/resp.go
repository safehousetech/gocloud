package vultrstorage

import "encoding/json"

//CreateDiskResp ...
type CreateDiskResp struct {
	StatusCode int
	SUBID      string `json:"SUBID"`
}

//ParseCreateDiskResp ..
func ParseCreateDiskResp(resp interface{}) (createDiskResp CreateDiskResp, err error) {
	response := resp.(map[string]interface{})
	err = json.Unmarshal([]byte(response["body"].(string)), &createDiskResp)
	createDiskResp.StatusCode = response["status"].(int)
	return
}
