package types

type Weather struct {
	Temperature string `json:"temperature"`
	Condition   string `json:"condition"`
	FeelsLike   string `json:"feels_like"`
	Pressure    string `json:"pressure"`
	Humidity    string `json:"humidity"`
	Wind        string `json:"wind"`
	UpdatedAt   string `json:"updated_at"`
}
