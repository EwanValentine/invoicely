package main

import (
	"net/http"
	"testing"

	"github.com/EwanValentine/invoicely/pkg/model"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

type fakeRepo struct{}

func (repo fakeRepo) Store(*model.Client) error {
	return nil
}
func TestCanStoreClient(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		Body: `{ "name": "test client", "description": "some test", "rate": 40 }`,
	}
	h := &handler{fakeRepo{}}
	response, err := h.Handler(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
}
