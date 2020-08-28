package models

// Topic : standard topic model
type Topic struct {
	ID                string            `json:"id"`
	LastUpdate        string            `json:"lastUpdate,omitempty"`
	PushConfiguration PushConfiguration `json:"pushConfiguration,omitempty"`
}
