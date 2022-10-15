package dm

import "time"

type BookingOption struct {
	BookingTimes           *BookingTimes         `json:"iata:bookingOption:bookingTimes,omitempty"`           // booking times details of the Booking Option (proposed or actual)
	Carrier                *Company              `json:"iata:bookingOption:carrier"`                          // Carrier details
	CarrierProductInfo     *CarrierProduct       `json:"iata:bookingOption:carrierProductInfo,omitempty"`     // Reference to the Carrier products included in the offer
	Consignee              *Company              `json:"iata:bookingOption:consignee,omitempty"`              // Consignee details
	FreightForwarder       *Company              `json:"iata:bookingOption:freightForwarder,omitempty"`       // Freight forwarder details
	NotifyParty            *Company              `json:"iata:bookingOption:notifyParty,omitempty"`            // Other parties to be notified of the booking evolution
	Parties                []*Party              `json:"iata:bookingOption:parties,omitempty"`                // Parties involved in the Booking Option (e.g. shipper, forwarder, ...)
	Price                  *Price                `json:"iata:bookingOption:price,omitempty"`                  // Price of the Booking (if different from the offer)
	RequestRef             *BookingOptionRequest `json:"iata:bookingOption:requestRef,omitempty"`             // Reference to the Booking option request
	Routing                *Routing              `json:"iata:bookingOption:routing,omitempty"`                // Routing details of the offer, to be compared with routing preferences of the quote request
	ShipmentDetails        *Shipment             `json:"iata:bookingOption:shipmentDetails,omitempty"`        // Details of the shipement that is to be shipped
	Shipper                *Company              `json:"iata:bookingOption:shipper,omitempty"`                // Shipper details
	WaybillNumber          *Waybill              `json:"iata:bookingOption:waybillNumber,omitempty"`          // House or Master Waybill unique identifier
	BookingStatus          string                `json:"iata:bookingOption:bookingStatus,omitempty"`          // Status of the Booking with regards to the step in the Sales and Booking process: Quoted, Booked (to be confirmed)
	OfferValidFrom         time.Time             `json:"iata:bookingOption:offerValidFrom,omitempty"`         // Starting datetime of availability of the booking option
	OfferValidTo           time.Time             `json:"iata:bookingOption:offerValidTo,omitempty"`           // Ending datetime of availability of the booking option
	RequestMatchInd        bool                  `json:"iata:bookingOption:requestMatchInd,omitempty"`        // Indicates if the offer is a perfect match to the quote request (boolean)
	ShipmentSecurityStatus string                `json:"iata:bookingOption:shipmentSecurityStatus,omitempty"` // Indicate the secruty state of the shipment, screened (SCR) or not (NCR)
	*LogisticsObject
}

func (o BookingOption) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
