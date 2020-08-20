package googledns

import "testing"

func TestCreateDns(t *testing.T) {
	var googledns Googledns

	createdns := map[string]interface{}{
		"Project":     "sheltermap-1493101612061",
		"Kind":        "dns#managedZone",
		"Description": "dns",
		"DnsName":     "rootmonk.me.",
		"Name":        "gocloud",
	}
	_, err := googledns.CreateDNS(createdns)

	if err != nil {
		t.Errorf("Test Fail")
	}

}

func TestListDns(t *testing.T) {
	var googledns Googledns

	listdns := map[string]string{
		"Project": "sheltermap-1493101612061",
	}

	_, err := googledns.ListDNS(listdns)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestListResourceDnsRecordSets(t *testing.T) {
	var googledns Googledns

	listResourcednsRecordSets := map[string]string{
		"Project":     "sheltermap-1493101612061",
		"managedZone": "gocloud3",
	}
	_, err := googledns.ListResourceDNSRecordSets(listResourcednsRecordSets)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestDeleteDns(t *testing.T) {
	var googledns Googledns

	deletedns := map[string]string{
		"Project":     "sheltermap-1493101612061",
		"managedZone": "gocloud3",
	}
	_, err := googledns.DeleteDNS(deletedns)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
