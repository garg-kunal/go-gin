package controllers

import (
	"fmt"
	"go-tutorial/internal/utils"
	"go-tutorial/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func InitAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: *authService,
	}
}

func (a *AuthController) InitRoutes(router *gin.Engine) {
	routes := router.Group("/auth")
	routes.POST("/login", a.Login())
	routes.POST("/register", a.Register())
}

func (*AuthController) Nope() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "connected",
		})
		return
	}
}

func (a *AuthController) Register() gin.HandlerFunc {
	type RegisterBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required,min=8,max=255"`
	}
	return func(c *gin.Context) {
		var registerBody RegisterBody
		if err := c.BindJSON(&registerBody); err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}
		user, err := a.authService.Register(&registerBody.Email, &registerBody.Password)
		if err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": user,
		})
		return
	}
}

func (a *AuthController) Login() gin.HandlerFunc {
	type RegisterBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required,min=8,max=255"`
	}
	return func(c *gin.Context) {
		var registerBody RegisterBody
		if err := c.ShouldBindJSON(&registerBody); err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}
		fmt.Print(registerBody)
		user, err := a.authService.Login(&registerBody.Email, &registerBody.Password)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		var token string
		token, err = utils.GenerateToken(user.Email, user.Id)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": user,
			"token":   token,
		})
		return
	}
}
