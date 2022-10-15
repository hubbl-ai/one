package dm

type MovementTimeType string

const (
	MovementTimeTypeActual    MovementTimeType = "Actual"
	MovementTimeTypeDeparted  MovementTimeType = "Departed"
	MovementTimeTypeScheduled MovementTimeType = "Scheduled"
)
