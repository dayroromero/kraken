package kraken

import "time"

// Object Type: Table
type ShippingBOL struct {
	uuid              string
	Enabled           bool      //for audit purposes
	CreationDate      time.Time //for audit purposes
	DisableDate       time.Time //for audit purposes
	CreatedBy         string    //for audit purposes
	Updated           time.Time //for audit purposes
	UpdatedBy         string    //for audit purposes
	ShippingETA       time.Time
	ShippingETD       time.Time
	ArrivalDate       time.Time
	CloseDate         time.Time
	IsDelivered       bool   // product has been delivered
	ShippingLineId    string // foreign key to business partner table with IsShippingAgent True
	VesselId          string // link to vessel data - foreign key to VesselStruct.uuid
	Vouage            string // vessel voyage
	ProductId         string // Product Id - foreign key
	BusinessPartnerId string // BP Id - foreign key to business partner table with IsCustomer True
	Quantity          float64
	DestinationPortId string // Port uuid - foreign key to port table
}

// Object Type: Table
type PortStr struct {
	uuid        string
	PortName    string
	IMOCode     string
	Coordinates LatLong // Port location
	FacilityId  string  // id of the corresponding facility (if any) - foreign key to facility table
}

// Object Type: Table
type VesselStruct struct {
	uuid        string
	VesselType  string      // truck, train, ship
	VesselSpecs interface{} // receives any of the alternative specs
}

type TruckSpecs struct {
	uuid         string
	LicensePlate string
}
type ShipSpecs struct {
	uuid       string
	IMOCode    string
	MMSICode   string
	VesselName string
}
