package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

// Represents the interface to interact with gorm's generic methods
type GoFacilities struct {
	FacilityId         uuid.UUID
	Status             bool
	DisableNote        string
	FacilityName       string
	IsPort             bool
	PortId             uuid.UUID
	TypeOfTerminal     string
	ThirdPartyServices bool
	FacilityAddress    pgtype.JSONB
	FacilityCoordinate pgtype.JSONB
	EntityId           uuid.UUID
	FacilityManager    pgtype.JSONB
}

// Object Type: Table, represents DTO
type Facilities struct {
	FacilityId         uuid.UUID    `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	Status             bool         //for audit purposes
	DisableNote        string       //Audit Ppurposes
	FacilityName       string       `json:"facility_name"` // Facility name
	IsPort             bool         // True if the facility is a port
	PortId             uuid.UUID    // PortId- foreign key from port table
	TypeOfTerminal     string       `json:"type_of_terminal"`                           // river, sea, truck, rail
	ThirdPartyServices bool         `json:"third_party_service"`                        // True if the terminal provides public services to third party customers (public terminal)
	FacilityAddress    pgtype.JSONB `json:"facility_address" gorm:"column:metadata"`    // Location
	FacilityCoordinate pgtype.JSONB `json:"facility_coordinate" gorm:"column:metadata"` // Geographic coordinates
	EntityId           uuid.UUID    // Entity id - foreign key from entity table
	FacilityManager    pgtype.JSONB `json:"-" gorm:"column:metadata"`
}

// FromEntity respects the gorm_generics.GormModel interface
// Creates new GORM model from Entity.
func (g Facilities) FromEntity(gofacilities GoFacilities) interface{} {
	return Facilities{
		FacilityId:         gofacilities.FacilityId,
		Status:             gofacilities.Status,
		DisableNote:        gofacilities.DisableNote,
		FacilityName:       gofacilities.FacilityName,
		IsPort:             gofacilities.IsPort,
		PortId:             gofacilities.PortId,
		TypeOfTerminal:     gofacilities.TypeOfTerminal,
		ThirdPartyServices: gofacilities.ThirdPartyServices,
		FacilityAddress:    gofacilities.FacilityAddress,
		FacilityCoordinate: gofacilities.FacilityCoordinate,
		EntityId:           gofacilities.EntityId,
		FacilityManager:    gofacilities.FacilityManager,
	}
}

// ToEntity respects the gorm_generics.GormModel interface
// Creates new Entity from GORM model.
func (f Facilities) ToEntity() GoFacilities {
	return GoFacilities{
		FacilityId:         f.FacilityId,
		Status:             f.Status,
		DisableNote:        f.DisableNote,
		FacilityName:       f.FacilityName,
		IsPort:             f.IsPort,
		PortId:             f.PortId,
		TypeOfTerminal:     f.TypeOfTerminal,
		ThirdPartyServices: f.ThirdPartyServices,
		FacilityAddress:    f.FacilityAddress,
		FacilityCoordinate: f.FacilityCoordinate,
		EntityId:           f.EntityId,
		FacilityManager:    f.FacilityManager,
	}
}
