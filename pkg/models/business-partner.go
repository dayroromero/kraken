package kraken

import "time"

// Object Type: Table
type BusinessPartnerStr struct {
	CustomerId      string    `json:"uuid"`
	Enabled         bool      //for audit purposes
	CreationDate    time.Time //for audit purposes
	DisableDate     time.Time //for audit purposes
	CreatedBy       string    //for audit purposes
	Updated         time.Time //for audit purposes
	UpdatedBy       string    //for audit purposes
	CustomerName    string
	CustomerContact *[]ContactStr
	IsVendor        bool
	IsCustomer      bool
	IsShippingAgent bool
	IsTrucker       bool
	TaxId           *TaxIdStr
	IsChild         bool
	ParentCustomer  string //uuid if the parent customer
}
