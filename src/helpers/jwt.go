package helpers

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type (
	ParamGenerateJWT struct {
		ExpiredInMinute int
		SecretKey       string
		UserId          int64
	}
	ParamsValidateJWT struct {
		Token     string
		SecretKey string
	}
)

func GenerateJwtToken(p *ParamGenerateJWT) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(int(p.UserId)),
		"expiresAt": time.Now().Add(time.Duration(p.ExpiredInMinute) * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString(p.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, err
}
