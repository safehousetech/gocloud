package googledns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	googleauth "github.com/safehousetech/gocloud/googleauth"
)

// ...
const (
	UnixDate = "Mon Jan _2 15:04:05 MST 2006"
	RFC3339  = "2006-01-02T15:04:05Z07:00"
)

//ListResourceDNSRecordSets list ListResourceDNSRecordSets.
func (googledns *Googledns) ListResourceDNSRecordSets(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/dns/v1/projects/" + options["project"] + "/managedZones/" + options["managedZone"] + "/rrsets"

	client := googleauth.Client()

	ListResourceDNSRecordSetsrequest, err := http.NewRequest("GET", url, nil)

	ListResourceDNSRecordSetsparam := ListResourceDNSRecordSetsrequest.URL.Query()

	if options["maxResults"] != "" {
		ListResourceDNSRecordSetsparam.Add("deviceName", options["maxResults"])
	}

	if options["pageToken"] != "" {
		ListResourceDNSRecordSetsparam.Add("pageToken", options["pageToken"])
	}

	if options["sortBy"] != "" {
		ListResourceDNSRecordSetsparam.Add("sortBy", options["sortBy"])
	}

	if options["sortOrder"] != "" {
		ListResourceDNSRecordSetsparam.Add("sortOrder", options["sortOrder"])
	}

	ListResourceDNSRecordSetsrequest.URL.RawQuery = ListResourceDNSRecordSetsparam.Encode()

	ListResourceDNSRecordSetsrequest.Header.Set("Content-Type", "application/json")

	ListResourceDNSRecordSetsResp, err := client.Do(ListResourceDNSRecordSetsrequest)

	if err != nil {
		fmt.Print(err)
		return
	}

	defer ListResourceDNSRecordSetsResp.Body.Close()

	body, err := ioutil.ReadAll(ListResourceDNSRecordSetsResp.Body)

	ListResourcednsRecordSetresponse := make(map[string]interface{})
	ListResourcednsRecordSetresponse["status"] = ListResourceDNSRecordSetsResp.StatusCode
	ListResourcednsRecordSetresponse["body"] = string(body)
	resp = ListResourcednsRecordSetresponse
	return resp, err
}

//CreateDNS creates DNS.
func (googledns *Googledns) CreateDNS(request interface{}) (resp interface{}, err error) {

	param := request.(map[string]interface{})
	var Project string
	var option CreateDNS

	for key, value := range param {
		switch key {

		case "Project":
			ProjectV, _ := value.(string)
			Project = ProjectV

		case "CreationTime":
			CreationTimeV, _ := value.(string)
			option.CreationTime = CreationTimeV
			option.CreationTime = time.Now().UTC().Format(time.RFC3339)

		case "Description":
			DescriptionV, _ := value.(string)
			option.Description = DescriptionV

		case "DnsName":
			DNSNameV, _ := value.(string)
			option.DNSName = DNSNameV

		case "nameServers":
			nameServersV, _ := value.([]string)
			option.NameServers = nameServersV

		case "Id":
			IDV, _ := value.(string)
			option.ID = IDV

		case "Kind":
			kindV, _ := value.(string)
			option.Kind = kindV

		case "Name":
			NameV, _ := value.(string)
			option.Name = NameV

		case "nameServerSet":
			NameServerSetV, _ := value.(string)
			option.NameServerSet = NameServerSetV
		}
	}

	CreateDnsjsonmap := make(map[string]interface{})

	CreateDnsedictnoaryconvert(option, CreateDnsjsonmap)

	CreateDnsjson, _ := json.Marshal(CreateDnsjsonmap)

	CreateDnsjsonstring := string(CreateDnsjson)

	var CreateDnsjsonstringbyte = []byte(CreateDnsjsonstring)

	url := "https://www.googleapis.com/dns/v1/projects/" + Project + "/managedZones"

	client := googleauth.Client()

	CreateDnsrequest, err := http.NewRequest("POST", url, bytes.NewBuffer(CreateDnsjsonstringbyte))

	CreateDnsrequest.Header.Set("Content-Type", "application/json")

	CreateDnsrresp, err := client.Do(CreateDnsrequest)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer CreateDnsrresp.Body.Close()

	body, err := ioutil.ReadAll(CreateDnsrresp.Body)

	CreateDnsresponse := make(map[string]interface{})
	CreateDnsresponse["status"] = CreateDnsrresp.StatusCode
	CreateDnsresponse["body"] = string(body)
	resp = CreateDnsresponse
	return resp, err
}

//ListDNS lists DNS.
func (googledns *Googledns) ListDNS(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/dns/v1/projects/" + options["Project"] + "/managedZones/"

	client := googleauth.Client()

	ListDnsrequest, err := http.NewRequest("GET", url, nil)

	ListDnsrequestparam := ListDnsrequest.URL.Query()

	if options["maxResults"] != "" {
		ListDnsrequestparam.Add("deviceName", options["maxResults"])
	}

	if options["pageToken"] != "0" {
		ListDnsrequestparam.Add("pageToken", options["pageToken"])
	}

	ListDnsrequest.URL.RawQuery = ListDnsrequestparam.Encode()

	ListDnsrequest.Header.Set("Content-Type", "application/json")

	ListDnsresp, err := client.Do(ListDnsrequest)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer ListDnsresp.Body.Close()

	body, err := ioutil.ReadAll(ListDnsresp.Body)

	ListDnsresponse := make(map[string]interface{})
	ListDnsresponse["status"] = ListDnsresp.StatusCode
	ListDnsresponse["body"] = string(body)
	resp = ListDnsresponse
	return resp, err
}

//DeleteDNS deletes DNS.
func (googledns *Googledns) DeleteDNS(request interface{}) (resp interface{}, err error) {

	options := request.(map[string]string)

	url := "https://www.googleapis.com/dns/v1/projects/" + options["Project"] + "/managedZones/" + options["managedZone"]

	client := googleauth.Client()

	DeleteDnsrequest, err := http.NewRequest("DELETE", url, nil)

	DeleteDnsrequest.Header.Set("Content-Type", "application/json")

	DeleteDnsresp, err := client.Do(DeleteDnsrequest)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer DeleteDnsresp.Body.Close()

	body, err := ioutil.ReadAll(DeleteDnsresp.Body)

	DeleteDnsresponse := make(map[string]interface{})
	DeleteDnsresponse["status"] = DeleteDnsresp.StatusCode
	DeleteDnsresponse["body"] = string(body)
	resp = DeleteDnsresponse
	return resp, err
}
