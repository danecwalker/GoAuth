package utils

import (
	"os"

	"github.com/golang-jwt/jwt"
)

type OTPClaims struct {
	Username string
	Verified bool
	jwt.StandardClaims
}

func GenerateOTPToken(claims OTPClaims) (string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString([]byte(os.Getenv("JWT_OTP_KEY")))
	return ss
}

// func GenerateAccessToken(username string) (string) {

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	ss, _ := token.SignedString([]byte(os.Getenv("JWT_ACCESS_KEY")))
// 	return ss
// }

// func GenerateRefreshToken(username string) (string) {

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	ss, _ := token.SignedString([]byte(os.Getenv("JWT_REFRESH_KEY")))
// 	return ss
// }