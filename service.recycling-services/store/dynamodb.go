package store

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/edstell/lambda/service.recycling-services/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *DynamoDB) ReadProperty(ctx context.Context, propertyID string) (*domain.Property, error) {
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
		return nil, status.Errorf(codes.NotFound, "'%s' was not found", propertyID)
	}

	property := &domain.Property{}
	if err := dynamodbattribute.UnmarshalMap(result.Item, property); err != nil {
		return nil, err
	}

	return property, nil
}

func (s *DynamoDB) WriteProperty(ctx context.Context, propertyID string, services []domain.Service) (*domain.Property, error) {
	property := &domain.Property{
		ID:        propertyID,
		Services:  services,
		UpdatedAt: s.timeNow(),
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
