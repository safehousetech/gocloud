package vultrdns

import (
	"errors"

	"github.com/safehousetech/gocloud/vultrauth"
)

//VultrDNS ...
type VultrDNS struct {
}

//CreateDNS ...
type CreateDNS struct {
	domain   string // Domain name to add record to
	name     string // Name (subdomain) of record
	Type     string // Type (A, AAAA, MX, etc) of record
	data     string // Data for this record
	ttl      int    // (optional) TTL of this record
	priority int    // (only required for MX and SRV) Priority of this record (omit the priority from the data)
}

//ListDNS ...
type ListDNS struct {
	domain string // Domain to list records for
}

//DeleteDNS ...
type DeleteDNS struct {
	domain   string // Domain name to delete record from
	RECORDID int    // ID of record to delete (see /dns/records)
}

// CreateDNSBuilder pattern code
type CreateDNSBuilder struct {
	createDNS *CreateDNS
}

//NewCreateDNSBuilder ...
func NewCreateDNSBuilder() *CreateDNSBuilder {
	createDNS := &CreateDNS{}
	b := &CreateDNSBuilder{createDNS: createDNS}
	return b
}

//Domain ..
func (b *CreateDNSBuilder) Domain(domain string) *CreateDNSBuilder {
	b.createDNS.domain = domain
	return b
}

//Name ..
func (b *CreateDNSBuilder) Name(name string) *CreateDNSBuilder {
	b.createDNS.name = name
	return b
}

//Type ..
func (b *CreateDNSBuilder) Type(typeV string) *CreateDNSBuilder {
	b.createDNS.Type = typeV
	return b
}

//Data ..
func (b *CreateDNSBuilder) Data(data string) *CreateDNSBuilder {
	b.createDNS.data = data
	return b
}

//TTL ...
func (b *CreateDNSBuilder) TTL(ttl int) *CreateDNSBuilder {
	b.createDNS.ttl = ttl
	return b
}

//Priority ...
func (b *CreateDNSBuilder) Priority(priority int) *CreateDNSBuilder {
	b.createDNS.priority = priority
	return b
}

//Build ...
func (b *CreateDNSBuilder) Build() (map[string]interface{}, error) {
	if b.createDNS.domain == "" {
		return nil, errors.New(vultrauth.StrMissRequired + "domain")
	}
	if b.createDNS.name == "" {
		return nil, errors.New(vultrauth.StrMissRequired + "name")
	}
	if b.createDNS.Type == "" {
		return nil, errors.New(vultrauth.StrMissRequired + "Type")
	}
	if b.createDNS.data == "" {
		return nil, errors.New(vultrauth.StrMissRequired + "data")
	}

	params := make(map[string]interface{})
	params["domain"] = b.createDNS.domain
	params["name"] = b.createDNS.name
	params["type"] = b.createDNS.Type
	params["data"] = b.createDNS.data

	if b.createDNS.ttl != 0 {
		params["ttl"] = b.createDNS.ttl
	}
	if b.createDNS.priority != 0 {
		params["priority"] = b.createDNS.priority
	}

	return params, nil
}

// ListDNSBuilder pattern code
type ListDNSBuilder struct {
	listDNS *ListDNS
}

//NewListDNSBuilder ..
func NewListDNSBuilder() *ListDNSBuilder {
	listDNS := &ListDNS{}
	b := &ListDNSBuilder{listDNS: listDNS}
	return b
}

//Domain ...
func (b *ListDNSBuilder) Domain(domain string) *ListDNSBuilder {
	b.listDNS.domain = domain
	return b
}

//Build ...
func (b *ListDNSBuilder) Build() (map[string]interface{}, error) {
	if b.listDNS.domain == "" {
		return nil, errors.New(vultrauth.StrMissRequired + "domain")
	}
	params := map[string]interface{}{
		"domain": b.listDNS.domain,
	}
	return params, nil
}

// DeleteDNSBuilder pattern code
type DeleteDNSBuilder struct {
	deleteDNS *DeleteDNS
}

//NewDeleteDNSBuilder ..
func NewDeleteDNSBuilder() *DeleteDNSBuilder {
	deleteDNS := &DeleteDNS{}
	b := &DeleteDNSBuilder{deleteDNS: deleteDNS}
	return b
}

//Domain ..
func (b *DeleteDNSBuilder) Domain(domain string) *DeleteDNSBuilder {
	b.deleteDNS.domain = domain
	return b
}

//RecordID ...
func (b *DeleteDNSBuilder) RecordID(recordID int) *DeleteDNSBuilder {
	b.deleteDNS.RECORDID = recordID
	return b
}

//Build ..
func (b *DeleteDNSBuilder) Build() (map[string]interface{}, error) {
	if b.deleteDNS.domain == "" {
		return nil, errors.New(vultrauth.StrMissRequired + "domain")
	}
	if b.deleteDNS.RECORDID == 0 {
		return nil, errors.New(vultrauth.StrMissRequired + "RECORDID")
	}
	params := map[string]interface{}{
		"domain":   b.deleteDNS.domain,
		"RECORDID": b.deleteDNS.RECORDID,
	}
	return params, nil
}
