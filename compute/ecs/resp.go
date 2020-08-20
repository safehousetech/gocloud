package ecs

import (
	"encoding/json"
)

//CreateNodeResp ...
type CreateNodeResp struct {
	StatusCode int
	InstanceID string
}

//ParseCreateNodeResp ...
func ParseCreateNodeResp(resp interface{}) (createNodeResp CreateNodeResp, err error) {
	response := resp.(map[string]interface{})
	err = json.Unmarshal([]byte(response["body"].(string)), &createNodeResp)
	createNodeResp.StatusCode = response["status"].(int)
	return
}
