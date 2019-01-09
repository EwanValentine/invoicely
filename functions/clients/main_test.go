package main

import (
	"net/http"
	"testing"

	"github.com/EwanValentine/invoicely/functions/clients/model"
	httpdelivery "github.com/EwanValentine/invoicely/pkg/http"
	"github.com/stretchr/testify/assert"
)

type MockClientRepository struct{}

func (r *MockClientRepository) Get(id string) (*model.Client, error) {
	return &model.Client{
		ID:          "123",
		Name:        "some client",
		Rate:        40,
		Description: "Some client!",
	}, nil
}

func (r *MockClientRepository) Store(*model.Client) error {
	return nil
}

func (r *MockClientRepository) List() (*[]model.Client, error) {
	return &[]model.Client{
		model.Client{
			ID:          "123",
			Name:        "some client",
			Rate:        40,
			Description: "Some client!",
		},
	}, nil
}

func TestCanFetchClient(t *testing.T) {
	request := httpdelivery.Req{
		HTTPMethod:     "GET",
		PathParameters: map[string]string{"id": "123"},
	}
	h := &Handler{&MockClientRepository{}}
	router := httpdelivery.Router(h)
	response, err := router(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestCanCreateClient(t *testing.T) {
	request := httpdelivery.Req{
		HTTPMethod: "POST",
		Body:       `{ "name": "test client", "description": "some test", "rate": 40 }`,
	}
	h := &Handler{&MockClientRepository{}}
	router := httpdelivery.Router(h)
	response, err := router(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
}

func TestCanListClients(t *testing.T) {
	request := httpdelivery.Req{
		HTTPMethod: "GET",
	}
	h := &Handler{&MockClientRepository{}}
	router := httpdelivery.Router(h)
	response, err := router(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestHandleInvalidJSON(t *testing.T) {
	request := httpdelivery.Req{
		HTTPMethod: "POST",
		Body:       "",
	}
	h := &Handler{&MockClientRepository{}}
	router := httpdelivery.Router(h)
	response, err := router(request)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}
