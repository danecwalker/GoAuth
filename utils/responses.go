package utils

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type RObj map[string]string

func BuildResponse(body RObj, statusCode int) (events.APIGatewayProxyResponse, error) {
	jsonStr, _ := json.Marshal(body)

	return events.APIGatewayProxyResponse{Body: string(jsonStr), StatusCode: statusCode, Headers: RObj{
		"Content-Type": "application/json",
		"Access-Control-Allow-Origin": "*",
	}}, nil
}