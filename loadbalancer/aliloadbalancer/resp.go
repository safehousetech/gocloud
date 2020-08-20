package aliloadbalancer

import "encoding/json"

//CreateLoadBalancerResp ...
type CreateLoadBalancerResp struct {
	StatusCode       int
	LoadBalancerID   string
	Address          string
	NetworkType      string
	MasterZoneID     string
	SlaveZoneID      string
	LoadBalancerName string
}

//AttachLoadBalancerResp ...
type AttachLoadBalancerResp struct {
	StatusCode     int
	LoadBalancerID string
	BackendServers struct {
		BackendServer []BackendServerInfo
	}
}

//DetachLoadBalancerResp ...
type DetachLoadBalancerResp struct {
	StatusCode     int
	LoadBalancerID string
	BackendServers struct {
		BackendServer []BackendServerInfo
	}
}

//BackendServerInfo ...
type BackendServerInfo struct {
	ServerID string
	Weight   int
	Type     string
}

//ParseCreateLoadBalancerResp ...
func ParseCreateLoadBalancerResp(resp interface{}) (createLoadBalancerResp CreateLoadBalancerResp, err error) {
	response := resp.(map[string]interface{})
	err = json.Unmarshal([]byte(response["body"].(string)), &createLoadBalancerResp)
	createLoadBalancerResp.StatusCode = response["status"].(int)
	return
}

//ParseAttachLoadBalancerResp ...
func ParseAttachLoadBalancerResp(resp interface{}) (attachLoadBalancerResp AttachLoadBalancerResp, err error) {
	response := resp.(map[string]interface{})
	err = json.Unmarshal([]byte(response["body"].(string)), &attachLoadBalancerResp)
	attachLoadBalancerResp.StatusCode = response["status"].(int)
	return
}

//ParseDetachLoadBalancerResp ...
func ParseDetachLoadBalancerResp(resp interface{}) (detachLoadBalancerResp DetachLoadBalancerResp, err error) {
	response := resp.(map[string]interface{})
	err = json.Unmarshal([]byte(response["body"].(string)), &detachLoadBalancerResp)
	detachLoadBalancerResp.StatusCode = response["status"].(int)
	return
}
