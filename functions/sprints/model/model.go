package model

// Sprint model
type Sprint struct {
	ID          string `json:"id"`
	Description string `json:"description"`

	// StartDate and EndDate are UNIX timestamps
	StartDate int64  `json:"start_date"`
	EndDate   int64  `json:"end_date"`
	Client    string `json:"client"`
}
