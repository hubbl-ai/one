package dm

type ServiceRequest struct {
	Code                      string `json:"iata:serviceRequest:code,omitempty"`                      // Service request code
	ServiceRequestDescription string `json:"iata:serviceRequest:serviceRequestDescription,omitempty"` // Service request description
	StatementText             string `json:"iata:serviceRequest:statementText,omitempty"`             // Service request statement text
	StatementType             string `json:"iata:serviceRequest:statementType,omitempty"`             // Service request statement type - e.g. Dangerous Goods, Lithium Ion Battery, Live Animal Certificate
}

func (o ServiceRequest) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
