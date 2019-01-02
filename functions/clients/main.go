package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/EwanValentine/invoicely/pkg/datastore"
	httpdelivery "github.com/EwanValentine/invoicely/pkg/delivery/http"
	"github.com/EwanValentine/invoicely/pkg/domain"
)

// ClientRepository -
type ClientRepository interface {
	Get(id string) (*domain.Client, error)
	List() (*[]domain.Client, error)
	Store(client *domain.Client) error
}

// Handler -
type Handler struct {
	repository ClientRepository
}

// Store a resource
func (h *Handler) Store(request httpdelivery.Req) (httpdelivery.Res, error) {
	var client *domain.Client

	if err := httpdelivery.ParseBody(request, &client); err != nil {
		return httpdelivery.ErrResponse(err, http.StatusBadRequest)
	}

	if err := h.repository.Store(client); err != nil {
		return httpdelivery.ErrResponse(err, http.StatusInternalServerError)
	}

	return httpdelivery.Response(map[string]bool{
		"success": true,
	}, http.StatusCreated)
}

// Get a single resource
func (h *Handler) Get(id string, request httpdelivery.Req) (httpdelivery.Res, error) {
	client, err := h.repository.Get(id)
	if err != nil {
		return httpdelivery.ErrResponse(err, http.StatusNotFound)
	}

	return httpdelivery.Response(map[string]interface{}{
		"client": client,
	}, http.StatusOK)
}

// List resources
func (h *Handler) List(request httpdelivery.Req) (httpdelivery.Res, error) {
	clients, err := h.repository.List()
	if err != nil {
		return httpdelivery.ErrResponse(err, http.StatusNotFound)
	}

	return httpdelivery.Response(map[string]interface{}{
		"clients": clients,
	}, http.StatusOK)
}

func main() {
	conn, err := datastore.CreateConnection(os.Getenv("REGION"))
	if err != nil {
		log.Panic(err)
	}
	ddb := datastore.NewDynamoDB(conn, os.Getenv("DB_TABLE"))
	repository := domain.NewClientRepository(ddb)
	handler := &Handler{repository}
	router := httpdelivery.Router(handler)
	lambda.Start(router)
}
