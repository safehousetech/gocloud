package gce

import (
	"testing"

	googleauth "github.com/safehousetech/gocloud/googleauth"
)

var fixture struct {
	project  string
	instance string
	zone     string
	region   string
}

func init() {
	googleauth.LoadConfig()
	fixture.project = "awesome-delight-283812"
	fixture.instance = "testing-scorelab2"
	fixture.region = "us-east4"
	fixture.zone = fixture.region + "-c"
}

func TestDeleteNode(t *testing.T) {
	var gce GCE
	config := map[string]string{
		"projectid": fixture.project,
		"instance":  fixture.instance,
		"Zone":      fixture.zone,
	}
	gce.DeleteNode(config)
}

func TestStopNode(t *testing.T) {
	var gce GCE
	config := map[string]string{
		"projectid": fixture.project,
		"instance":  fixture.instance,
		"Zone":      fixture.zone,
	}
	_, err := gce.StopNode(config)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestRebootNode(t *testing.T) {
	var gce GCE
	config := map[string]string{
		"projectid": fixture.project,
		"instance":  fixture.instance,
		"Zone":      fixture.zone,
	}
	_, err := gce.RebootNode(config)

	if err != nil {
		t.Errorf("Test Fail")
	}
}

func TestCreateNode(t *testing.T) {
	var gce GCE
	InitializeParams := map[string]string{
		"SourceImage": "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-8-jessie-v20160301",
		"DiskType":    "projects/" + fixture.project + "/zones/" + fixture.zone + "/diskTypes/pd-standard",
		"DiskSizeGb":  "10",
	}

	disks := []map[string]interface{}{
		{
			"Boot":             true,
			"AutoDelete":       false,
			"DeviceName":       "bokya",
			"Type":             "PERSISTENT",
			"Mode":             "READ_WRITE",
			"InitializeParams": InitializeParams,
		},
	}

	AccessConfigs := []map[string]string{{
		"Name": "external-nat",
		"Type": "ONE_TO_ONE_NAT",
	},
	}

	NetworkInterfaces := []map[string]interface{}{
		{
			"Network":       "https://www.googleapis.com/compute/v1/projects/" + fixture.project + "/global/networks/default",
			"Subnetwork":    "projects/sheltermap-1493101612061/regions/" + fixture.region + "/subnetworks/default",
			"AccessConfigs": AccessConfigs,
		},
	}

	createnode := map[string]interface{}{
		"projectid":         fixture.project,
		"InstanceTemplate":  "sb1",
		"Name":              fixture.instance,
		"MachineType":       "https://www.googleapis.com/compute/v1/projects/" + fixture.project + "/zones/" + fixture.zone + "/machineTypes/n1-standard-1",
		"Zone":              fixture.zone,
		"disks":             disks,
		"NetworkInterfaces": NetworkInterfaces,
	}

	_, err := gce.CreateNode(createnode)

	if err != nil {
		t.Errorf("Test Fail")
	}
}
