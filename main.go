package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Coordinate struct {
	X *int `json:"x,omitempty"`
	Y *int `json:"y,omitempty"`
}

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

func clientError(status int, body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       body,
	}, nil
}

func handler(req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var pair Coordinate
	var projection Coordinate

	err := json.Unmarshal([]byte(req.Body), &pair)
	if err != nil {
		return clientError(401, "Couldn't unmarshal json pair")
	}

	if pair.X == nil || pair.Y == nil {
		return clientError(500, "Missing x or y")
	}

	x, y := *pair.X, *pair.Y
	projection.X, projection.Y = &x, &y

	prj, err := json.Marshal(projection)
	if err != nil {
		return clientError(500, "Couldn't marshal json response")
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(prj),
	}, nil
}

func main() {
	lambda.Start(handler)
}
