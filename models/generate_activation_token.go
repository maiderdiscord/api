package models

import "time"

type GenerateActivationTokenRequest struct {
	Data GenerateActivationTokenRequestData `json:"data"`
}

type GenerateActivationTokenRequestData struct {
	Type       string                                  `json:"type"`
	Attributes GenerateActivationTokenRequestAttribute `json:"attributes"`
}

type GenerateActivationTokenRequestAttribute struct {
	MaxActivations   int        `json:"maxActivations,omitempty"`
	MaxDeactivations int        `json:"maxDeactivations,omitempty"`
	Expiry           *time.Time `json:"expiry,omitempty"`
}

type GenerateActivationTokenResponse struct {
	Data GenerateActivationTokenResponseData `json:"data"`
}

type GenerateActivationTokenResponseData struct {
	ID            string                                      `json:"id"`
	Type          string                                      `json:"type"`
	Attributes    GenerateActivationTokenResponseAttribute    `json:"attributes"`
	Relationships GenerateActivationTokenResponseRelationship `json:"relationships"`
}

type GenerateActivationTokenResponseAttribute struct {
	Kind             string     `json:"kind"`
	Token            string     `json:"token"`
	Expiry           *time.Time `json:"expiry,omitempty"`
	MaxActivations   int        `json:"maxActivations,omitempty"`
	Activations      int        `json:"activations"`
	MaxDeactivations int        `json:"maxDeactivations,omitempty"`
	Deactivations    int        `json:"deactivations"`
	Created          time.Time  `json:"created"`
	Updated          time.Time  `json:"updated"`
}

type GenerateActivationTokenResponseRelationship struct {
	Account Relationship `json:"account"`
	Bearer  Relationship `json:"bearer"`
}
