package main

import (
	"GoAuth/services"
	"GoAuth/utils"
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const HEALTH_PATH = "/health"
const LOGIN_PATH = "/login"
const REGISTER_PATH = "/register"
const VERIFY_PATH = "/verify"
const REFRESH_PATH = "/refresh"

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	method := req.HTTPMethod
	path := req.Path

	switch true {
	case method == "GET" && path == HEALTH_PATH:
        return utils.BuildResponse(utils.RObj{"message": "healthy"}, 200)
	case method == "POST" && path == LOGIN_PATH:
        return services.Login(req.Body)
	case method == "POST" && path == REGISTER_PATH:
        return services.Register(req.Body)
	case method == "POST" && path == VERIFY_PATH:
        return events.APIGatewayProxyResponse{Body: "Verify", StatusCode: 200}, nil
	case method == "POST" && path == REFRESH_PATH:
        return events.APIGatewayProxyResponse{Body: "Refresh", StatusCode: 200}, nil
	default:
        return utils.BuildResponse(utils.RObj{
			"message": "Page Not Found",
		}, 404)
	}
}

func main() {
	lambda.Start(HandleRequest)
}