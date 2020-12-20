package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func testSum(t *testing.T, a, b, sum int) {
	res, err := handler(events.APIGatewayV2HTTPRequest{
		Body:            fmt.Sprintf(`{"x": %d, "y": %d}`, a, b),
		IsBase64Encoded: false,
	})
	if err != nil {
		t.Fatal("Everything should be ok")
	}

	if res.StatusCode != 200 {
		t.Fatal(fmt.Sprintf("StatusCode should be 200, got %d", res.StatusCode))
	}

	var pair Coordinate
	err = json.Unmarshal([]byte(res.Body), &pair)
	if err != nil {
		t.Fatal("Couldn't unmarshal json pair ")
	}

	if (*pair.X + *pair.Y) != sum {
		t.Fatal("Sum should be should be ", sum)
	}

	fmt.Printf("%d + %d = %d\n", a, b, sum)
}

func TestHandler(t *testing.T) {
	t.Run("Successful Request", func(t *testing.T) {
		testSum(t, 0, 0, 0)
		testSum(t, 1, 0, 1)
		testSum(t, 100, 200, 300)
		testSum(t, 500, 400, 900)
		testSum(t, 1000, 2000, 3000)
	})
}
