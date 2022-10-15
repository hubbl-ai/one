package dm

import "time"

type Shipment struct {
	ContainedPiece           []*Piece             `json:"iata:shipment:containedPiece,omitempty"`           // Details of contained piece(s)
	DeliveryLocation         *Location            `json:"iata:shipment:deliveryLocation,omitempty"`         // Name and UN/LOCODE code of the point or port of departure, shipment or destination, as required under the applicable delivery term
	Dimensions               *Dimensions          `json:"iata:shipment:dimensions,omitempty"`               // Dimensions details
	ExternalReferences       []*ExternalReference `json:"iata:shipment:externalReferences,omitempty"`       // Reference document details
	FreightForwarder         *Company             `json:"iata:shipment:freightForwarder,omitempty"`         // Reference to the freight forwarder
	Insurance                *Insurance           `json:"iata:shipment:insurance,omitempty"`                // Insurance details
	Parties                  []*Party             `json:"iata:shipment:parties,omitempty"`                  // Parties details
	Shipper                  *Company             `json:"iata:shipment:shipper,omitempty"`                  // Reference to the shipper
	TotalGrossWeight         *Value               `json:"iata:shipment:totalGrossWeight,omitempty"`         // Weight details
	VolumetricWeight         *VolumetricWeight    `json:"iata:shipment:volumetricWeight,omitempty"`         // Volumetric weight details
	WaybillNumber            *Waybill             `json:"iata:shipment:waybillNumber,omitempty"`            // Waybill unique identifier (AWB or HWB)
	DeliveryDate             time.Time            `json:"iata:shipment:deliveryDate,omitempty"`             // The date at which the delivery is supposed to happen. Format is YYYYMMDD.
	GoodsDescription         string               `json:"iata:shipment:goodsDescription,omitempty"`         // General goods description
	INCOTerms                string               `json:"iata:shipment:incoterms,omitempty"`                // Standard codes as defined by UNCEFACT and ICC that correspond to international rules for the interpretation of the most commonly used trade terms in different countries. UNECE recommendation n. 5 Incoterms 2000.
	OtherChargesIndicator    string               `json:"iata:shipment:otherChargesIndicator,omitempty"`    // payment of Other Charges will be made at origin (prepaid) or at destination (collect) C or P
	TotalPieceCount          uint32               `json:"iata:shipment:totalPieceCount,omitempty"`          // Total Piece Count
	TotalSLAC                uint32               `json:"iata:shipment:totalSLAC,omitempty"`                // Total SLAC of all piece groups
	WeightValuationIndicator string               `json:"iata:shipment:weightValuationIndicator,omitempty"` // payment for the Weight/Valuation will be made at origin (prepaid) or at destination (collect)
}

func (o Shipment) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
