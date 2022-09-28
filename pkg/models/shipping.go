package kraken

import (
	"time"

	"github.com/google/uuid"
)

// Object Type: Table
type ShippingBOL struct {
	Uuid              uuid.UUID `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	Enabled           bool      //for audit purposes
	CreationDate      time.Time //for audit purposes
	DisableDate       time.Time //for audit purposes
	CreatedBy         uuid.UUID //for audit purposes
	Updated           time.Time //for audit purposes
	UpdatedBy         uuid.UUID //for audit purposes
	ShippingETA       time.Time
	ShippingETD       time.Time
	ArrivalDate       time.Time
	CloseDate         time.Time
	IsNullandVoid     bool      // Used in case of having generated new sets
	OriginalBOL       uuid.UUID // Used to track changes in BOLs (new sets)
	IsDelivered       bool      // product has been delivered
	ShippingLineId    uuid.UUID // foreign key to business partner table with IsShippingAgent True
	VesselId          uuid.UUID // link to vessel data - foreign key to VesselStruct.uuid
	Vouage            string    // vessel voyage
	ProductId         uuid.UUID // Product Id - foreign key
	BusinessPartnerId uuid.UUID // BP Id - foreign key to business partner table with IsCustomer True
	Quantity          float64
	DestinationPortId uuid.UUID // Port uuid - foreign key to port table
}

// Object Type: Table
type VesselStruct struct {
	Uuid        uuid.UUID   `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	VesselType  string      // truck, train, ship
	VesselSpecs interface{} // receives any of the alternative specs
}

type TruckSpecs struct {
	Uuid         uuid.UUID `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	LicensePlate string
}
type ShipSpecs struct {
	Uuid       uuid.UUID `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	IMOCode    string
	MMSICode   string
	VesselName string
}
