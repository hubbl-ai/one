package dm

type Product struct {
	Characteristics        []*Characteristics `json:"iata:product:characteristics,omitempty"`        // Charateristics of the product
	IsInItems              []*Item            `json:"iata:product:isInItems,omitempty"`              // Reference to the Items in which the product can be found.
	IsInPieces             []*Piece           `json:"iata:product:isInPieces,omitempty"`             // Reference to the pieces where the product can be found. This needs to be filled in case there is no Item
	Manufacturer           *Company           `json:"iata:product:manufacturer,omitempty"`           // Manufacturing company details and contacts
	OtherIdentifier        []*OtherIdentifier `json:"iata:product:otherIdentifier,omitempty"`        // Other product identifier (e.g. Bar code, UPC, EAN, Amazon)
	CommodityItemNumber    string             `json:"iata:product:commodityItemNumber,omitempty"`    // Indicates the specific commodity on which the rate class code is applied
	HSCode                 string             `json:"iata:product:hsCode,omitempty"`                 // Harmonized Commodity code, refer to hsType used. 6 minimum digits are expected.
	HSCommodityDescription string             `json:"iata:product:hsCommodityDescription,omitempty"` // Commodity description
	HSCommodityName        string             `json:"iata:product:hsCommodityName,omitempty"`        // If no Code provided, name of commodity
	HSType                 string             `json:"iata:product:hsType,omitempty"`                 // Reference identifying the type of standard code to be used for the Commodity Classification (Brussels Tariff Nomenclature, EU Harmonized System Code, UN Standard International Trade Classification). Mandatory if the commodity code is more than 6 digits
	ProductDescription     string             `json:"iata:product:productDescription,omitempty"`     // Product description
	ProductIdentifier      string             `json:"iata:product:productIdentifier,omitempty"`      // Manufacturer's unique product identifier
}

func (o Product) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
