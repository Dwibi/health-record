package helpers

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type (
	ParamGenerateJWT struct {
		ExpiredInMinute int
		SecretKey       []byte
		UserId          int64
		Nip             int
	}
	ParamsValidateJWT struct {
		Token     string
		SecretKey string
	}

	contextKey string

	Claims struct {
		UserId int64 `json:"userId"`
		Nip    int   `json:"nip"`
		jwt.StandardClaims
	}
)

const UserContextKey contextKey = "user"

func GenerateJwtToken(p *ParamGenerateJWT) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    p.UserId,
		"nip":       p.Nip,
		"expiresAt": time.Now().Add(time.Duration(p.ExpiredInMinute) * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString(p.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, err
}
