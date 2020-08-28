package models

// PushConfiguration : standard PushConfiguration model
type PushConfiguration struct {
	AttempsLimit int      `json:"attempsLimit"`
	Urls         []string `json:"urls"`
}
