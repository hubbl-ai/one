package dm

type ContactType string

const (
	ContactTypeAlternateEmailAddress ContactType = "Alternate email address"
	ContactTypeAlternatePhoneNumber  ContactType = "Alternate phone number"
	ContactTypeEmailAddress          ContactType = "Email address"
	ContactTypeFaxNumber             ContactType = "Fax number"
	ContactTypeOther                 ContactType = "Other"
	ContactTypePhoneNumber           ContactType = "Phone number"
	ContactTypeTelex                 ContactType = "Telex"
	ContactTypeWebsite               ContactType = "Website"
)
