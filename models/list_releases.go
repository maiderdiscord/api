package models

import "time"

type ListReleasesResponse struct {
	Data []ListReleasesData `json:"data"`
}

type ListReleasesData struct {
	ID            string                    `json:"id"`
	Type          string                    `json:"type"`
	Attributes    ListReleasesAttributes    `json:"attributes"`
	Relationships ListReleasesRelationships `json:"relationships"`
	Links         ListReleasesLinks         `json:"links"`
}

type ListReleasesAttributes struct {
	Name        string               `json:"name"`
	Description interface{}          `json:"description"`
	Signature   string               `json:"signature"`
	Checksum    string               `json:"checksum"`
	Filename    string               `json:"filename"`
	Filetype    string               `json:"filetype"`
	Filesize    int                  `json:"filesize"`
	Platform    string               `json:"platform"`
	Channel     string               `json:"channel"`
	Status      string               `json:"status"`
	Downloads   int                  `json:"downloads"`
	Upgrades    int                  `json:"upgrades"`
	Version     string               `json:"version"`
	Semver      ListReleasesSemver   `json:"semver"`
	Metadata    ListReleasesMetadata `json:"metadata"`
	Created     time.Time            `json:"created"`
	Updated     time.Time            `json:"updated"`
	Yanked      interface{}          `json:"yanked"`
}

type ListReleasesRelationships struct {
	Account     Relationship `json:"account"`
	Product     Relationship `json:"product"`
	Constraints Relationship `json:"constraints"`
	Artifact    Relationship `json:"artifact"`
}

type ListReleasesLinks struct {
	Self string `json:"self"`
}

type ListReleasesSemver struct {
	Major      int         `json:"major"`
	Minor      int         `json:"minor"`
	Patch      int         `json:"patch"`
	Prerelease interface{} `json:"prerelease"`
	Build      string      `json:"build"`
}

type ListReleasesMetadata struct {
	Sha512 string `json:"sha512"`
}
