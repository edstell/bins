package store

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/edstell/lambda/errors"
)

type DynamoDB struct {
	db                dynamodbiface.DynamoDBAPI
	timeNow           func() time.Time
	propertyTableName string
}

var _ Store = &DynamoDB{}

func NewDynamoDB(db dynamodbiface.DynamoDBAPI, timeNow func() time.Time) *DynamoDB {
	return &DynamoDB{
		db:                db,
		timeNow:           timeNow,
		propertyTableName: "RecyclingServicesProperty",
	}
}

func (s *DynamoDB) ReadProperty(ctx context.Context, propertyID string) (*Property, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(s.propertyTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"property_id": {
				S: aws.String(propertyID),
			},
		},
	}

	result, err := s.db.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, errors.NotFound(fmt.Sprintf("property: %s", propertyID))
	}

	property := &Property{}
	if err := dynamodbattribute.UnmarshalMap(result.Item, property); err != nil {
		return nil, err
	}

	return property, nil
}

func (s *DynamoDB) WriteProperty(ctx context.Context, property Property) error {
	item, err := dynamodbattribute.MarshalMap(property)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(s.propertyTableName),
		Item:      item,
	}

	if _, err := s.db.PutItem(input); err != nil {
		return err
	}

	return nil
}
