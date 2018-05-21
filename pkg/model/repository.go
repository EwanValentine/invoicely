package model

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/satori/go.uuid"
)

// ClientRepository stores and fetches clients
type ClientRepository struct {
	Conn *dynamodb.DynamoDB
}

// Store a new client
func (repository *ClientRepository) Store(client *Client) error {
	id := uuid.NewV4()
	client.ID = id.String()
	av, err := dynamodbattribute.MarshalMap(client)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Clients"),
	}
	_, err = repository.Conn.PutItem(input)
	if err != nil {
		return err
	}
	return err
}

// Fetch a client
func (repository *ClientRepository) Fetch(key string) (*Client, error) {
	var client *Client
	result, err := repository.Conn.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("Metrics"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(key),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	if err := dynamodbattribute.UnmarshalMap(result.Item, &client); err != nil {
		return nil, err
	}
	return client, nil
}
