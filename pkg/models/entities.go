package models

import "github.com/google/uuid"

// Represents the interface to interact with gorm's generic methods
type GoEntity struct {
	EntityId     uuid.UUID
	Status       bool
	DisableNote  string
	EntityName   string
	FriendlyName string
	CountryId    uuid.UUID
	SegmentId    uuid.UUID
}

// Object Type: Table, represents DTO
type Entity struct {
	EntityId     uuid.UUID `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	Status       bool      //for audit purposes
	DisableNote  string    //Audit purposes
	EntityName   string    // Name of the entity (e.g. andikem Mexico S.A. de C.V)
	FriendlyName string    // Name to be displayed in afriendly way (e.g. Andikem Mexico)
	CountryId    uuid.UUID // 3-char ISO code - foreign key from country table
	SegmentId    uuid.UUID // Segment id - foreign key from segment table
}

// FromEntity respects the gorm_generics.GormModel interface
// Creates new GORM model from Entity.
func (e Entity) FromEntity(goEntity GoEntity) interface{} {
	return Entity{
		EntityId:     goEntity.EntityId,
		Status:       goEntity.Status,
		DisableNote:  goEntity.DisableNote,
		EntityName:   goEntity.EntityName,
		FriendlyName: goEntity.FriendlyName,
		CountryId:    goEntity.CountryId,
		SegmentId:    goEntity.SegmentId,
	}
}

// ToEntity respects the gorm_generics.GormModel interface
// Creates new Entity from GORM model.
func (e Entity) ToEntity() GoEntity {
	return GoEntity{
		EntityId:     e.EntityId,
		Status:       e.Status,
		DisableNote:  e.DisableNote,
		EntityName:   e.EntityName,
		FriendlyName: e.FriendlyName,
		CountryId:    e.CountryId,
		SegmentId:    e.SegmentId,
	}
}
