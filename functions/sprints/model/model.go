package model

// Item model
type Item struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
	Ticket      string `json:"ticket"`
	Description string `json:"description"`

	// Duration of time spent on this item, in minutes
	Duration uint32 `json:"duration"`
	Sprint   string `json:"sprint"`
}
