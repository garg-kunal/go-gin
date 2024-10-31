package utils

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func sanitizeMessage(originalMessage string) string {
	parts := strings.Split(originalMessage, " ")
	if len(parts) > 0 {
		return strings.Join(parts[2:], " ")
	}
	return "Validation failed."
}

func handleValidationError(c *gin.Context, fieldErr validator.FieldError) string {
	fmt.Print(fieldErr)
	sanitizedMessage := sanitizeMessage(fieldErr.Error())

	return sanitizedMessage
}

func ValidateFields(c *gin.Context, err error, fields ...string) string {
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrs {
			for _, field := range fields {
				if fieldErr.Field() == field {
					if err := handleValidationError(c, fieldErr); err != "" {
						return err
					}
					return "Something went wrong. Try again later."
				}
			}
		}
	}
	return ""
}
