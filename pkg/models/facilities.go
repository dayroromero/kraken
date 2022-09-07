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
	FacilityName       string      `json:"facility_name"`
	TypeOfTerminal     string      `json:"type_of_terminal"`
	ThirdPartyServices bool        `json:"third_party_service"`
	FacilityAddress    *AddressStr `json:"facility_address"`
	FacilityCoord      *LatLong    `json:"facility_coord"`
	EntityId           string      // Entity id - foreign key from entity table
}

// Object Type: Table
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
type TkCalTableStr struct {
	CalibrationId    string
	CalibrationTable []byte //raw json
	enabled          bool
	CreationDate     time.Time
	DisableDate      time.Time
}

type TkCapStr struct {
	NominalCap float64 // values in m3
	SafeCap    float64 // values in m3
}

type TkGeomStr struct {
	Diameter float64 // values in m
	Height   float64 // values in m
}

type TkIntCoveringStr struct {
	Bottom string
	Walls  string
	Cap    string
}

type TkExtCoveringStr struct {
	Primary   string
	Linking   string
	Finishing string
}

type TkLinkedAssetsStr struct {
	LinkedAssets []FacAssetStr
}

type FacAssetStr struct {
	AssetId string //uuid
	Type    string
}
