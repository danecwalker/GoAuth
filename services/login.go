package services

import (
	"GoAuth/utils"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type UserCredentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func Login(body string) (events.APIGatewayProxyResponse, error) {
	var data UserCredentials

	json.Unmarshal([]byte(body), &data)

	if data.Username == "" || data.Password == "" {
		return utils.BuildResponse(utils.RObj{
			"message": "You must enter both a username and password to login",
		}, 401)
	}

	return utils.BuildResponse(utils.RObj{
		"username": data.Username,
		"password": data.Password,
	}, 200)
}