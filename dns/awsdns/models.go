package awsdns

//Awsdns represents Awsdns attribute and method associates with it.
type Awsdns struct {
}

//ListResourceDNSRecordSets represents ListResourceDNSRecordSets.
type ListResourceDNSRecordSets struct {
	Name       string
	Type       string
	Identifier string
	MaxItems   int
}

//ChangeResourceRecordSetsRequest respeptes ChangeResourceRecordSetsRequest.
type ChangeResourceRecordSetsRequest struct {
	Comment string   `xml:"ChangeBatch>Comment,omitempty"`
	Changes []Change `xml:"ChangeBatch>Changes>Change"`
}

//Change represents Change.
type Change struct {
	Action string            `xml:"Action"`
	Record ResourceRecordSet `xml:"ResourceRecordSet"`
}

//AliasTarget represents AliasTarget.
type AliasTarget struct {
	HostedZoneID         string
	DNSName              string
	EvaluateTargetHealth bool
}

//HostedZone represents HostedZone.
type HostedZone struct {
	ID              string `xml:"Id"`
	Name            string `xml:"Name"`
	CallerReference string `xml:"CallerReference"`
	Comment         string `xml:"Config>Comment"`
	ResourceCount   int    `xml:"ResourceRecordSetCount"`
}

//ChangeInfo represents ChangeInfo.
type ChangeInfo struct {
	ID          string `xml:"Id"`
	Status      string `xml:"Status"`
	SubmittedAt string `xml:"SubmittedAt"`
}

//DelegationSet represents DelegationSet.
type DelegationSet struct {
	NameServers []string `xml:"NameServers>NameServer"`
}

//ResourceRecordSet represents ResourceRecordSet.
type ResourceRecordSet struct {
	Name          string       `xml:"Name"`
	Type          string       `xml:"Type"`
	TTL           int          `xml:"TTL"`
	Records       []string     `xml:"ResourceRecords>ResourceRecord>Value,omitempty"`
	SetIdentifier string       `xml:"SetIdentifier,omitempty"`
	Weight        int          `xml:"Weight,omitempty"`
	HealthCheckID string       `xml:"HealthCheckId,omitempty"`
	Region        string       `xml:"Region,omitempty"`
	Failover      string       `xml:"Failover,omitempty"`
	AliasTarget   *AliasTarget `xml:"AliasTarget,omitempty"`
	RecordsXML    string       `xml:",innerxml"`
}

//CreateHostedZoneRequest represents CreateHostedZoneRequest.
type CreateHostedZoneRequest struct {
	Name            string `xml:"Name"`
	CallerReference string `xml:"CallerReference"`
	Comment         string `xml:"HostedZoneConfig>Comment"`
}

//CreateHostedZoneResponse represents CreateHostedZoneResponse.
type CreateHostedZoneResponse struct {
	HostedZone    HostedZone    `xml:"HostedZone"`
	ChangeInfo    ChangeInfo    `xml:"ChangeInfo"`
	DelegationSet DelegationSet `xml:"DelegationSet"`
}

//DeleteHostedZoneResponse represents DeleteHostedZoneResponse.
type DeleteHostedZoneResponse struct {
	ChangeInfo ChangeInfo `xml:"ChangeInfo"`
}

//ListHostedZonesResponse represents ListHostedZonesResponse.
type ListHostedZonesResponse struct {
	HostedZones []HostedZone `xml:"HostedZones>HostedZone"`
	Marker      string       `xml:"Marker"`
	IsTruncated bool         `xml:"IsTruncated"`
	NextMarker  string       `xml:"NextMarker"`
	MaxItems    int          `xml:"MaxItems"`
}

//ChangeResourceRecordSetsResponse represents ChangeResourceRecordSetsResponse
type ChangeResourceRecordSetsResponse struct {
	ChangeInfo ChangeInfo `xml:"ChangeInfo"`
}

//ListResourceRecordSetsResponse represents ListResourceRecordSetsResponse.
type ListResourceRecordSetsResponse struct {
	Records              []ResourceRecordSet `xml:"ResourceRecordSets>ResourceRecordSet"`
	IsTruncated          bool                `xml:"IsTruncated"`
	MaxItems             int                 `xml:"MaxItems"`
	NextRecordName       string              `xml:"NextRecordName"`
	NextRecordType       string              `xml:"NextRecordType"`
	NextRecordIdentifier string              `xml:"NextRecordIdentifier"`
}
