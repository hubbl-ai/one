package dm

type ULD struct {
	ExternalReference  []*ExternalReference `json:"iata:uld:externalReference,omitempty"`  // Reference documents details
	TareWeight         *Value               `json:"iata:uld:tareWeight,omitempty"`         // Tare weight of the empty ULD
	TransportMovements []*TransportMovement `json:"iata:uld:transportMovements,omitempty"` // Transport Movements on which the ULD are transported
	UPID               []*Piece             `json:"iata:uld:upid,omitempty"`               // Details of contained (virtual) piece(s)
	ATADesignator      string               `json:"iata:uld:ataDesignator,omitempty"`      // US / ATA Unit Load Device type code e.g. M2
	DamageFlag         bool                 `json:"iata:uld:damageFlag,omitempty"`         // Indicates if the ULD is Damaged
	DemurrageCode      string               `json:"iata:uld:demurrageCode,omitempty"`      // Contains three designator of demurrage code, refer to RP 1654 (BCC, HHH, XXX, ZZZ)
	LoadingIndicator   string               `json:"iata:uld:loadingIndicator,omitempty"`   // ULD height or loading limitation code. Refer CXML Code List 1.47,  e.g. R - ULD Height above 244 centimetres
	NBDoor             uint32               `json:"iata:uld:nbDoor,omitempty"`             // Number of doors
	NBFittings         uint32               `json:"iata:uld:nbFittings,omitempty"`         // Number of fittings
	NBNets             uint32               `json:"iata:uld:nbNets,omitempty"`             // Number of nets
	NBStraps           uint32               `json:"iata:uld:nbStraps,omitempty"`           // Number of straps
	ODLNCode           string               `json:"iata:uld:odlnCode,omitempty"`           // Contains two designator codes of ODLN or Operational Damage Limit Notices. ODLN code is used to define type of damage after visually check the serviceability of ULDs section 7, Standard Specifications 40/3 or 40/4 in ULD Regulations
	OwnerCode          string               `json:"iata:uld:ownerCode,omitempty"`          // Owner code of the ULD in aa, an or na format - owner can be an airline or leasing company
	SerialNumber       string               `json:"iata:uld:serialNumber,omitempty"`       // ULD serial number
	ServiceabilityCode ServiceabilityCode   `json:"iata:uld:serviceabilityCode,omitempty"` // Designator of serviceablity condition e.g. SER or DAM
	ULDRemarks         string               `json:"iata:uld:uldRemarks,omitempty"`         // Remarks or Supplement Information
}

func (o ULD) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
