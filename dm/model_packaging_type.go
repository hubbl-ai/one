package dm

type PackagingType struct {
	Piece                    *Piece `json:"iata:packagingType:piece,omitempty"`                    // Piece on which the Packaging type is applicable
	PackagingTypeDescription string `json:"iata:packagingType:packagingTypeDescription,omitempty"` // Free Text. Describes the package type.
	// Packaging type identifier as per UNECE Rec 21 Annex V and
	// VI e.g. 1A - Drum, steel - Packaging material
	// code. Identifies the Logistic Unit package type. UN
	// Recommendation on Transport of Dangerous Goods, Model
	// Regulations
	TypeCode string `json:"iata:packagingType:typeCode,omitempty"`
}

func (o PackagingType) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
