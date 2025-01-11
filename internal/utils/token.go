package utils

import (
	"errors"
	"go-tutorial/internal/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secret = "supersecret"

func GenerateToken(email string, id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"nbf":   time.Date(2023, 01, 01, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(secret))

	return tokenStr, err
}

func parseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}

func TokenCheck(jwtToken string) (*model.UserIdentification, error) {
	token, err := parseToken(jwtToken)
	if err != nil {
		return nil, err
	}

	data, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		return nil, errors.New("unable to map claims")
	}

	userData := &model.UserIdentification{
		Email: data["email"].(string),
		Id:    int(data["id"].(float64)),
		Nbf:   data["nbf"].(float64),
	}
	return userData, nil
}
