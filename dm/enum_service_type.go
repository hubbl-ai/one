package dm

type ServiceType string

const (
	ServiceTypeOther   ServiceType = "OSI" // Other Service Information
	ServiceTypeCode    ServiceType = "SPH" // Special Handling Codes
	ServiceTypeRequest ServiceType = "SSR" // Special Service Request
)
