package dm

import "time"

type Waybill struct {
	Booking                       *BookingOption `json:"iata:waybill:booking,omitempty"`                       // Refers to the Booking option
	CarrierDeclarationPlace       *Location      `json:"iata:waybill:carrierDeclarationPlace,omitempty"`       // Location of individual or company involved in the movement of a consignment or Coded representation of a specific airport/city code
	ContainedWaybills             []*Waybill     `json:"iata:waybill:containedWaybills,omitempty"`             // Refers to the Waybill(s) contained
	AccountingInformation         *string        `json:"iata:waybill:accountingInformation,omitempty"`         // Indicates the details of accounting information. Free text e.g. PAYMENT BY CERTIFIED CHEQUE etc.
	CarrierDeclarationDate        time.Time      `json:"iata:waybill:carrierDeclarationDate,omitempty"`        // Date upon which the certification is made by the carrier
	CarrierDeclarationSignature   string         `json:"iata:waybill:carrierDeclarationSignature,omitempty"`   // Contains the authentication of the Carrier
	ConsignorDeclarationSignature string         `json:"iata:waybill:consignorDeclarationSignature,omitempty"` // Name of consignor signatory
	DestinationCharges            float64        `json:"iata:waybill:destinationCharges,omitempty"`            // Charges levied at destination accruing to the last carrier, in destination currency
	DestinationCurrencyCode       string         `json:"iata:waybill:destinationCurrencyCode,omitempty"`       // ISO 3-letter currency code of destination. Refer to ISO 4217
	DestinationCurrencyRate       float64        `json:"iata:waybill:destinationCurrencyRate,omitempty"`       // Conversion rate applied
	OptionalShippingInfo          string         `json:"iata:waybill:optionalShippingInfo,omitempty"`          // The shipper or its Agent may enter the appropriate optional shipping
	OptionalShippingRefNo         string         `json:"iata:waybill:optionalShippingRefNo,omitempty"`         // Optional shipping reference number if any
	OriginCurrency                string         `json:"iata:waybill:originCurrency,omitempty"`                // ISO alpha 3 Code used to indicate the Origin Currency, refer to ISO 4217 currency codes
	WaybillNumber                 string         `json:"iata:waybill:waybillNumber,omitempty"`                 // House or Master Waybill unique identifier
	WaybillPrefix                 string         `json:"iata:waybill:waybillPrefix,omitempty"`                 // Prefix used for the Waybill Number. Refer to IATA Airlines Codes
	WaybillType                   WaybillType    `json:"iata:waybill:waybillType,omitempty"`                   // Type of the Waybill: House, Direct or Master
}

func (o Waybill) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
