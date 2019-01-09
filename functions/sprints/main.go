package main

import (
	"log"
	"net/http"
	"os"

	"github.com/EwanValentine/invoicely/functions/sprints/model"
	"github.com/EwanValentine/invoicely/pkg/datastore"
	httpdelivery "github.com/EwanValentine/invoicely/pkg/http"
	"github.com/aws/aws-lambda-go/lambda"
)

// SprintRepository -
type SprintRepository interface {
	Get(id string) (*model.Sprint, error)
	List() (*[]model.Sprint, error)
	Store(sprint *model.Sprint) error
}

// Handler -
type Handler struct {
	repository SprintRepository
}

// Get a single sprint
func (h *Handler) Get(id string, request httpdelivery.Req) (httpdelivery.Res, error) {
	sprint, err := h.repository.Get(id)
	if err != nil {
		return httpdelivery.ErrResponse(err, http.StatusNotFound)
	}
	return httpdelivery.Response(map[string]interface{}{
		"sprint": sprint,
	}, http.StatusOK)
}

// List all sprints
func (h *Handler) List(request httpdelivery.Req) (httpdelivery.Res, error) {
	sprints, err := h.repository.List()
	if err != nil {
		return httpdelivery.ErrResponse(err, http.StatusNotFound)
	}
	return httpdelivery.Response(map[string]interface{}{
		"sprints": sprints,
	}, http.StatusOK)
}

// Store a sprint
func (h *Handler) Store(request httpdelivery.Req) (httpdelivery.Res, error) {
	var sprint *model.Sprint
	if err := httpdelivery.ParseBody(request, &sprint); err != nil {
		return httpdelivery.ErrResponse(err, http.StatusBadRequest)
	}
	if err := h.repository.Store(sprint); err != nil {
		return httpdelivery.ErrResponse(err, http.StatusBadRequest)
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
	repository := model.NewSprintRepository(ddb)
	handler := &Handler{repository}
	router := httpdelivery.Router(handler)
	lambda.Start(router)
}
