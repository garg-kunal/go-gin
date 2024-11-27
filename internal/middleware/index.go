package middleware

import (
	"fmt"
	"go-tutorial/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckMiddleware(c *gin.Context) {

	headers := c.GetHeader("Authorization")

	fmt.Println(headers)

	if headers == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Headers not provided",
		})
		return
	}

	token := strings.Split(headers, " ")

	if len(token) < 2 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Token not provided",
		})
		return
	}

	data, err := utils.TokenCheck(token[1])
	fmt.Println(data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Claims not matched!!!",
		})
		return
	}

	c.Next()

}

func CheckAuthMiddleware(c *gin.Context) {

	token, err := c.Request.Cookie("userToken")
	print("token", token)
	if err != nil || token == nil || token.Value == "" {
		if c.Request.URL.Path == "/auth/login" || c.Request.URL.Path == "/auth/register" {
			c.Next()
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/auth/login")
		return
	}

	fmt.Println(token)

	data, err := utils.TokenCheck(token.Value)
	fmt.Println(data)
	if err != nil {
		c.SetSameSite(http.SameSiteNoneMode)
		c.SetCookie("userToken", "", -1, "", "", true, false)

		if c.Request.URL.Path == "/auth/login" || c.Request.URL.Path == "/auth/register" {
			c.Next()
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/auth/login")
		return
	}

	c.Set("isLoggedIn", true)
	c.Next()

}
