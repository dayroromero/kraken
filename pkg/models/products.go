package models

import (
	"github.com/gofrs/uuid"
)

// Object Type: Table
type GoProduct struct {
	ProductId         uuid.UUID `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	Status            bool      //for audit purposes
	DisableNote       string    //Audit purposes
	SKU               string    //Andikem standard SKU
	ProductName       string
	ProductDescriptor string // type of product. when there is a product that can have different presentations (e.g caustic soda ash, pearls, flakes, etc.)
	// ProdNameLocalization pgtype.JSONB `json:"-" gorm:"column:metadata"`
	UnCode       string
	ChrisCode    string
	EinecNumber  string
	HSCode       string    //US HS Code - Schedule B
	ChemUniqueId uuid.UUID // Unique chemical substance ID
}

// Object Type: Table, represents DTO
type Product struct {
	ProductId         uuid.UUID `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	Status            bool      //for audit purposes
	DisableNote       string    //Audit purposes
	SKU               string    //Andikem standard SKU
	ProductName       string
	ProductDescriptor string // type of product. when there is a product that can have different presentations (e.g caustic soda ash, pearls, flakes, etc.)
	// ProdNameLocalization pgtype.JSONB `json:"-" gorm:"column:metadata"`
	UnCode       string
	ChrisCode    string
	EinecNumber  string
	HSCode       string    //US HS Code - Schedule B
	ChemUniqueId uuid.UUID // Unique chemical substance ID
}

// FromEntity respects the gorm_generics.GormModel interface
// Creates new GORM model from Entity.
func (g Product) FromEntity(goproduct GoProduct) interface{} {
	return Product{
		Status:            goproduct.Status,
		DisableNote:       goproduct.DisableNote,
		ProductId:         goproduct.ProductId,
		SKU:               goproduct.SKU,
		ProductName:       goproduct.ProductName,
		ProductDescriptor: goproduct.ProductDescriptor,
		UnCode:            goproduct.UnCode,
		ChrisCode:         goproduct.ChrisCode,
		EinecNumber:       goproduct.EinecNumber,
		HSCode:            goproduct.HSCode,
		ChemUniqueId:      goproduct.ChemUniqueId,
	}
}

// ToEntity respects the gorm_generics.GormModel interface
// Creates new Entity from GORM model.
func (p Product) ToEntity() GoProduct {
	return GoProduct{
		ProductId:         p.ProductId,
		Status:            p.Status,
		DisableNote:       p.DisableNote,
		SKU:               p.SKU,
		ProductName:       p.ProductName,
		ProductDescriptor: p.ProductDescriptor,
		UnCode:            p.UnCode,
		ChrisCode:         p.ChrisCode,
		EinecNumber:       p.EinecNumber,
		HSCode:            p.HSCode,
		ChemUniqueId:      p.ChemUniqueId,
	}
}
