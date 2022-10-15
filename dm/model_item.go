package dm

import "time"

type Item struct {
	Dimensions           *Dimensions        `json:"iata:item:dimensions,omitempty"`           // Dimensions of the item
	IsInPiece            *Piece             `json:"iata:item:isInPiece,omitempty"`            // URI of the PIECE that contains the Item
	OtherIdentifiers     []*OtherIdentifier `json:"iata:item:otherIdentifiers,omitempty"`     // Other identifier details
	Product              *Product           `json:"iata:item:product,omitempty"`              // URI of the product
	ProductionCountry    *Country           `json:"iata:item:productionCountry,omitempty"`    // Production country details
	Quantity             *Value             `json:"iata:item:quantity,omitempty"`             // Quantity of the item when applicable, witth associated units of measure
	TargetCountry        *Country           `json:"iata:item:targetCountry,omitempty"`        // Item target country
	UnitPrice            *Value             `json:"iata:item:unitPrice,omitempty"`            // Product price per unit in the base
	Weight               *Value             `json:"iata:item:weight,omitempty"`               // Weight of the item
	BatchNumber          string             `json:"iata:item:batchNumber,omitempty"`          // Production batch number / reference
	LotNumber            *string            `json:"iata:item:lotNumber,omitempty"`            // Production lot number / reference
	ProductExpiryDate    time.Time          `json:"iata:item:productExpiryDate,omitempty"`    // Product expiry date - e.g. for perishables goods or goods with programmed obsolescence
	ProductionDate       time.Time          `json:"iata:item:productionDate,omitempty"`       // Production date
	QuantityForUnitPrice float64            `json:"iata:item:quantityForUnitPrice,omitempty"` // Product quantity for unit price - e.g. 12 (eggs for one USD 1)
}

func (o Item) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
