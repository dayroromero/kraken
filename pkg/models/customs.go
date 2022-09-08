package kraken

import "time"

// Object Type: Table
// Stores the customs document information for each import operation.
type ImportDocumentStr struct {
	uuid                  string
	Enabled               bool      //for audit purposes
	CreationDate          time.Time //for audit purposes
	DisableDate           time.Time //for audit purposes
	CreatedBy             string    //for audit purposes
	Updated               time.Time //for audit purposes
	UpdatedBy             string    //for audit purposes
	CustomsDocumentNumber string    // Unique document id issued by the local national customs office
	ProductId             string    // Product Id - foreign key
	HSCode                string    // product HSCode at destination (foreign code)
	BusinessPartnerId     string    // BP Id - foreign key to business partner table with IsCustomer True
	Quantity              float64   // Quantity in Metric Tonnes
	ShippingBOLId         string    // foreign key to bill of lading table
	Closed                bool      // closed when the inventory available is 0
}
