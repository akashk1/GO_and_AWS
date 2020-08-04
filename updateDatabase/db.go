package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func updateTrade(data map[string]interface{}) error {
	item, err := dynamodbattribute.MarshalMap(data)

	if err != nil {
		return err
	}
	log.Println("akash")
	log.Println(item)
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("ApiTable"),
	}

	_, err = db.PutItem(input)
	if err != nil {
		return err
	}

	return nil

}
