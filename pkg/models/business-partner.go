package kraken

import (
	"time"

	"github.com/google/uuid"
)

// Object Type: Table
// Stores the information regarding business partners, which can be customers, or vendors.
type BusinessPartnerStr struct {
	CustomerId      uuid.UUID           `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	Enabled         bool                //for audit purposes
	CreationDate    time.Time           //for audit purposes
	DisableDate     time.Time           //for audit purposes
	CreatedBy       uuid.UUID           //for audit purposes
	Updated         time.Time           //for audit purposes
	UpdatedBy       uuid.UUID           //for audit purposes
	CustomerName    string              //Bussiness partner's name
	CustomerContact *[]ContactStr       //Contact information for comunication purposes
	Address         *[]EntityAddressStr //Addressess for shipping or mailing
	IsSupplier      bool                //True if the BP is a vendor of products or services to the company (supplier)
	IsCustomer      bool                //True if the BP is a customer of the company
	IsShippingAgent bool                //True if the supplier provides shipping services (for sea transportation)
	IsTrucker       bool                //True if the BP is a supplier of trucking services
	TaxId           *TaxIdStr           //Unique national Tax ID - from the country of residence
	IsChild         bool                //true if the BP is part of another BP (e.g. subsidiary, branch, etc)
	ParentCustomer  uuid.UUID           //uuid if the parent customer (if any)
}
