package models

import "time"

type ValidateByLicenseScope struct {
	Fingerprint string `json:"fingerprint,omitempty"`
}

type ValidateByLicenseKeyRequest struct {
	Meta ValidateByLicenseKeyRequestMeta `json:"meta"`
}

type ValidateByLicenseKeyRequestMeta struct {
	Key   string                  `json:"key"`
	Scope *ValidateByLicenseScope `json:"scope,omitempty"`
}

type ValidateByLicenseKeyResponse struct {
	Meta ValidateByLicenseKeyResponseMeta `json:"meta"`
	Data ValidateByLicenseKeyResponseData `json:"data"`
}

type ValidateByLicenseKeyResponseMeta struct {
	Ts       time.Time              `json:"ts"`
	Valid    bool                   `json:"valid"`
	Detail   string                 `json:"detail"`
	Constant string                 `json:"constant"`
	Scope    ValidateByLicenseScope `json:"scope"`
}

type ValidateByLicenseKeyResponseData struct {
	ID            string                                    `json:"id"`
	Type          string                                    `json:"type"`
	Links         ValidateByLicenseKeyResponseDataLinks     `json:"links"`
	Attributes    ValidateByLicenseKeyResponseAttributes    `json:"attributes"`
	Relationships ValidateByLicenseKeyResponseRelationships `json:"relationships"`
}

type ValidateByLicenseKeyResponseDataLinks struct {
	Self string `json:"self"`
}

// TODO: interface{} をなんとかする

type ValidateByLicenseKeyResponseAttributes struct {
	Name           string                 `json:"name,omitempty"`
	Key            string                 `json:"key"`
	Expiry         time.Time              `json:"expiry"`
	Status         string                 `json:"status"`
	Uses           int                    `json:"uses"`
	Protected      bool                   `json:"protected"`
	Suspended      bool                   `json:"suspended"`
	Scheme         interface{}            `json:"scheme,omitempty"`
	Encrypted      bool                   `json:"encrypted"`
	Floating       bool                   `json:"floating"`
	Concurrent     bool                   `json:"concurrent"`
	Strict         bool                   `json:"strict"`
	MaxMachines    int                    `json:"maxMachines"`
	MaxCores       int                    `json:"maxCores"`
	MaxUses        int                    `json:"maxUses,omitempty"`
	RequireCheckIn bool                   `json:"requireCheckIn"`
	LastValidated  time.Time              `json:"lastValidated"`
	LastCheckIn    interface{}            `json:"lastCheckIn,omitempty"`
	NextCheckIn    interface{}            `json:"nextCheckIn,omitempty"`
	Metadata       map[string]interface{} `json:"metadata"`
	Created        time.Time              `json:"created"`
	Updated        time.Time              `json:"updated"`
}

type ValidateByLicenseKeyResponseRelationships struct {
	Account      Relationship `json:"account"`
	Product      Relationship `json:"product"`
	Policy       Relationship `json:"policy"`
	User         Relationship `json:"user"`
	Machines     Relationship `json:"machines"`
	Tokens       Relationship `json:"tokens"`
	Entitlements Relationship `json:"entitlements"`
}
