package kraken

import "time"

// Object Type: Table
//EntityStr implements he data structure to characterise entities on the Kraken application
type EntityStr struct {
	EntityId     string    `json:"uuid"`
	Enabled      bool      //for audit purposes
	CreationDate time.Time //for audit purposes
	DisableDate  time.Time //for audit purposes
	CreatedBy    string    //for audit purposes
	Updated      time.Time //for audit purposes
	UpdatedBy    string    //for audit purposes
	EntityName   string    // Name of the entity (e.g. andikem Mexico S.A. de C.V)
	FriendlyName string    // Name to be displayed in afriendly way (e.g. Andikem Mexico)
	CountryId    string    // 3-char ISO code - foreign key from country table
	SegmentId    string    // Segment id - foreign key from segment table
}

// Object Type: Table
//FacilityStr implements the data structure to characterise facilities on the Kraken application
type FacilityStr struct {
	FacilityId         string      `json:"uuid"`
	Enabled            bool        //for audit purposes
	CreationDate       time.Time   //for audit purposes
	DisableDate        time.Time   //for audit purposes
	CreatedBy          string      //for audit purposes
	Updated            time.Time   //for audit purposes
	UpdatedBy          string      //for audit purposes
	FacilityName       string      `json:"facility_name"` // Facility name
	IsPort             bool        // True if the facility is a port
	TypeOfTerminal     string      `json:"type_of_terminal"`    // river, sea, truck, rail
	ThirdPartyServices bool        `json:"third_party_service"` // True if the terminal provides public services to third party customers (public terminal)
	FacilityAddress    *AddressStr `json:"facility_address"`    // Location
	FacilityCoord      *LatLong    `json:"facility_coord"`      // Geographic coordinates
	EntityId           string      // Entity id - foreign key from entity table
}

// Object Type: Table
// Stores all tanks which have been documented, whether they are under andikem's control or not
type TankStr struct {
	TkId            string        `json:"uuid"`
	Enabled         bool          //for audit purposes
	CreationDate    time.Time     //for audit purposes
	DisableDate     time.Time     //for audit purposes
	CreatedBy       string        //for audit purposes
	Updated         time.Time     //for audit purposes
	UpdatedBy       string        //for audit purposes
	TkName          string        `json:"tk_name"`         // tank name as it is defined by the facility
	TkDescr         *TankDescrStr `json:"tk_description"`  // jsonb data that holds the technical description of the tank
	TkCalibrationId string        `json:"tk_cal_table_id"` // uuid for the current tank calibration table - foreign key from tank callibration table
	TkCoord         *LatLong      `json:"tk_coord"`        // Geographic location of the tank
	FacilityId      string        `json:"facility_id"`     // uuid for the tank's facility - foreign key from facility table
}

// Standard tank description layout as defined by Anditerminals
type TankDescrStr struct {
	ConstrDate         time.Time
	DesignStandard     string
	WorkingPreasure    float64 // value in atm
	DesignDensity      float64 // value in kg/dm3
	Material           string
	ProductTypes       []string // e.g. Class I, II , IIIA
	TkCapacity         *TkCapStr
	TkGeom             *TkGeomStr
	NominalThickness   map[string]float64
	TkIntCovering      *TkIntCoveringStr
	TkExtCovering      *TkExtCoveringStr
	Instrumentation    *TkLinkedAssetsStr //requires an instrumentation list
	LevelMeasureMethod string
	SafetyMeasures     *TkLinkedAssetsStr //requires an safety measures' list
}

// Object Type: Table
// Stores all tank calibration tables, which serve as base for the estimation of inventory per tank in MT
type TkCalTableStr struct {
	CalibrationId    string
	CalibrationTable []byte //raw json
	enabled          bool
	CreationDate     time.Time
	DisableDate      time.Time
}

// Tank capacity
type TkCapStr struct {
	NominalCap float64 // values in m3
	SafeCap    float64 // values in m3
}

// Tank Geometry
type TkGeomStr struct {
	Diameter float64 // values in m
	Height   float64 // values in m
}

// Tank internal covering
type TkIntCoveringStr struct {
	Bottom string
	Walls  string
	Cap    string
}

// Tank external covering
type TkExtCoveringStr struct {
	Primary   string
	Linking   string
	Finishing string
}

// Asset list related to the tank
type TkLinkedAssetsStr struct {
	LinkedAssets []FacAssetStr
}

// Linked asset
type FacAssetStr struct {
	AssetId string // uuid of the individual asset - foreign key to asset table
	Type    string // Assett type / function
}

// Object Type: Table
// Should store all available data for all berths which have been identified
type BerthStr struct{}
