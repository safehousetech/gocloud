package digioceandns

import (
	"testing"

	digioceanAuth "github.com/safehousetech/gocloud/digioceanauth"
)

func init() {
	digioceanAuth.LoadConfig()
}

func TestCreateDns(t *testing.T) {

	var digioceancloud Digioceandns

	create := map[string]interface{}{
		"DomainName": "example.com",
		"Type":       "A",
		"Name":       "www",
		"Data":       "162.10.66.0",
		"Priority":   nil,
		"Port":       nil,
		"TimeToLive": 1800,
		"Weight":     nil,
		"Flags":      nil,
		"Tag":        nil,
	}

	_, err := digioceancloud.CreateDNS(create)

	if err != nil {
		t.Errorf("Test to create DigitalOcean DNS record failed.")
	}
}

func TestDeleteDns(t *testing.T) {

	var digioceancloud Digioceandns

	delete := map[string]string{
		"DomainName": "example.com",
		"RecordID":   "28448433",
	}

	_, err := digioceancloud.DeleteDNS(delete)

	if err != nil {
		t.Errorf("Test to delete DigitalOcean DNS record failed.")
	}
}

func TestListDns(t *testing.T) {

	var digioceancloud Digioceandns

	listRecords := map[string]string{
		"DomainName": "example.com",
	}

	_, err := digioceancloud.ListDNS(listRecords)

	if err != nil {
		t.Errorf("Test to list DigitalOcean DNS record failed.")
	}
}
