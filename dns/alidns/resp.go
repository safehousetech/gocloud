package alidns

import "encoding/json"

//ListDNSResp ..
type ListDNSResp struct {
	StatusCode    int
	DomainRecords struct {
		Record []RecordInfo
	}
}

//RecordInfo ...
type RecordInfo struct {
	DomainName string
	RecordID   string
	RR         string
	Type       string
	Value      string
	Line       string
	Priority   int
	TTL        int
	Status     string
	Locked     bool
}

//ParseListDNSResp ..
func ParseListDNSResp(resp interface{}) (listDNSResp ListDNSResp, err error) {
	response := resp.(map[string]interface{})
	err = json.Unmarshal([]byte(response["body"].(string)), &listDNSResp)
	listDNSResp.StatusCode = response["status"].(int)
	return
}
