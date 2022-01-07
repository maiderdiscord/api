package models

type Relationship struct {
	Links struct {
		Related string `json:"related"`
	} `json:"links"`
	Data struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"data"`
}
