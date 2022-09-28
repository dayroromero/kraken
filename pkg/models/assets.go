package kraken

import (
	"time"

	"github.com/google/uuid"
)

// Object Type: Table
// Stores critical assets related to the operation of a facility.
type FacilityAssetsStr struct {
	Uuid          uuid.UUID `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	Enabled       bool      //for audit purposes
	CreationDate  time.Time //for audit purposes
	DisableDate   time.Time //for audit purposes
	CreatedBy     uuid.UUID //for audit purposes
	Updated       time.Time //for audit purposes
	UpdatedBy     uuid.UUID //for audit purposes
	Name          string    //Asset name
	Description   string    //Asset description
	InventoryCode string    // unique for facility
	AssetType     string    // requires additional table
	FacilityId    uuid.UUID // uuid for owner facility - foreign key to facilityies table uuid
}
