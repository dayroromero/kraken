package kraken

import "time"

// Object Type: Table
type FacilityAssetsStr struct {
	uuid          string
	Enabled       bool      //for audit purposes
	CreationDate  time.Time //for audit purposes
	DisableDate   time.Time //for audit purposes
	CreatedBy     string    //for audit purposes
	Updated       time.Time //for audit purposes
	UpdatedBy     string    //for audit purposes
	Name          string
	InventoryCode string // unique for facility
	AssetType     string // requires additional table
	FacilityId    string // uuid for owner facility - foreign key to facilityies table uuid
}
