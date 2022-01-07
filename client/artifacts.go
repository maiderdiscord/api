package client

import (
	"fmt"
	"net/http"

	"golang.org/x/xerrors"
)

func (c *Client) GetArtifactDownloadLink(releaseID string) (string, error) {
	url := fmt.Sprintf("%s/accounts/%s/releases/%s/artifact", BaseUrl, c.Account, releaseID)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", xerrors.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		return "", xerrors.Errorf("failed to send: %w", err)
	}

	return resp.Header.Get("Location"), nil
}
