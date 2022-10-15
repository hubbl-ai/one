package dm

type HandlingInstructions struct {
	// Refers to the type of handling information provided:
	// Special Service Request (SSR), Special Handling Codes (SPH)
	// or Other Service Information (OSI)
	ServiceType string `json:"iata:handlinginstructions:serviceType,omitempty"`
	// Service Type code linked to the Service Refers to Code List
	// 1.14 or 1.16 if service type is Special Handling.
	ServiceTypeCode []string `json:"iata:handlinginstructions:serviceTypeCode,omitempty"`
	// Free text description of the handling instructions
	ServiceDescription []string `json:"iata:handlinginstructions:serviceDescription,omitempty"`
	// Refers to the person that requests the handling/service
	RequestedBy *Person `json:"iata:handlinginstructions:requestedBy,omitempty"`
}

func (o HandlingInstructions) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
