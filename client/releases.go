package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/maiderdiscord/api/models"
	"golang.org/x/xerrors"
)

type GetReleasesFilter struct {
	Product string
	Limit   int
}

func (c *Client) GetReleases(filter *GetReleasesFilter) (*models.ListReleasesResponse, error) {
	u, err := url.Parse(fmt.Sprintf("%s/accounts/%s/releases", BaseUrl, c.Account))
	if err != nil {
		return nil, xerrors.Errorf("failed to parse URL: %w", err)
	}

	if filter != nil {
		query := u.Query()

		if product := filter.Product; product != "" {
			query.Add("product", product)
		}

		if limit := filter.Limit; limit != 0 {
			query.Add("limit", strconv.Itoa(limit))
		}

		u.RawQuery = query.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, xerrors.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, xerrors.Errorf("failed to get: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, xerrors.Errorf("failed to read body: %w", err)
	}

	data := new(models.ListReleasesResponse)

	if err := json.Unmarshal(body, data); err != nil {
		return nil, xerrors.Errorf("failed to parse body: %w", err)
	}

	return data, nil
}
