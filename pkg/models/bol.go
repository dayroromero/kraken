package models

import (
	"time"

	"github.com/google/uuid"
)

// Represents the interface to interact with gorm's generic methods
type GoBol struct {
	ShippingBOLId     uuid.UUID
	Status            bool
	DisableNote       string
	ShippingETA       time.Time
	ShippingETD       time.Time
	ArrivalDate       time.Time
	CloseDate         time.Time
	IsNullandVoid     bool
	OriginalBOL       uuid.UUID
	IsDelivered       bool
	ShippingLineId    uuid.UUID
	VesselId          uuid.UUID
	Voyage            string
	ProductId         uuid.UUID
	BPReceiverId      uuid.UUID
	Quantity          float32
	DestinationPortId uuid.UUID
}

// Object Type: Table, represents DTO
type Bol struct {
	ShippingBOLId     uuid.UUID `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	Status            bool      //for audit purposes
	DisableNote       string    //for audit purposes
	ShippingETA       time.Time
	ShippingETD       time.Time
	ArrivalDate       time.Time
	CloseDate         time.Time
	IsNullandVoid     bool      // Used in case of having generated new sets
	OriginalBOL       uuid.UUID // Used to track changes in BOLs (new sets)
	IsDelivered       bool      // product has been delivered
	ShippingLineId    uuid.UUID // foreign key to business partner table with IsShippingAgent True
	VesselId          uuid.UUID // link to vessel data - foreign key to VesselStruct.uuid
	Voyage            string    // vessel voyage
	ProductId         uuid.UUID // Product Id - foreign key
	BPReceiverId      uuid.UUID // BP Id - foreign key to business partner table with IsCustomer True
	Quantity          float32
	DestinationPortId uuid.UUID // Port uuid - foreign key to port table
}

// FromEntity respects the gorm_generics.GormModel interface
// Creates new GORM model from Entity.
func (g Bol) FromEntity(gobol GoBol) interface{} {
	return Bol{
		ShippingBOLId:     gobol.ShippingBOLId,
		Status:            gobol.Status,
		DisableNote:       gobol.DisableNote,
		ShippingETA:       gobol.ShippingETA,
		ShippingETD:       gobol.ShippingETD,
		ArrivalDate:       gobol.ArrivalDate,
		CloseDate:         gobol.CloseDate,
		IsNullandVoid:     gobol.IsNullandVoid,
		OriginalBOL:       gobol.OriginalBOL,
		IsDelivered:       gobol.IsDelivered,
		ShippingLineId:    gobol.ShippingLineId,
		VesselId:          gobol.VesselId,
		Voyage:            gobol.Voyage,
		ProductId:         gobol.ProductId,
		BPReceiverId:      gobol.BPReceiverId,
		Quantity:          gobol.Quantity,
		DestinationPortId: gobol.DestinationPortId,
	}
}

// ToEntity respects the gorm_generics.GormModel interface
// Creates new Entity from GORM model.
func (p Bol) ToEntity() GoBol {
	return GoBol{
		ShippingBOLId:     p.ShippingBOLId,
		Status:            p.Status,
		DisableNote:       p.DisableNote,
		ShippingETA:       p.ShippingETA,
		ShippingETD:       p.ShippingETD,
		ArrivalDate:       p.ArrivalDate,
		CloseDate:         p.CloseDate,
		IsNullandVoid:     p.IsNullandVoid,
		OriginalBOL:       p.OriginalBOL,
		IsDelivered:       p.IsDelivered,
		ShippingLineId:    p.ShippingLineId,
		VesselId:          p.VesselId,
		Voyage:            p.Voyage,
		ProductId:         p.ProductId,
		BPReceiverId:      p.BPReceiverId,
		Quantity:          p.Quantity,
		DestinationPortId: p.DestinationPortId,
	}
}
