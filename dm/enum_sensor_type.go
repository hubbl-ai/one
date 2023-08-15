package dm

type SensorType string

const (
	SensorTypeAccelerometer SensorType = "Accelerometer"
	SensorTypeGeolocation   SensorType = "Geolocation"
	SensorTypeHumidity      SensorType = "Humidity"
	SensorTypeLight         SensorType = "Light"
	SensorTypePressure      SensorType = "Pressure"
	SensorTypeThermometer   SensorType = "Thermometer"
	SensorTypeTilt          SensorType = "Tilt"
	SensorTypeVibration     SensorType = "Vibration"
)
