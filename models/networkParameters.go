package models

type Payload struct {
	MessagingProduct string   `json:"messaging_product"`
	To               string   `json:"to"`
	Type             string   `json:"type"`
	Template         Template `json:"template"`
}

type Language struct {
	Code string `json:"code"`
}

type Template struct {
	Name     string   `json:"name"`
	Language Language `json:"language"`
}
