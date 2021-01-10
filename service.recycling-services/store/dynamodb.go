package store

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/edstell/lambda/libraries/errors"
	recyclingservicesproto "github.com/edstell/lambda/service.recycling-services/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		propertyTableName: "recyclingservicesprotoProperty",
	}
}

func (s *DynamoDB) ReadProperty(ctx context.Context, propertyID string) (*recyclingservicesproto.Property, error) {
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

	property := &recyclingservicesproto.Property{}
	if err := dynamodbattribute.UnmarshalMap(result.Item, property); err != nil {
		return nil, err
	}

	return property, nil
}

func (s *DynamoDB) WriteProperty(ctx context.Context, propertyID string, services []*recyclingservicesproto.Service) (*recyclingservicesproto.Property, error) {
	property := &recyclingservicesproto.Property{
		Id:        propertyID,
		Services:  services,
		UpdatedAt: timestamppb.New(s.timeNow()),
	}
	item, err := dynamodbattribute.MarshalMap(property)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(s.propertyTableName),
		Item:      item,
	}

	if _, err := s.db.PutItem(input); err != nil {
		return nil, err
	}

	return property, nil
}
