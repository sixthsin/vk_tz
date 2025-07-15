package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct {
	Username string
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{
		Secret: secret,
	}
}

func (j *JWT) Create(data JWTData) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": data.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	s, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}
