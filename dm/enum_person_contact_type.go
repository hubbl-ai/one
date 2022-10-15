package dm

type PersonContactType string

const (
	PersonContactTypeCustomer  PersonContactType = "Customer contact"
	PersonContactTypeCustoms   PersonContactType = "Customs contact"
	PersonContactTypeEmergency PersonContactType = "Emergency contact"
	PersonContactTypeOther     PersonContactType = "Other"
)
