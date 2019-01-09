package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/EwanValentine/invoicely/functions/items/model"
	"github.com/EwanValentine/invoicely/pkg/datastore"
	httpdelivery "github.com/EwanValentine/invoicely/pkg/http"
)

// ItemRepository -
type ItemRepository interface {
	Store(item *model.Item) error
	Get(id string) (*model.Item, error)
	List() (*[]model.Item, error)
}

// Handler -
type Handler struct {
	repository ItemRepository
}

// Get a single item
func (h *Handler) Get(id string, request httpdelivery.Req) (httpdelivery.Res, error) {
	item, err := h.repository.Get(id)
	if err != nil {
		return httpdelivery.ErrResponse(err, http.StatusNotFound)
	}
	return httpdelivery.Response(map[string]interface{}{
		"item": item,
	}, http.StatusOK)
}

// List all items
func (h *Handler) List(request httpdelivery.Req) (httpdelivery.Res, error) {
	items, err := h.repository.List()
	if err != nil {
		return httpdelivery.ErrResponse(err, http.StatusNotFound)
	}
	return httpdelivery.Response(map[string]interface{}{
		"items": items,
	}, http.StatusOK)
}

// Store a new item
func (h *Handler) Store(request httpdelivery.Req) (httpdelivery.Res, error) {
	var item *model.Item
	if err := httpdelivery.ParseBody(request, &item); err != nil {
		return httpdelivery.ErrResponse(err, http.StatusBadRequest)
	}

	if err := h.repository.Store(item); err != nil {
		return httpdelivery.ErrResponse(err, http.StatusInternalServerError)
	}

	return httpdelivery.Response(map[string]bool{
		"created": true,
	}, http.StatusCreated)
}

func main() {
	conn, err := datastore.CreateConnection(os.Getenv("REGION"))
	if err != nil {
		log.Panic(err)
	}
	ddb := datastore.NewDynamoDB(conn, os.Getenv("DB_TABLE"))
	repository := model.NewItemRepository(ddb)
	handler := &Handler{repository}
	router := httpdelivery.Router(handler)
	lambda.Start(router)
}
