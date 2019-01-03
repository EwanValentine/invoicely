package datastore

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// CreateConnection to dynamodb
func CreateConnection(region string) (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return nil, err
	}
	return dynamodb.New(sess), nil
}

// DynamoDB is a concrete implementation
// to interface with common DynamoDB operations
type DynamoDB struct {
	table string
	conn  *dynamodb.DynamoDB
}

// NewDynamoDB - creates new dynamodb instance
func NewDynamoDB(conn *dynamodb.DynamoDB, table string) *DynamoDB {
	return &DynamoDB{
		conn: conn, table: table,
	}
}

// List gets a collection of resources
func (ddb *DynamoDB) List(castTo interface{}) error {
	results, err := ddb.conn.Scan(&dynamodb.ScanInput{
		TableName: aws.String(ddb.table),
	})
	if err != nil {
		return err
	}
	if err := dynamodbattribute.UnmarshalListOfMaps(results.Items, &castTo); err != nil {
		return err
	}
	return nil
}

// Store a new Item
func (ddb *DynamoDB) Store(item interface{}) error {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(ddb.table),
	}
	_, err = ddb.conn.PutItem(input)
	if err != nil {
		return err
	}
	return err
}

// Get an item
func (ddb *DynamoDB) Get(key string, castTo interface{}) error {
	result, err := ddb.conn.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(ddb.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(key),
			},
		},
	})
	if err != nil {
		return err
	}
	if err := dynamodbattribute.UnmarshalMap(result.Item, &castTo); err != nil {
		return err
	}
	return nil
}
