package client

type Topic string

const (
	TopicWaybill         Topic = "WAYBILL"
	TopicShipment        Topic = "SHIPMENT"
	TopicBooking         Topic = "BOOKING"
	TopicULD             Topic = "ULD"
	TopicItem            Topic = "ITEM"
	TopicInsurance       Topic = "INSURANCE"
	TopicProduct         Topic = "PRODUCT"
	TopicCharacteristics Topic = "CHARACTERISTICS"
	TopicCarrierProduct  Topic = "CARRIER_PRODUCT"
	TopicPiece           Topic = "PIECE"
)
