package model

// Client model
type Client struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rate        int32  `json:"rate"`
}
