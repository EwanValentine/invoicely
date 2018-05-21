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
	Fetch(string) (*model.Client, error)
}

type handler struct {
	repository
}

// Response is a wrapper around the api gateway proxy response, which takes
// a interface argument to be marshalled to json and returned, and an error code
func Response(data interface{}, code int) (events.APIGatewayProxyResponse, error) {
	body, _ := json.Marshal(data)
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: code,
	}, nil
}

// Handler is our lambda handler
func (h handler) Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var client *model.Client
	if err := json.Unmarshal([]byte(request.Body), &client); err != nil {
		log.Println(err)
		return helpers.Response(map[string]string{
			"err": err.Error(),
		}, http.StatusInternalServerError)
	}
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
