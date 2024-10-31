package controllers

import (
	"fmt"
	"go-tutorial/internal/utils"
	"go-tutorial/services"
	"net/http"

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
	routes.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})
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

	type RegisterBodyForm struct {
		Email    string `form:"email" binding:"required,min=8,max=255"`
		Password string `form:"password" binding:"required,min=8,max=255"`
	}
	return func(c *gin.Context) {
		var registerBody RegisterBodyForm
		if err := c.ShouldBind(&registerBody); err != nil {
			errStr := utils.ValidateFields(c, err, "Email", "Password")
			c.HTML(http.StatusOK, "register.html", gin.H{
				"errMessage": errStr,
			})
			return
		}
		_, err := a.authService.Register(&registerBody.Email, &registerBody.Password)
		if err != nil {
			c.HTML(http.StatusOK, "register.html", gin.H{
				"errMessage":  err.Error(),
			})
			return
		}

		c.HTML(http.StatusOK, "index.html", nil)
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
