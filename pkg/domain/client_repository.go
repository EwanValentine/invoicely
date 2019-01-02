package domain

// Client model
type Client struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rate        int32  `json:"rate"`
}

// Datastore -
type Datastore interface {
	List(castTo interface{}) error
	Get(key string, castTo interface{}) error
	Store(item interface{}) error
}

// NewClientRepository instance
func NewClientRepository(datastore Datastore) *ClientRepository {
	return &ClientRepository{datastore: datastore}
}

// ClientRepository stores and fetches items
type ClientRepository struct {
	datastore Datastore
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
