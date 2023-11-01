package models

type BoxData struct {
	BoxName         string  `json:"box_name"`
	Humidity        float64 `json:"humidity"`
	Ghum            int     `json:"ghum"`
	Temperature     float64 `json:"temperature"`
	CO2             int     `json:"co2"`
	FanState        uint8   `json:"fan_state"`
	HumidifierState uint8   `json:"humidifier_state"`
	PumpState       uint8   `json:"pump_state"`
	Time            string  `json:"time"`
	Phase           string  `json:"phase"`
}
