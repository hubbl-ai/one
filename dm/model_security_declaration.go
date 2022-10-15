package dm

import "time"

type SecurityDeclaration struct {
	IssuedBy                      *Person            `json:"iata:securityDeclaration:issuedBy,omitempty"`                      // Name of person (or employee ID) who issued the security status
	OtherRegulatedEntity          []*RegulatedEntity `json:"iata:securityDeclaration:otherRegulatedEntity,omitempty"`          // Any other regulated entity that accepts custody of the cargo and accepts the security status originally issued
	Piece                         *Piece             `json:"iata:securityDeclaration:piece,omitempty"`                         // Piece linked to the Security Declaration
	ReceivedFrom                  *RegulatedEntity   `json:"iata:securityDeclaration:receivedFrom,omitempty"`                  // Regulated entity that tendered the consignment
	RegulatedEntityIssuer         *RegulatedEntity   `json:"iata:securityDeclaration:regulatedEntityIssuer,omitempty"`         // Regulated entity issuing the Security Declaration
	AdditionalSecurityInformation string             `json:"iata:securityDeclaration:additionalSecurityInformation,omitempty"` // Any additional information that may be required by an ICAO Member State
	// Exemption code - e.g.
	// SMUS - small undersized shipments MAIL - mail
	// BIOM - bio-medical samples
	// DIPL - diplomatic bags or diplomatic mail
	// LFSM - life-saving materials
	// NUCL - nuclear materials
	// TRNS - transfer or transshipment
	GroundsForExemption   []string        `json:"iata:securityDeclaration:groundsForExemption,omitempty"`
	IssuedOn              time.Time       `json:"iata:securityDeclaration:issuedOn,omitempty"`              // Date and time when the security status was issued
	OtherScreeningMethods []string        `json:"iata:securityDeclaration:otherScreeningMethods,omitempty"` // Other methods used to secure the cargo
	ScreeningMethod       ScreeningMethod `json:"iata:securityDeclaration:screeningMethod,omitempty"`       // Screening methods which have been used to secure the
	SecurityStatus        string          `json:"iata:securityDeclaration:securityStatus,omitempty"`        // Security status indicator (CXML 1.103) - e.g. SPX- Cargo Secure for Passenger and All-Cargo Aircraft
}

func (o SecurityDeclaration) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
