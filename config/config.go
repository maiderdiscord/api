package config

import "os"

const (
	BaseUrl     = "https://api.keygen.sh/v1"
	Account     = "54de58e6-e0a0-46dd-8e7a-c4b310c2ebfc"
	ContentType = "application/vnd.api+json"
)

var APIKey = os.Getenv("KEYGEN_API_KEY")
