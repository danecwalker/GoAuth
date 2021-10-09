package services

import (
	"GoAuth/utils"
	"crypto/rand"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/golang-jwt/jwt"
)

type UserInfo struct {
	IsVerified bool `json:"is_verified,omitempty"`
	Username string `json:"username,omitempty"`
	Phone string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
}

func Register(body string) (events.APIGatewayProxyResponse, error) {
	var data UserInfo

	json.Unmarshal([]byte(body), &data)

	data.Username = strings.ToLower(data.Username)

	if data.Username == "" || data.Phone == "" || data.Password == "" {
		return utils.BuildResponse(utils.RObj{
			"message": "You must enter a username, phone number and password to login",
		}, 401)
	}

	dsvc := utils.DynamoDB()

	

	result, err := dsvc.Query(&dynamodb.QueryInput{
		TableName: aws.String("GoAuth_Users"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":uname": {
				S: aws.String(data.Username),
			},
		},
		KeyConditionExpression: aws.String("username = :uname"),
	})

	if err != nil {
		log.Fatalf("Got error calling Query: %s", err)
	}


	if len(result.Items) > 0 {
		return utils.BuildResponse(utils.RObj{
			"message": "User with this username already exists",
		}, 406)
	}

	sresult, serr := dsvc.Scan(&dynamodb.ScanInput{
		TableName: aws.String("GoAuth_Users"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":ph": {
				S: aws.String(data.Phone),
			},
		},
		FilterExpression: aws.String("phone = :ph"),
	})

	if serr != nil {
		log.Fatalf("Got error calling Query: %s", err)
	}

	if len(sresult.Items) > 0 {
		return utils.BuildResponse(utils.RObj{
			"message": "User with this phone number already exists",
		}, 406)
	}


	itemToPut := UserInfo{
		Username: data.Username,
		Phone: data.Phone,
		Password: utils.HashPass(data.Password),
		IsVerified: false,
	}

	av, _ := dynamodbattribute.MarshalMap(itemToPut)

	dsvc.PutItem(&dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String("GoAuth_Users"),
	})

	return GenerateOTP(itemToPut.Username, itemToPut.IsVerified)
}

	

func GenerateOTP(uname string, isVerified bool) (events.APIGatewayProxyResponse, error) {
	const length = 6
	const otpChars = "0123456789"
	buffer := make([]byte, length)
    rand.Read(buffer)

    otpCharsLength := len(otpChars)
    for i := 0; i < length; i++ {
        buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
    }

    otp := string(buffer)

	claims := utils.OTPClaims{
		Username: uname,
		Verified: isVerified,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer: "Test",
		},
	}


	return utils.BuildResponse(utils.RObj{
		"username": uname,
		"verified": strconv.FormatBool(isVerified),
		"otp_code": otp,
		"otp_token": utils.GenerateOTPToken(claims),
	}, 200)
}