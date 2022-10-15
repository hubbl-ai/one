package dm

type Person struct {
	ContactType      ContactType          `json:"iata:person:contactType,omitempty"`
	Salutation       string               `json:"iata:person:salutation,omitempty"`       // Salutation
	FirstName        string               `json:"iata:person:firstName,omitempty"`        // First name / given name
	MiddleName       string               `json:"iata:person:middleName,omitempty"`       // Middle name/ other name
	LastName         string               `json:"iata:person:lastName,omitempty"`         // Last name / family name / surname
	JobTitle         string               `json:"iata:person:jobTitle,omitempty"`         // Job title / position
	Department       string               `json:"iata:person:department,omitempty"`       // Department / Division / Unit
	EmployeeID       string               `json:"iata:person:employeeId,omitempty"`       // Employee ID
	AssociatedBranch *CompanyBranch       `json:"iata:person:associatedBranch,omitempty"` // Refers to the Branch the person is associated with
	Contact          []*Contact           `json:"iata:person:contact,omitempty"`          // Contact details
	Documents        []*ExternalReference `json:"iata:person:documents,omitempty"`        // Linked documents to the person, e.g. driver's license, ID, etc.
}

func (o Person) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
