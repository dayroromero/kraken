package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type Tank struct {
	TankId          uuid.UUID
	Status          bool
	DisableNote     string
	TankName        string
	TankDescription pgtype.JSONB
	TankCoordinate  pgtype.JSONB
	FacilityId      uuid.UUID
}

type GoTank struct {
	TankId          uuid.UUID    `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	Status          bool         //for audit purposes
	DisableNote     string       //for audit purposes
	TankName        string       `json:"tk_name"`                                 // tank name as it is defined by the facility
	TankDescription pgtype.JSONB `json:"tank_description" gorm:"column:metadata"` // jsonb data that holds the technical description of the tank
	TankCoordinate  pgtype.JSONB `json:"tank_coordinate" gorm:"column:metadata"`  // uuid for the tank's facility - foreign key from facility table
	FacilityId      uuid.UUID    // FacilityId - foreign key Facilities table
}

// FromEntity respects the gorm_generics.GormModel interface
// Creates new GORM model from Entity.
func (t Tank) FromEntity(gotank GoTank) interface{} {
	return Tank{
		TankId:          gotank.TankId,
		Status:          gotank.Status,
		DisableNote:     gotank.DisableNote,
		TankName:        gotank.TankName,
		TankDescription: gotank.TankDescription,
		TankCoordinate:  gotank.TankCoordinate,
		FacilityId:      gotank.FacilityId,
	}
}

// ToEntity respects the gorm_generics.GormModel interface
// Creates new Entity from GORM model.
func (t Tank) ToEntity() GoTank {
	return GoTank{
		TankId:          t.TankId,
		Status:          t.Status,
		DisableNote:     t.DisableNote,
		TankName:        t.TankName,
		TankDescription: t.TankDescription,
		TankCoordinate:  t.TankCoordinate,
		FacilityId:      t.FacilityId,
	}
}
