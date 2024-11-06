package controllers

import (
	"fmt"
	"go-tutorial/internal/middleware"
	"go-tutorial/internal/utils"
	"go-tutorial/services"
	"net/http"
	"time"

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
	routes.Use(middleware.CheckAuthMiddleware)
	routes.POST("/login", a.Login())
	routes.POST("/register", a.Register())
	routes.GET("/register", func(c *gin.Context) {
		value, exists := c.Get("isLoggedIn")
		if exists && value == true {
			c.Redirect(http.StatusMovedPermanently, "/home")
			return
		}
		c.HTML(http.StatusOK, "register.html", nil)
	})
	routes.GET("/login", func(c *gin.Context) {
		value, exists := c.Get("isLoggedIn")
		if exists && value == true {
			c.Redirect(http.StatusMovedPermanently, "/home")
			return
		}
		c.HTML(http.StatusOK, "login.html", nil)
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
				"errMessage": err.Error(),
			})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/auth/login")
	}
}

func (a *AuthController) Login() gin.HandlerFunc {
	type RegisterBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required,min=8,max=255"`
	}
	type RegisterBodyForm struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required,min=8,max=255"`
	}
	return func(c *gin.Context) {
		var registerBody RegisterBodyForm
		if err := c.ShouldBind(&registerBody); err != nil {
			errStr := utils.ValidateFields(c, err, "Email", "Password")
			c.HTML(http.StatusOK, "login.html", gin.H{
				"errMessage": errStr,
			})
			return
		}
		fmt.Print(registerBody)
		user, err := a.authService.Login(&registerBody.Email, &registerBody.Password)
		if err != nil {
			c.HTML(http.StatusBadRequest, "login.html", gin.H{
				"errMessage": err.Error(),
			})
			return
		}

		var token string
		token, err = utils.GenerateToken(user.Email, user.Id)
		if err != nil {
			c.HTML(404, "login.html", gin.H{
				"errMessage": err.Error(),
			})
			return
		}

		c.SetSameSite(http.SameSiteNoneMode)
		c.SetCookie("userToken", token, int(time.Now().Add(time.Hour*24).Unix()), "", "", true, false)
		c.Redirect(http.StatusMovedPermanently, "/home")
	}
}
