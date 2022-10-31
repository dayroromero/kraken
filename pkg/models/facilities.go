package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

// Object Type: Table
//Facilities implements the data structure to characterise facilities on the Kraken application
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
