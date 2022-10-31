package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

// Represents the interface to interact with gorm's generic methods
type GoBussinesPartner struct {
	BusinessPartnerId uuid.UUID
	Status            bool
	DisableNote       string
	CustomerName      string
	Contact           pgtype.JSONB
	Address           pgtype.JSONB
	IsSupplier        bool
	IsCustomer        bool
	IsShippingAgent   bool
	IsTrucker         bool
	TaxId             pgtype.JSONB
	IsChild           bool
	ParentCustomer    uuid.UUID
}

// Object Type: Table, represents DTO
type BussinesPartner struct {
	BusinessPartnerId uuid.UUID    `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	Status            bool         //for audit purposes
	DisableNote       string       //for audit purposes
	CustomerName      string       //Bussiness partner's name
	Contact           pgtype.JSONB `json:"-" gorm:"column:metadata"` //Contact information for comunication purposes
	Address           pgtype.JSONB `json:"-" gorm:"column:metadata"` //Addressess for shipping or mailing
	IsSupplier        bool         //True if the BP is a vendor of products or services to the company (supplier)
	IsCustomer        bool         //True if the BP is a customer of the company
	IsShippingAgent   bool         //True if the supplier provides shipping services (for sea transportation)
	IsTrucker         bool         //True if the BP is a supplier of trucking services
	TaxId             pgtype.JSONB `json:"-" gorm:"column:metadata"` //Unique national Tax ID - from the country of residence
	IsChild           bool         //true if the BP is part of another BP (e.g. subsidiary, branch, etc)
	ParentCustomer    uuid.UUID    //uuid if the parent customer (if any)
}

// FromEntity respects the gorm_generics.GormModel interface
// Creates new GORM model from Entity.
func (g BussinesPartner) FromEntity(goBP GoBussinesPartner) interface{} {
	return BussinesPartner{
		BusinessPartnerId: goBP.BusinessPartnerId,
		Status:            goBP.Status,
		DisableNote:       goBP.DisableNote,
		CustomerName:      goBP.CustomerName,
		Contact:           goBP.Contact,
		Address:           goBP.Address,
		IsSupplier:        goBP.IsSupplier,
		IsCustomer:        goBP.IsCustomer,
		IsShippingAgent:   goBP.IsShippingAgent,
		IsTrucker:         goBP.IsTrucker,
		TaxId:             goBP.TaxId,
		IsChild:           goBP.IsChild,
		ParentCustomer:    goBP.ParentCustomer,
	}
}

// ToEntity respects the gorm_generics.GormModel interface
// Creates new Entity from GORM model.
func (bp BussinesPartner) ToEntity() GoBussinesPartner {
	return GoBussinesPartner{
		BusinessPartnerId: bp.BusinessPartnerId,
		Status:            bp.Status,
		DisableNote:       bp.DisableNote,
		CustomerName:      bp.CustomerName,
		Contact:           bp.Contact,
		Address:           bp.Address,
		IsSupplier:        bp.IsSupplier,
		IsCustomer:        bp.IsCustomer,
		IsShippingAgent:   bp.IsShippingAgent,
		IsTrucker:         bp.IsTrucker,
		TaxId:             bp.TaxId,
		IsChild:           bp.IsChild,
		ParentCustomer:    bp.ParentCustomer,
	}
}

// ContactStr stores basic contact information
type Contact struct {
	IsPrimary bool
	Name      string
	Email     string
	Phone     string
	Address   pgtype.JSONB `json:"-" gorm:"column:metadata"`
}
