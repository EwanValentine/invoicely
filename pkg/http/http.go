package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// ResponseError -
type ResponseError map[string]error

// Req is an alias for an api gateway request
type Req events.APIGatewayProxyRequest

// Res is an alis for an api gateway response
type Res events.APIGatewayProxyResponse

// Response is a wrapper around the api gateway proxy response, which takes
// a interface argument to be marshalled to json and returned, and an error code
func Response(data interface{}, code int) (Res, error) {
	body, _ := json.Marshal(data)
	return Res{
		Body:       string(body),
		StatusCode: code,
	}, nil
}

// ErrResponse returns an error in a specified format
func ErrResponse(err error, code int) (Res, error) {
	data := map[string]string{
		"err": err.Error(),
	}
	body, _ := json.Marshal(data)
	return Res{
		Body:       string(body),
		StatusCode: code,
	}, err
}

// RestHandler represents a RESTful Lambda handler
type RestHandler interface {
	Get(id string, request Req) (Res, error)
	Store(request Req) (Res, error)
	List(request Req) (Res, error)
}

// ParseBody takes the body from the request, parses the json to a given struct pointer
func ParseBody(request Req, castTo interface{}) error {
	return json.Unmarshal([]byte(request.Body), &castTo)
}

// RequestHandleFunc is an alias for an api gateway request signature
type RequestHandleFunc func(request Req) (Res, error)

// Router routes restful endpoints to the correct method
// GET without an ID in the path parameters, calls the List method,
// GET with an ID calls the Get method,
// POST calls the Store method.
func Router(h RestHandler) RequestHandleFunc {
	return func(request Req) (Res, error) {
		switch request.HTTPMethod {
		case "GET":
			id := request.PathParameters["id"]
			if id != "" {
				return h.Get(id, request)
			}
			return h.List(request)
		case "POST":
			return h.Store(request)
		default:
			return ErrResponse(errors.New("method not allowed"), http.StatusMethodNotAllowed)
		}
	}
}
