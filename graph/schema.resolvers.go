package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/maiderdiscord/api/config"
	"github.com/maiderdiscord/api/graph/generated"
	"github.com/maiderdiscord/api/graph/model"
	"github.com/maiderdiscord/api/models"
)

func (r *queryResolver) ActivateToken(ctx context.Context, licenseKey string) (string, error) {
	licenseKeyReq := &models.ValidateByLicenseKeyRequest{
		models.ValidateByLicenseKeyRequestMeta{
			Key: licenseKey,
		},
	}
	reqBody, err := json.Marshal(licenseKeyReq)
	if err != nil {
		return "", errInternalServerError
	}

	resp, err := http.Post(
		fmt.Sprintf(
			"%s/accounts/%s/licenses/actions/validate-key",
			config.BaseUrl,
			config.Account,
		),
		config.ContentType,
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return "", errInternalServerError
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errInternalServerError
	}

	licenseKeyData := new(models.ValidateByLicenseKeyResponse)
	if err := json.Unmarshal(respBody, licenseKeyData); err != nil {
		return "", errInternalServerError
	}

	activationTokenReq := &models.GenerateActivationTokenRequest{
		models.GenerateActivationTokenRequestData{
			Type: "tokens",
			Attributes: models.GenerateActivationTokenRequestAttribute{
				MaxActivations: 1,
			},
		},
	}

	reqBody, err = json.Marshal(activationTokenReq)
	if err != nil {
		return "", errInternalServerError
	}

	log.Println(string(reqBody))

	httpReq, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"%s/accounts/%s/licenses/%s/tokens",
			config.BaseUrl,
			config.Account,
			licenseKeyData.Data.ID,
		),
		bytes.NewReader(reqBody),
	)
	if err != nil {
		return "", errInternalServerError
	}

	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.APIKey))
	log.Printf("Request Headers: %+v", httpReq.Header)

	resp, err = http.DefaultClient.Do(httpReq)
	if err != nil {
		return "", errInternalServerError
	}

	defer resp.Body.Close()

	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", errInternalServerError
	}

	log.Println(string(respBody))

	activateTokenData := new(models.GenerateActivationTokenResponse)
	if err := json.Unmarshal(respBody, activateTokenData); err != nil {
		return "", errInternalServerError
	}

	log.Printf("Activate token is generated for key %s", licenseKey)

	return activateTokenData.Data.Attributes.Token, nil
}

func (r *queryResolver) DownloadLinks(ctx context.Context, licenseKey string) (*model.Links, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
