package model

// Company model
type Company struct {
	CpnCode     string `json:"cpnCode"`
	Certificate struct {
		ID          string `json:"id"`
		ContentType string `json:"contentType"`
		Base64      string `json:"base64"`
		Pass        string `json:"pass"`
	} `json:"certificate"`
}
