package alidns

import (
	"errors"

	"github.com/safehousetech/gocloud/aliauth"
)

//Alidns represents Alidns attribute and method associates with it.
type Alidns struct {
}

// CreateDNS to store all attribute to create Ali-cloud DNS
type CreateDNS struct {
	DomainName string
	RR         string
	Type       string
	Value      string
	TTL        int
	Priority   int
	Line       string
}

// DeleteDNS to store all attribute to delete Ali-cloud DNS
type DeleteDNS struct {
	RecordID string
}

// ListDNS ...
type ListDNS struct {
	DomainName   string
	PageNumber   int
	PageSize     int
	RRKeyWord    string
	TypeKeyWord  string
	ValueKeyWord string
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

//DomainName ...
func (b *CreateDNSBuilder) DomainName(domainName string) *CreateDNSBuilder {
	b.createDNS.DomainName = domainName
	return b
}

//RR ...
func (b *CreateDNSBuilder) RR(rR string) *CreateDNSBuilder {
	b.createDNS.RR = rR
	return b
}

//Type ...
func (b *CreateDNSBuilder) Type(typeV string) *CreateDNSBuilder {
	b.createDNS.Type = typeV
	return b
}

//Value ...
func (b *CreateDNSBuilder) Value(value string) *CreateDNSBuilder {
	b.createDNS.Value = value
	return b
}

//TTL ...
func (b *CreateDNSBuilder) TTL(tTL int) *CreateDNSBuilder {
	b.createDNS.TTL = tTL
	return b
}

//Priority ...
func (b *CreateDNSBuilder) Priority(priority int) *CreateDNSBuilder {
	b.createDNS.Priority = priority
	return b
}

//Line ..
func (b *CreateDNSBuilder) Line(line string) *CreateDNSBuilder {
	b.createDNS.Line = line
	return b
}

//Build ...
func (b *CreateDNSBuilder) Build() (map[string]interface{}, error) {
	if b.createDNS.DomainName == "" {
		return nil, errors.New(aliauth.StrMissRequired + "DomainName")
	}
	if b.createDNS.RR == "" {
		return nil, errors.New(aliauth.StrMissRequired + "RR")
	}
	if b.createDNS.Type == "" {
		return nil, errors.New(aliauth.StrMissRequired + "Type")
	}
	if b.createDNS.Value == "" {
		return nil, errors.New(aliauth.StrMissRequired + "Value")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.createDNS)
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

//RecordID ..
func (b *DeleteDNSBuilder) RecordID(recordID string) *DeleteDNSBuilder {
	b.deleteDNS.RecordID = recordID
	return b
}

//Build ..
func (b *DeleteDNSBuilder) Build() (map[string]interface{}, error) {
	if b.deleteDNS.RecordID == "" {
		return nil, errors.New(aliauth.StrMissRequired + "RecordId")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.deleteDNS)
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

//DomainName ...
func (b *ListDNSBuilder) DomainName(domainName string) *ListDNSBuilder {
	b.listDNS.DomainName = domainName
	return b
}

//PageNumber ...
func (b *ListDNSBuilder) PageNumber(pageNumber int) *ListDNSBuilder {
	b.listDNS.PageNumber = pageNumber
	return b
}

//PageSize ..
func (b *ListDNSBuilder) PageSize(pageSize int) *ListDNSBuilder {
	b.listDNS.PageSize = pageSize
	return b
}

//RRKeyWord ..
func (b *ListDNSBuilder) RRKeyWord(rRKeyWord string) *ListDNSBuilder {
	b.listDNS.RRKeyWord = rRKeyWord
	return b
}

//TypeKeyWord ..
func (b *ListDNSBuilder) TypeKeyWord(typeKeyWord string) *ListDNSBuilder {
	b.listDNS.TypeKeyWord = typeKeyWord
	return b
}

//ValueKeyWord ..
func (b *ListDNSBuilder) ValueKeyWord(valueKeyWord string) *ListDNSBuilder {
	b.listDNS.ValueKeyWord = valueKeyWord
	return b
}

//Build ..
func (b *ListDNSBuilder) Build() (map[string]interface{}, error) {
	if b.listDNS.DomainName == "" {
		return nil, errors.New(aliauth.StrMissRequired + "DomainName")
	}
	params := make(map[string]interface{})
	// Put all of options into params
	params = aliauth.PutStructIntoMap(b.listDNS)
	return params, nil
}
