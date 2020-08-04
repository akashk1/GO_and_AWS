package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetTrade(id string) (*[]map[string]interface{}, error) {

	input := &dynamodb.QueryInput{
		TableName:              aws.String("ApiTable"),
		KeyConditionExpression: aws.String("id = :tId"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":tId": {
				S: aws.String(id),
			},
		},
	}

	result, err := db.Query(input)
	if err != nil {
		return nil, err
	}

	if result.Items == nil {
		return nil, nil
	}
	log.Println("akash")
	log.Println(result)
	log.Println("akash")
	trade := new([]map[string]interface{})
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &trade)
	if err != nil {
		return nil, err
	}
	log.Println(trade)
	log.Println("akash")
	return trade, nil
}
