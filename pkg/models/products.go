package models

import (
	"github.com/gofrs/uuid"
)

// Object Type: Table
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

// Object Type: Table
type ProductGorm struct {
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
func (g ProductGorm) FromEntity(product Product) interface{} {
	return ProductGorm{
		ProductId:         product.ProductId,
		Status:            product.Status,
		DisableNote:       product.DisableNote,
		SKU:               product.SKU,
		ProductName:       product.ProductName,
		ProductDescriptor: product.ProductDescriptor,
		UnCode:            product.UnCode,
		ChrisCode:         product.ChrisCode,
		EinecNumber:       product.EinecNumber,
		HSCode:            product.HSCode,
		ChemUniqueId:      product.ChemUniqueId,
	}
}

// ToEntity respects the gorm_generics.GormModel interface
// Creates new Entity from GORM model.
func (g ProductGorm) ToEntity() Product {
	return Product{
		ProductId:         g.ProductId,
		Status:            g.Status,
		DisableNote:       g.DisableNote,
		SKU:               g.SKU,
		ProductName:       g.ProductName,
		ProductDescriptor: g.ProductDescriptor,
		UnCode:            g.UnCode,
		ChrisCode:         g.ChrisCode,
		EinecNumber:       g.EinecNumber,
		HSCode:            g.HSCode,
		ChemUniqueId:      g.ChemUniqueId,
	}
}
