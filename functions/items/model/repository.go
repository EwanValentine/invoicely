package model

import (
	"github.com/EwanValentine/invoicely/pkg/datastore"
	uuid "github.com/satori/go.uuid"
)

// ItemRepository interfaces with items table
type ItemRepository struct {
	datastore datastore.Datastore
}

// NewItemRepository returns a new instance of item repository
func NewItemRepository(ds datastore.Datastore) *ItemRepository {
	return &ItemRepository{datastore: ds}
}

// Get a single item
func (r *ItemRepository) Get(id string) (*Item, error) {
	var item *Item
	if err := r.datastore.Get(id, &item); err != nil {
		return nil, err
	}
	return item, nil
}

// List all items
func (r *ItemRepository) List() (*[]Item, error) {
	var items *[]Item
	if err := r.datastore.List(&items); err != nil {
		return nil, err
	}
	return items, nil
}

// Store a new item
func (r *ItemRepository) Store(item *Item) error {
	id := uuid.NewV4()
	item.ID = id.String()
	return r.datastore.Store(item)
}
