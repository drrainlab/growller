package models

type BoxData struct {
	BoxName         string  `json:"box_name"`
	Humidity        float64 `json:"humidity"`
	Ghum            int     `json:"ghum"`
	Temperature     float64 `json:"temperature"`
	CO2             int     `json:"co2"`
	FanState        bool    `json:"fan_state"`
	HumidifierState bool    `json:"humidifier_state"`
	PumpState       bool    `json:"pump_state"`
	Time            string  `json:"time"`
	Phase           string  `json:"phase"`
}
