package utils

import (
	"errors"
	"go-tutorial/internal/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const (
	MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), MinCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserDataFromToken(c *gin.Context) (*model.UserIdentification, error) {
	value, exists := c.Get("User")
	if !exists {
		return nil, errors.New("User not found in context")
	}

	user, ok := value.(*model.UserIdentification)
	println("User ID:", user.Id, ok)
	if !ok {
		return nil, errors.New("Failed to cast to User Identification")
	}

	return user, nil
}
