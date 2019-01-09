package model

import (
	"github.com/EwanValentine/invoicely/pkg/datastore"
	uuid "github.com/satori/go.uuid"
)

// SprintRepository -
type SprintRepository struct {
	datastore datastore.Datastore
}

// NewSprintRepository -
func NewSprintRepository(ds datastore.Datastore) *SprintRepository {
	return &SprintRepository{datastore: ds}
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
	id := uuid.NewV4()
	sprint.ID = id.String()
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
