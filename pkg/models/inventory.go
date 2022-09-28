package kraken

import (
	"time"

	"github.com/google/uuid"
)

// Object type: Table
// TankAssignStr implements the tank assignment data structure that holds the information related to
// the operational link between a tank, a customer and a product, both for inventory tracking, and
// operations.
type TankAssignStr struct {
	Uuid             uuid.UUID `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	Enabled          bool      //for audit purposes
	CreationDate     time.Time //for audit purposes
	DisableDate      time.Time //for audit purposes
	CreatedBy        uuid.UUID //for audit purposes
	Updated          time.Time //for audit purposes
	UpdatedBy        uuid.UUID //for audit purposes
	CustomerId       uuid.UUID
	TkId             uuid.UUID
	ProductId        string
	StartDate        time.Time //Start date of the tank space assignment to any given customer-product combination
	EndDate          time.Time //End date of the tank space assignment to any given customer-product combination
	OpeningInventory float64   //In Metric Tonnes
}

// Object Type: Table
type ProdInOut struct {
	Uuid           uuid.UUID `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	Enabled        bool      //for audit purposes
	CreationDate   time.Time //for audit purposes
	DisableDate    time.Time //for audit purposes
	CreatedBy      uuid.UUID //for audit purposes
	Updated        time.Time //for audit purposes
	UpdatedBy      uuid.UUID //for audit purposes
	MovementType   string    // in, out, gain, loss, transfer
	TkAssignmentId uuid.UUID
	Quantity       float64
	MovementDate   time.Time
	MovementId     uuid.UUID // only used for in, out, and transfer operations - foreign key MovementOperation.uuid
}

// Object Type: Table
type MeasuredInventory struct {
	Uuid           uuid.UUID `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	Enabled        bool      //for audit purposes
	CreationDate   time.Time //for audit purposes
	DisableDate    time.Time //for audit purposes
	CreatedBy      uuid.UUID //for audit purposes
	Updated        time.Time //for audit purposes
	UpdatedBy      uuid.UUID //for audit purposes
	TkAssignmentId uuid.UUID
	Quantity       float64
}

// Object Type: Table
type Inventory struct {
	Uuid           uuid.UUID `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	Period         time.Time
	TkAssignmentId uuid.UUID
	OpeningInv     float64
	ProdIn         float64
	ProdOut        float64
	Gains          float64
	Losses         float64
	ClosingInv     float64
	Closed         bool // if the period is closed, no recalculations are needed for showing updated inventory
}

// Object Type: Table
type MovementOperation struct {
	Uuid              uuid.UUID `gorm:"primaryKey; type:uuid;default:uuid_generate_v4()"`
	Enabled           bool      //for audit purposes
	CreationDate      time.Time //for audit purposes
	DisableDate       time.Time //for audit purposes
	CreatedBy         uuid.UUID //for audit purposes
	Updated           time.Time //for audit purposes
	UpdatedBy         uuid.UUID //for audit purposes
	TkAssignmentId    uuid.UUID // links the product, tank and owner of the product - foreign key
	BusinessPartnerId uuid.UUID // BP to whom, or from who, the inventory is transferred
	Finalized         bool      // Operation has been finalized
	Accrued           bool      // is the record accrued for inventary purposes?
	MovementType      string    // in, out, transfer
	MovementInterface string    // sea, truck, internal, disposal
	VesselId          uuid.UUID // link to vessel - foreign key to Vessel Table uuid
	StartDate         time.Time
	EndDate           time.Time
	MovementRate      float64 // transfer, charge, or discharge rate
	Quantity          float64
	ImportDocId       uuid.UUID // to link the dispatch to the import document if any. For reporting to customs authorities - Foreign key
}
