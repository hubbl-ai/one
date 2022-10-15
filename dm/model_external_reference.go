package dm

import "time"

type ExternalReference struct {
	DocumentChecksum string `json:"iata:externalReference:documentChecksum,omitempty"` // Checksum of the document to validate its integrity
	DocumentID       string `json:"iata:externalReference:documentId,omitempty"`       // Unique document identifier
	DocumentLink     string `json:"iata:externalReference:documentLink,omitempty"`     // Link to the document, e.g. URL of the file where it is hosted
	DocumentName     string `json:"iata:externalReference:documentName,omitempty"`     // If no DocumentType provided, name of the referenced document
	DocumentType     string `json:"iata:externalReference:documentType,omitempty"`     // Type of the referenced document . Can refer UNEDIFACT 1001  e.g. 740 - Air Waybill, but not limited to
	DocumentVersion  string `json:"iata:externalReference:documentVersion,omitempty"`  // Document version number

	ExpiryDate         time.Time `json:"iata:externalReference:expiryDate,omitempty"` // Document expiry date
	ValidFrom          time.Time `json:"iata:externalReference:validFrom,omitempty"`  // Document validity start date
	DocumentOriginator *Company  `json:"iata:externalReference:documentOriginator,omitempty"`
	Location           *Location `json:"iata:externalReference:location,omitempty"`
}

func (o ExternalReference) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
