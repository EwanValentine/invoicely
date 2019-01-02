package domain

import "time"

// Sprint model
type Sprint struct {
	ID        string    `json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Client    string    `json:"sprint"`
}

// SprintRepository -
type SprintRepository struct {
	datastore Datastore
}

// NewSprintRepository -
func NewSprintRepository(datastore Datastore) *SprintRepository {
	return &SprintRepository{datastore: datastore}
}

// Get a single sprint by id
func (r *SprintRepository) Get(id string) (*Sprint, error) {
	var sprint *Sprint
	if err := r.datastore.Get(id, &sprint); err != nil {
		return nil, err
	}
	return sprint, nil
}

// Store a new sprint
func (r *SprintRepository) Store(sprint *Sprint) error {
	return r.datastore.Store(sprint)
}

// List all sprints
func (r *SprintRepository) List() (*[]Sprint, error) {
	var sprints *[]Sprint
	if err := r.datastore.List(&sprints); err != nil {
		return nil, err
	}
	return sprints, nil
}
