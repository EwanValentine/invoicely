package main

import (
	"net/http"
	"testing"

	"github.com/EwanValentine/invoicely/pkg/model"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

type fakeRepo struct{}

func (repo *fakeRepo) Fetch(key string) (*model.Client, error) {
	return &model.Client{
		ID:          "abc123",
		Name:        "test client",
		Description: "Test client",
		Rate:        40,
	}, nil
}

func TestCanFetchClient(t *testing.T) {
	request := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{"id": "123"},
	}
	h := &handler{&fakeRepo{}}
	response, err := h.Handler(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
}
