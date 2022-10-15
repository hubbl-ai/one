package dm

type PartyRole string

const (
	PartyRoleAgent                      PartyRole = "Agent"
	PartyRoleAirline                    PartyRole = "Airline"
	PartyRoleAirportAuthority           PartyRole = "Airport Authority"
	PartyRoleBroker                     PartyRole = "Broker"
	PartyRoleCommissionableAgent        PartyRole = "Commissionable Agent"
	PartyRoleConsignee                  PartyRole = "Consignee"
	PartyRoleCustoms                    PartyRole = "Customs"
	PartyRoleDeclarant                  PartyRole = "Declarant"
	PartyRoleDeconsolidator             PartyRole = "Deconsolidator"
	PartyRoleFreightForwarder           PartyRole = "Freight Forwarder"
	PartyRoleGroundHandlingAgent        PartyRole = "Ground Handling Agent"
	PartyRoleNominatedfreightcompany    PartyRole = "Nominated freight company"
	PartyRoleNotifyParty                PartyRole = "Notify Party"
	PartyRoleOtherParticipantIdentifier PartyRole = "Other Participant Identifier"
	PartyRolePostOffice                 PartyRole = "Post Office"
	PartyRoleShipper                    PartyRole = "Shipper"
	PartyRoleTrucker                    PartyRole = "Trucker"
)
