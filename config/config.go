package config

import "os"

const (
	BaseUrl           = "https://api.keygen.sh/v1"
	Account           = "54de58e6-e0a0-46dd-8e7a-c4b310c2ebfc"
	ContentType       = "application/vnd.api+json"
	ProductMaider     = "785c9307-19a0-49b7-8992-caf236e307d0"
	ProductHWIDIssuer = "ea0afda5-fe63-4105-8756-88d5260d5cf5"
)

var APIKey = os.Getenv("KEYGEN_API_KEY")
