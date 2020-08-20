package googledns

//Googledns struct represents Googledns attribute and methods associates with it.
type Googledns struct {
}

//CreateDNS struct represents CreateDNS attribute.
type CreateDNS struct {
	CreationTime  string   `json:"creationTime"`
	Description   string   `json:"description"`
	DNSName       string   `json:"dnsName"`
	ID            string   `json:"id"`
	Kind          string   `json:"kind"`
	Name          string   `json:"name"`
	NameServerSet string   `json:"nameServerSet"`
	NameServers   []string `json:"nameServers"`
}
