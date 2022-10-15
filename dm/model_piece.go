package dm

type Piece struct {
	ContainedItems           []*Item                 `json:"iata:piece:containedItems,omitempty"`           // Reference to the item(s) contained in the piece
	ContainedPieces          []*Piece                `json:"iata:piece:containedPieces,omitempty"`          // Details of contained piece(s)
	CustomsInfo              *CustomsInfo            `json:"iata:piece:customsInfo,omitempty"`              // Customs details
	Dimensions               *Dimensions             `json:"iata:piece:dimensions,omitempty"`               // Dimensions details
	ExternalReferences       *ExternalReference      `json:"iata:piece:externalReferences,omitempty"`       // Reference documents details
	GrossWeight              *Value                  `json:"iata:piece:grossWeight,omitempty"`              // Weight details
	HandlingInstructions     []*HandlingInstructions `json:"iata:piece:handlingInstructions,omitempty"`     // Links to Handling instructions / service requests of the piece
	OtherIdentifiers         []*OtherIdentifier      `json:"iata:piece:otherIdentifiers,omitempty"`         // Other piece identification ( e.g. Shipping Marks, Seal)
	OtherParty               *Company                `json:"iata:piece:otherParty,omitempty"`               // Other party company details - e.g. the party to be notified
	PackagingType            *PackagingType          `json:"iata:piece:packagingType,omitempty"`            // Packaging details
	Parties                  []*Party                `json:"iata:piece:parties,omitempty"`                  // Other party company details - e.g. the party to be notified
	Product                  *Product                `json:"iata:piece:product,omitempty"`                  // Product of the piece, mandatory when there are no items
	ProductionCountry        *Country                `json:"iata:piece:productionCountry,omitempty"`        // Goods production country, mandatory when there are no Items
	SecurityDeclaration      *SecurityDeclaration    `json:"iata:piece:securityDeclaration,omitempty"`      // Security details of the piece
	SecurityStatus           *SecurityDeclaration    `json:"iata:piece:securityStatus,omitempty"`           // Security details
	ServiceRequest           *ServiceRequest         `json:"iata:piece:serviceRequest,omitempty"`           // Security requests
	Shipment                 *Shipment               `json:"iata:piece:shipment,omitempty"`                 // Shipment on which the piece is assigned to
	Shipper                  *Company                `json:"iata:piece:shipper,omitempty"`                  // Shipper company details - e.g. the party shipping the piece
	TransportMovements       []*TransportMovement    `json:"iata:piece:transportMovements,omitempty"`       // Transport Movements on which the piece is transported
	ULDReference             []*ULD                  `json:"iata:piece:uldReference,omitempty"`             // ULD on which the (virtual) piece has been loaded into - URIs of the ULD
	VolumetricWeight         *VolumetricWeight       `json:"iata:piece:volumetricWeight,omitempty"`         // Volumetric weight details
	Coload                   bool                    `json:"iata:piece:coload,omitempty"`                   // Coload indicator for the pieces (boolean)
	DeclaredValueForCarriage string                  `json:"iata:piece:declaredValueForCarriage,omitempty"` // The value of a shipment declared for carriage purposes , NVD if no value declared
	DeclaredValueForCustoms  string                  `json:"iata:piece:declaredValueForCustoms,omitempty"`  // The value of a shipment declared for customs purposes , NVD if no value declared
	GoodsDescription         *string                 `json:"iata:piece:goodsDescription,omitempty"`         // General goods description
	LoadType                 LoadType                `json:"iata:piece:loadType,omitempty"`                 // Specify how the piece will be delivered (bulk or ULD)
	NVDForCarriage           bool                    `json:"iata:piece:nvdForCarriage,omitempty"`           // When no value is declared for Carriage, this field may be completed with the value TRUE otherwise FALSE
	NVDForCustoms            bool                    `json:"iata:piece:nvdForCustoms,omitempty"`            // When no value is declared for Customs, this field may be completed with the value TRUE otherwise FALSE
	PackageMarkCoded         PackageMarkCoded        `json:"iata:piece:packageMarkCoded,omitempty"`         // Reference identifying how the package is marked. Field is hardcode to \"SSCC-18\", \"UPC\" or \"Other\"
	PackagedeIdentifier      string                  `json:"iata:piece:packagedeIdentifier,omitempty"`      // SSCC-18 code for the value of the package mark, company or bar code, free text, pallet code, etc.
	ShippingMarks            []string                `json:"iata:piece:shippingMarks,omitempty"`            // Shipping marks
	SLAC                     uint32                  `json:"iata:piece:slac,omitempty"`                     // Shipper's Load And Count  ( total contained piece count as provided by shipper)
	Stackable                bool                    `json:"iata:piece:stackable,omitempty"`                // Stackable indicator for the pieces (boolean)
	Turnable                 bool                    `json:"iata:piece:turnable,omitempty"`                 // Turnable indicator for the pieces (boolean)
	UPID                     string                  `json:"iata:piece:upid,omitempty"`                     // Unique Piece Identifier (UPID) of the piece. Refer IATA Recommended Practice 1689
	Events                   []*Event                `json:"iata:piece:events,omitempty"`                   //
}

func (o Piece) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
