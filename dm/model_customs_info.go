package dm

// NB: The data model adds the type name as a prefix.

type CustomsInfo struct {
	// Customs content code. Refer CXML Code List 1.100, e.g. IST
	// - Security Textual StatementNumber, M - Movement Reference
	// Number
	ContentCode string `json:"iata:customsInfo:customsInfoContentCode,omitempty"`
	CountryCode string `json:"iata:customsInfo:customsInfoCountryCode,omitempty"` // Customs country code. Refer ISO 3166-2
	Note        string `json:"iata:customsInfo:customsInfoNote,omitempty"`        // Free text for customs remarks
	// Customs subject code. Refer CXML Code List 1.19, e.g. IMP
	// for import, EXP for export, AGT for Agent, ISS for The
	// Regulated Agent Issuing the Security Status for rdf:type
	// Consignment etc.  At least one of the three elements
	// (Country Code, Information Identifier or Customs, Security
	// and Regulatory Control Information Identifier) must be
	// completed
	SubjectCode string `json:"iata:customsInfo:customsInfoSubjectCode,omitempty"`
	Information string `json:"iata:customsInfo:customsInformation,omitempty"` // Information for customs submission
	*LogisticsObject
}

func (o CustomsInfo) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
