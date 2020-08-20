package vultrdns

import (
	"fmt"
	"net/http"

	"github.com/safehousetech/gocloud/vultrauth"
)

// ListResourceDNSRecordSets function lists DNS record sets.
func (vultrDNS *VultrDNS) ListResourceDNSRecordSets(request interface{}) (resp interface{}, err error) {
	fmt.Println("\nThis API is not provided by Vultr cloud")
	return resp, err
}

// ListDNS function lists DNS records.
func (vultrDNS *VultrDNS) ListDNS(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	var domain string

	for key, value := range param {
		switch key {
		case "domain":
			domain = value.(string)
		}
	}

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodGet, "/v1/dns/records?domain="+domain, param, response)

	resp = response
	return resp, err
}

// DeleteDNS function deletes a DNS record.
func (vultrDNS *VultrDNS) DeleteDNS(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/dns/delete_record", param, response)

	resp = response
	return resp, err
}

// CreateDNS function creates a new DNS record.
func (vultrDNS *VultrDNS) CreateDNS(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})
	param = request.(map[string]interface{})

	response := make(map[string]interface{})

	err = vultrauth.SignAndDoRequest(http.MethodPost, "/v1/dns/create_record", param, response)

	resp = response
	return resp, err
}
