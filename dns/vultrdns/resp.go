package vultrdns

import (
	"encoding/json"
)

//ListDNSResp ...
type ListDNSResp struct {
	StatusCode int
	DNSSlice   []DNSInfo
}

//DNSInfo ...
type DNSInfo struct {
	Tpye     string `json:"tpye"`
	Name     string `json:"name"`
	Data     string `json:"data"`
	Priority int    `json:"priority"`
	RecordID int    `json:"RECORDID"`
	TTL      int    `json:"ttl"`
}

//ParseListDNSResp ...
func ParseListDNSResp(resp interface{}) (listDNSResp ListDNSResp, err error) {
	response := resp.(map[string]interface{})
	err = json.Unmarshal([]byte(response["body"].(string)), &listDNSResp.DNSSlice)
	listDNSResp.StatusCode = response["status"].(int)
	return
}
