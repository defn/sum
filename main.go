package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

type Coordinate struct {
	X *int `json:"x,omitempty"`
	Y *int `json:"y,omitempty"`
}

type Answer struct {
	Sum int `json:"sum"`
}

func clientError(status int, body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       body,
	}, nil
}

func handler(req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var pair Coordinate
	var sum Answer

	err := json.Unmarshal([]byte(req.Body), &pair)
	if err != nil {
		return clientError(401, "Couldn't unmarshal json pair")
	}

	if pair.X == nil || pair.Y == nil {
		return clientError(500, "Missing x or y")
	}

	sum.Sum = *pair.X + *pair.Y

	sm, err := json.Marshal(sum)
	if err != nil {
		return clientError(500, "Couldn't marshal json response")
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(sm),
	}, nil
}

func main() {
	lambda.Start(handler)
}
