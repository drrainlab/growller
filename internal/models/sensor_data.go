package models

type SensorData struct {
	BoxName     string  `json:"box_name"`
	Humidity    float64 `json:"humidity"`
	Ghum        int     `json:"ghum"`
	Temperature float64 `json:"temperature"`
	CO2         int     `json:"co2"`
	Time        string  `json:"time"`
	Phase       string  `json:"phase"`
}
