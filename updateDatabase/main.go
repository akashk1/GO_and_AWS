package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var db *dynamodb.DynamoDB

func init() {
	db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-south-1"))

}

type Data struct {
	id   string  `json:"id"`
	rate float64 `json:"rate"`
}
type Trade struct {
	Rates map[string]interface{} `json:"rates"`

	Code int `json:"code"`
}

func handler(ctx context.Context) error {
	resp, err := http.Get("https://www.freeforexapi.com/api/live?pairs=EURUSD")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result Trade
	var rateData map[string]interface{}
	json.Unmarshal(body, &result)
	rateBytes, err := json.Marshal(result.Rates["EURUSD"])
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(rateBytes, &rateData)
	log.Println(result)
	// data := &Data{
	// 	id:   "EURUSD",
	// 	rate: rateData["rate"].(float64),
	// }
	data := map[string]interface{}{
		"id":   "EURUSD",
		"rate": rateData["rate"].(float64),
	}
	log.Println(rateData)
	log.Println(data)
	err = updateTrade(data)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
func main() {
	lambda.Start(handler)

}
