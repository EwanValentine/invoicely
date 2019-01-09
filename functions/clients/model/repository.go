package model

import (
	"github.com/EwanValentine/invoicely/pkg/datastore"
	uuid "github.com/satori/go.uuid"
)

// NewClientRepository instance
func NewClientRepository(ds datastore.Datastore) *ClientRepository {
	return &ClientRepository{datastore: ds}
}

// ClientRepository stores and fetches items
type ClientRepository struct {
	datastore datastore.Datastore
}

// Get a single client
func (r *ClientRepository) Get(id string) (*Client, error) {
	var client *Client
	if err := r.datastore.Get(id, &client); err != nil {
		return nil, err
	}
	return client, nil
}

// Store a new client
func (r *ClientRepository) Store(client *Client) error {
	id := uuid.NewV4()
	client.ID = id.String()
	return r.datastore.Store(client)
}

// List all clients
func (r *ClientRepository) List() (*[]Client, error) {
	var clients *[]Client
	if err := r.datastore.List(&clients); err != nil {
		return nil, err
	}
	return clients, nil
}
