package kraken

//AddressStr implements a reusable structure to handle addresses
type AddressStr struct {
	Street1 string `json:"street_1"`
	Street2 string `json:"street_2"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	ZipCode string `json:"zip_code"`
}

//LatLong implements the geographic latitude/longitude data structure
type LatLong struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Object Type: Table
type Country struct {
	IsoCode     string
	CountryName string
}

// ContactStr stores basic contact information
type ContactStr struct {
	IsPrimary bool
	Name      string
	Email     string
	Phone     string
	Address   *AddressStr
}

// EntityAddressStr stores entyty addresses: shipping, mailing
type EntityAddressStr struct {
	IsPrimary  bool
	IsShipping bool
	Address    *AddressStr
}

// TaxIdStr stores generic Unique Tax information.
// It asusmes the existance of an unique national tax id
type TaxIdStr struct {
	TaxId   string // Unique national tax id
	Country string // 3 Char ISO code
}
