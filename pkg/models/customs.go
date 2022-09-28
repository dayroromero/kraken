package kraken

import (
	"time"

	"github.com/google/uuid"
)

// Object Type: Table
// Stores the customs document information for each import operation.
type ImportDocumentStr struct {
	Uuid                  uuid.UUID `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	Enabled               bool      //for audit purposes
	CreationDate          time.Time //for audit purposes
	DisableDate           time.Time //for audit purposes
	CreatedBy             uuid.UUID //for audit purposes
	Updated               time.Time //for audit purposes
	UpdatedBy             uuid.UUID //for audit purposes
	CustomsDocumentNumber string    // Unique document id issued by the local national customs office
	ProductId             uuid.UUID // Product Id - foreign key
	HSCode                string    // product HSCode at destination (foreign code)
	BusinessPartnerId     uuid.UUID // BP Id - foreign key to business partner table with IsCustomer True
	Quantity              float64   // Quantity in Metric Tonnes
	ShippingBOLId         uuid.UUID // foreign key to bill of lading table
	Closed                bool      // closed when the inventory available is 0
}
