package models

import (
	"github.com/google/uuid"
)

// Represents the interface to interact with gorm's generic methods
type GoFacilityAsset struct {
	FacilityAssetId uuid.UUID
	Enabled         bool
	Name            string
	Description     string
	InventoryId     string
	AssetType       string
	FacilityId      uuid.UUID
	Facility        Facilities
}

// Object Type: Table, represents DTO
type FacilityAsset struct {
	FacilityAssetId uuid.UUID `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	Enabled         bool      //for audit purposes
	Name            string    //Asset name
	Description     string    //Asset description
	InventoryId     string    // unique for facility - it is the inventory code used for identifying the asset
	AssetType       string    // requires additional table
	FacilityId      uuid.UUID // uuid for owner facility - foreign key to facilityies table uuid
}

// FromEntity respects the gorm_generics.GormModel interface
// Creates new GORM model from Entity.
func (f FacilityAsset) FromEntity(gofacilityasset GoFacilityAsset) interface{} {
	return FacilityAsset{
		FacilityAssetId: gofacilityasset.FacilityAssetId,
		Enabled:         gofacilityasset.Enabled,
		Name:            gofacilityasset.Name,
		Description:     gofacilityasset.Description,
		InventoryId:     gofacilityasset.InventoryId,
		AssetType:       gofacilityasset.AssetType,
		FacilityId:      gofacilityasset.FacilityId,
	}
}

// ToEntity respects the gorm_generics.GormModel interface
// Creates new Entity from GORM model.
func (f FacilityAsset) ToEntity() GoFacilityAsset {
	return GoFacilityAsset{
		FacilityAssetId: f.FacilityAssetId,
		Enabled:         f.Enabled,
		Name:            f.Name,
		Description:     f.Description,
		InventoryId:     f.InventoryId,
		AssetType:       f.AssetType,
		FacilityId:      f.FacilityId,
	}
}
