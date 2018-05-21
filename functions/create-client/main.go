package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/EwanValentine/invoicely/pkg/datastore"
	"github.com/EwanValentine/invoicely/pkg/helpers"
	"github.com/EwanValentine/invoicely/pkg/model"
)

type repository interface {
	Store(*model.Client) error
}

type handler struct {
	repository
}

// Handler is our lambda handler
func (h handler) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var client *model.Client
	// Marshal request bodt into our client model
	if err := json.Unmarshal([]byte(request.Body), &client); err != nil {
		return helpers.ErrResponse(err, http.StatusInternalServerError)
	}

	// Call our repository and store our client
	if err := h.repository.Store(client); err != nil {
		return helpers.ErrResponse(err, http.StatusInternalServerError)
	}

	// Return a success response
	return helpers.Response(map[string]bool{
		"success": true,
	}, http.StatusCreated)
}

func main() {
	conn, err := datastore.CreateConnection(os.Getenv("REGION"))
	if err != nil {
		log.Panic(err)
	}
	repository := &model.ClientRepository{Conn: conn}
	h := handler{repository}
	lambda.Start(h.Handler)
}
