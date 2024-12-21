package main

import (
	"go-tutorial/controllers"
	internal "go-tutorial/internal/database"
	"go-tutorial/internal/middleware"
	"go-tutorial/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := internal.InitDB()

	if db == nil {
		// error while connection to the database
		return
	}

	router.LoadHTMLGlob("templates/*.html")
	notesService := &services.NotesService{}
	notesService.InitService(db)
	notesController := &controllers.NotesController{}
	notesController.InitController(*notesService)
	notesController.InitRoutes(router)

	authService := services.InitAuthService(db)

	authController := controllers.InitAuthController(authService)
	authController.InitRoutes(router)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong from gin",
		})
	})

	router.GET("/", middleware.CheckAuthMiddleware, func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/notes/note-ui")
	})

	router.GET("/home", middleware.CheckAuthMiddleware, func(c *gin.Context) {
		// titles := []string{"Hi", "All", "Bye"}
		// c.HTML(http.StatusOK, "index.html", gin.H{
		// 	"title":   "Home Page",
		// 	"content": "This is a Home Page of Go Gin 2.0",
		// 	"titles":  titles,
		// })

		c.Redirect(http.StatusMovedPermanently, "/notes/note-ui")
	})

	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", gin.H{
			"title":   "About Page",
			"content": "This is a About Page of Go Gin 2.0",
		})
	})

	// router.GET("/me/:id/:newId", func(c *gin.Context) {
	//   var id=c.Param("id")
	//   var newId=c.Param("newId")

	//   c.JSON(http.StatusOK, gin.H{
	//    "user_id":id,
	//    "user_new_id":newId,
	//   })
	// })

	// router.POST("/me",func(c *gin.Context) {
	//   // Email,Password

	//   type MeRequest struct {
	//     Email string `json:"email" binding:"required"`
	//     Password string  `json:"password"`
	//   }

	//   var meRequest MeRequest

	//  if err:= c.BindJSON(&meRequest); err!=nil{
	//   c.JSON(http.StatusBadRequest, gin.H{
	//    "error": err.Error(),
	//    })

	//    return
	//  }

	//   c.JSON(http.StatusOK, gin.H{
	//    "email":meRequest.Email,
	//    "password":meRequest.Password,
	//   })
	// })

	// router.PUT("/me",func(c *gin.Context) {
	//   // Email,Password

	//   type MeRequest struct {
	//     Email string `json:"email" binding:"required"`
	//     Password string  `json:"password"`
	//   }

	//   var meRequest MeRequest

	//  if err:= c.BindJSON(&meRequest); err!=nil{
	//   c.JSON(http.StatusBadRequest, gin.H{
	//    "error": err.Error(),
	//    })

	//    return
	//  }

	//   c.JSON(http.StatusOK, gin.H{
	//    "email":meRequest.Email,
	//    "password":meRequest.Password,
	//   })
	// })

	// router.PATCH("/me",func(c *gin.Context) {
	//   // Email,Password

	//   type MeRequest struct {
	//     Email string `json:"email" binding:"required"`
	//     Password string  `json:"password"`
	//   }

	//   var meRequest MeRequest

	//  if err:= c.BindJSON(&meRequest); err!=nil{
	//   c.JSON(http.StatusBadRequest, gin.H{
	//    "error": err.Error(),
	//    })

	//    return
	//  }

	//   c.JSON(http.StatusOK, gin.H{
	//    "email":meRequest.Email,
	//    "password":meRequest.Password,
	//   })
	// })

	// router.DELETE("/me/:id",func(c *gin.Context) {
	//   var id=c.Param("id")

	//   c.JSON(http.StatusOK, gin.H{
	//    "id":id,
	//    "message":"deleted!!",
	//   })
	// })

	router.Run(":8000")
}
