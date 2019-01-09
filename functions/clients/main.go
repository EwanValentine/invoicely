package main

import (
	"log"
	"net/http"
	"os"

	"github.com/EwanValentine/invoicely/functions/clients/model"
	"github.com/EwanValentine/invoicely/pkg/datastore"
	helpers "github.com/EwanValentine/invoicely/pkg/http"
	"github.com/aws/aws-lambda-go/lambda"
)

// ClientRepository -
type ClientRepository interface {
	Get(id string) (*model.Client, error)
	List() (*[]model.Client, error)
	Store(client *model.Client) error
}

// Handler -
type Handler struct {
	repository ClientRepository
}

// Store a resource
func (h *Handler) Store(request helpers.Req) (helpers.Res, error) {
	var client *model.Client

	if err := helpers.ParseBody(request, &client); err != nil {
		return helpers.ErrResponse(err, http.StatusBadRequest)
	}

	if err := h.repository.Store(client); err != nil {
		return helpers.ErrResponse(err, http.StatusInternalServerError)
	}

	return helpers.Response(map[string]bool{
		"success": true,
	}, http.StatusCreated)
}

// Get a single resource
func (h *Handler) Get(id string, request helpers.Req) (helpers.Res, error) {
	client, err := h.repository.Get(id)
	if err != nil {
		return helpers.ErrResponse(err, http.StatusNotFound)
	}

	return helpers.Response(map[string]interface{}{
		"client": client,
	}, http.StatusOK)
}

// List resources
func (h *Handler) List(request helpers.Req) (helpers.Res, error) {
	clients, err := h.repository.List()
	if err != nil {
		return helpers.ErrResponse(err, http.StatusNotFound)
	}

	return helpers.Response(map[string]interface{}{
		"clients": clients,
	}, http.StatusOK)
}

func main() {
	// Create a connection to the datastore, in this case, DynamoDB
	conn, err := datastore.CreateConnection(os.Getenv("REGION"))
	if err != nil {
		log.Panic(err)
	}

	// Create a new Dynamodb Table instance
	ddb := datastore.NewDynamoDB(conn, os.Getenv("DB_TABLE"))

	// Create a repository
	repository := model.NewClientRepository(ddb)

	// Create the handler instance, with the repository
	handler := &Handler{repository}

	// Pass the handler into the router
	router := helpers.Router(handler)

	// Start the Lambda process
	lambda.Start(router)
}
