package awsdns

import (
	"testing"

	awsAuth "github.com/safehousetech/gocloud/auth"
)

func init() {
	awsAuth.LoadConfig()
}

func TestCreateDns(t *testing.T) {
	var awsdns Awsdns
	createdns := map[string]interface{}{
		"name":             "rootmonk.me",
		"hostedZoneConfig": "hostedZoneConfig",
	}

	awsdns.CreateDNS(createdns)
}

func TestDeleteDns(t *testing.T) {
	var awsdns Awsdns
	deletedns := map[string]string{
		"ID": "ZOD7SUP0ZGGQQ",
	}

	_, err := awsdns.DeleteDNS(deletedns)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestListDns(t *testing.T) {
	var awsdns Awsdns

	listdns := map[string]interface{}{
		"marker":   "",
		"maxItems": 2,
	}

	_, err := awsdns.ListDNS(listdns)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestListResourceDnsRecordSets(t *testing.T) {
	var awsdns Awsdns
	listResourcednsRecordSets := map[string]interface{}{
		"zone": "ZBNX5TIW033J2",
	}

	_, err := awsdns.ListResourceDNSRecordSets(listResourcednsRecordSets)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
