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

type ContactStr struct {
	Name    string
	Email   string
	Phone   string
	Address *AddressStr
}

type TaxIdStr struct {
	TaxId   string // Unique national tax id
	Country string // 3 Char ISO code
}
