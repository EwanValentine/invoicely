package main

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/EwanValentine/invoicely/functions/sprints/model"
	httpdelivery "github.com/EwanValentine/invoicely/pkg/http"
	"github.com/stretchr/testify/assert"
)

type MockSprintRepository struct{}

func (r *MockSprintRepository) Get(id string) (*model.Sprint, error) {
	return &model.Sprint{
		Client: "def456",
	}, nil
}

func (r *MockSprintRepository) Store(sprint *model.Sprint) error {
	return nil
}

func (r *MockSprintRepository) List() (*[]model.Sprint, error) {
	return &[]model.Sprint{
		model.Sprint{},
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
	var res map[string]*model.Sprint
	err = json.Unmarshal([]byte(response.Body), &res)
	assert.NoError(t, err)
	assert.Equal(t, "def456", res["sprint"].Client)
}
