package kraken

import (
	"time"

	"github.com/google/uuid"
)

// Object Type: Table
type ProductStr struct {
	ProdId               uuid.UUID `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	Enabled              bool      //for audit purposes
	CreationDate         time.Time //for audit purposes
	DisableDate          time.Time //for audit purposes
	CreatedBy            uuid.UUID //for audit purposes
	Updated              time.Time //for audit purposes
	UpdatedBy            uuid.UUID //for audit purposes
	SKU                  string    //Andikem standard SKU
	ProdName             string
	ProductDescriptor    string // type of product. when there is a product that can have different presentations (e.g caustic soda ash, pearls, flakes, etc.)
	ProdNameLocalization map[string]string
	UnCode               string
	ChrisCode            string
	EinecNumber          string
	HSCode               string //US HS Code - Schedule B
}
