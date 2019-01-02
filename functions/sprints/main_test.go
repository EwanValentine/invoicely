package main

import (
	"encoding/json"
	"net/http"
	"testing"

	httpdelivery "github.com/EwanValentine/invoicely/pkg/delivery/http"
	"github.com/EwanValentine/invoicely/pkg/domain"
	"github.com/stretchr/testify/assert"
)

type MockSprintRepository struct{}

func (r *MockSprintRepository) Get(id string) (*domain.Sprint, error) {
	return &domain.Sprint{
		Client: "def456",
	}, nil
}

func (r *MockSprintRepository) Store(sprint *domain.Sprint) error {
	return nil
}

func (r *MockSprintRepository) List() (*[]domain.Sprint, error) {
	return &[]domain.Sprint{
		domain.Sprint{},
	}, nil
}

func TestCanStoreSprint(t *testing.T) {
	request := httpdelivery.Req{
		HTTPMethod: "POST",
		Body: `{
			"status": "test sprint",
			"description": "did some work",
			"duration": 60,
			"sprint": "abc123"
		}`,
	}
	h := &Handler{&MockSprintRepository{}}
	response, err := h.Store(request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, response.StatusCode)
}

func TestStoreSprintInvalidJSON(t *testing.T) {
	request := httpdelivery.Req{
		HTTPMethod: "POST",
		Body:       ``,
	}
	h := &Handler{&MockSprintRepository{}}
	response, err := h.Store(request)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestCanFetchSprint(t *testing.T) {
	request := httpdelivery.Req{
		HTTPMethod: "GET",
		PathParameters: map[string]string{
			"id": "abc123",
		},
	}
	h := &Handler{&MockSprintRepository{}}
	response, err := h.Get("abc123", request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	var res map[string]*domain.Sprint
	err = json.Unmarshal([]byte(response.Body), &res)
	assert.NoError(t, err)
	assert.Equal(t, "def456", res["sprint"].Client)
}
