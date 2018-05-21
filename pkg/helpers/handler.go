package helpers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// Response is a wrapper around the api gateway proxy response, which takes
// a interface argument to be marshalled to json and returned, and an error code
func Response(data interface{}, code int) (events.APIGatewayProxyResponse, error) {
	body, _ := json.Marshal(data)
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: code,
	}, nil
}

// ErrResponse returns an error in a specified format
func ErrResponse(err error, code int) (events.APIGatewayProxyResponse, error) {
	data := map[string]string{
		"err": err.Error(),
	}
	body, _ := json.Marshal(data)
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: code,
	}, err
}
