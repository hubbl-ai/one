package dm

import "time"

type MovementTimes struct {
	Direction MovementDirection `json:"iata:movementTimes:direction,omitempty"` // Direction to indicate if it's Inbound or Outbound
	// The milestone list still needs to be defined, it includes
	// elements from CXML Code List 1.92 but is not limited to
	// those values, e.g. block-on and block-off times might be
	// added as a comparison to wheels off and touchdown.
	MovementMilestone string `json:"iata:movementTimes:movementMilestone,omitempty"`
	// Timestamp (date and time) of the movement time. If the
	// movement time is recorded asynchronously, the timestamp
	// should reflect the actual time, not when the data was
	// created.
	MovementTimestamp time.Time        `json:"iata:movementTimes:movementTimestamp,omitempty"`
	TimeType          MovementTimeType `json:"iata:movementTimes:timeType,omitempty"` // The type of time can be Actual, Estimated ot Scheduled
}

func (o MovementTimes) MarshalJSON() (buf []byte, err error) {
	return serialize(o)
}
