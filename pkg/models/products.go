package kraken

import "time"

// Object Type: Table
type ProductStr struct {
	ProdId               string    `json:"uuid"`
	Enabled              bool      //for audit purposes
	CreationDate         time.Time //for audit purposes
	DisableDate          time.Time //for audit purposes
	CreatedBy            string    //for audit purposes
	Updated              time.Time //for audit purposes
	UpdatedBy            string    //for audit purposes
	SKU                  string    //Andikem standard SKU
	ProdName             string
	ProductDescriptor    string // type of product. when there is a product that can have different presentations (e.g caustic soda ash, pearls, flakes, etc.)
	ProdNameLocalization map[string]string
	UnCode               string
	ChrisCode            string
	EinecNumber          string
	HSCode               string //US HS Code - Schedule B
}
