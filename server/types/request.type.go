package types

type AuthRequestStruct struct {
	Auth struct {
		Username string `json:"username"`
		APIKey   string `json:"apiKey"`
	} `json:"auth"`
}

type RecordRequestStruct struct {
	Type   string
	Record struct{}
}
