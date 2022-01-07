package client

import (
	"os"

	"github.com/maiderdiscord/api/config"
)

const (
	BaseUrl     = "https://api.keygen.sh/v1"
	ContentType = "application/vnd.api+json"
)

type Client struct {
	Account string
	APIKey  string
}

func New() *Client {
	return &Client{
		Account: config.Account,
		APIKey:  os.Getenv("KEYGEN_API_KEY"),
	}
}
