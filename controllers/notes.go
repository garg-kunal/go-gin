package controllers

import (
	"go-tutorial/internal/middleware"
	"go-tutorial/internal/utils"
	"go-tutorial/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NotesController struct {
	notesService services.NotesService
}

func (n *NotesController) InitController(notesService services.NotesService) *NotesController {
	n.notesService = notesService
	return n
}

func (n *NotesController) InitRoutes(router *gin.Engine) {
	notes := router.Group("/notes")
	notes.Use(middleware.CheckAuthMiddleware)
	notes.GET("/", n.GetNotes())
	notes.GET("/:id", n.GetNote())
	notes.POST("/", n.CreateNotes())
	notes.PUT("/:id", n.UpdateNotes())
	notes.POST("/update-note/:id", n.UpdateNotes())
	notes.DELETE("/:id", n.DeleteNotes())
	notes.GET("/notes", func(c *gin.Context) {
		c.HTML(http.StatusOK, "note.html", nil)
	})
	notes.GET("/note-ui", n.GetNotesUI())
	notes.GET("/update/:id", n.GetNoteUI())
}

func (n *NotesController) GetNotesUI() gin.HandlerFunc {
	return func(c *gin.Context) {

		status := c.Query("status")
		var actualStatus *bool
		if status != "" {
			aS, err := strconv.ParseBool(status)
			actualStatus = &aS
			if err != nil {
				c.HTML(http.StatusOK, "notes-ui.html", gin.H{
					"errMessage": err.Error(),
				})
				return
			}
		}

		notes, err := n.notesService.GetNotesService(actualStatus)

		if err != nil {
			// c.JSON(400,gin.H{
			// 	"message":err.Error(),
			// })
			c.HTML(http.StatusOK, "notes-ui.html", gin.H{
				"errMessage": err.Error(),
			})
			return
		}
		// c.JSON(200,gin.H{
		// 	"notes": notes,
		// })
		c.HTML(http.StatusOK, "notes-ui.html", gin.H{
			"notes": notes,
		})
	}
}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {

		status := c.Query("status")
		var actualStatus *bool
		if status != "" {
			aS, err := strconv.ParseBool(status)
			actualStatus = &aS
			if err != nil {
				c.JSON(400, gin.H{
					"message": err.Error(),
				})
				return
			}
		}

		notes, err := n.notesService.GetNotesService(actualStatus)

		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"notes": notes,
		})
	}
}

func (n *NotesController) CreateNotes() gin.HandlerFunc {
	type NoteBody struct {
		Title  string `json:"title" binding:"required"`
		Status bool   `json:"status"`
	}
	type NoteBodyForm struct {
		Title  string `form:"title" binding:"required"`
		Status bool   `form:"status"`
	}
	return func(c *gin.Context) {
		var noteBody NoteBodyForm
		if err := c.ShouldBind(&noteBody); err != nil {
			errStr := utils.ValidateFields(c, err, "Title", "Status")
			c.HTML(http.StatusOK, "note.html", gin.H{
				"errMessage": errStr,
			})
			return
		}

		_, err := n.notesService.CreateNotesService(noteBody.Title, noteBody.Status)
		if err != nil {
			// c.JSON(404,gin.H{
			// 	"message":err.Error(),
			// })
			c.HTML(http.StatusOK, "note.html", gin.H{
				"errMessage": err.Error(),
			})
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/notes/note-ui")
		// c.JSON(200,gin.H{
		// 	"note":note,
		// })
	}
}

func (n *NotesController) UpdateNotes() gin.HandlerFunc {
	type NoteBody struct {
		Title  string `json:"title" binding:"required"`
		Status bool   `json:"status"`
		Id     int    `json:"id" binding:"required"`
	}
	type NoteBodyForm struct {
		Title  string `form:"title" binding:"required"`
		Status bool   `form:"status"`
	}
	return func(c *gin.Context) {
		var noteBody NoteBodyForm
		id := c.Param("id")
		noteId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.HTML(http.StatusOK, "note-update.html", gin.H{
				"errMessage": err.Error(),
			})
			return
		}

		if err := c.ShouldBind(&noteBody); err != nil {
			c.HTML(http.StatusOK, "note-update.html", gin.H{
				"errMessage": err.Error(),
			})
			return
		}
		_, err = n.notesService.UpdateNotesService(noteBody.Title, noteBody.Status, int(noteId))
		if err != nil {
			c.HTML(http.StatusOK, "note-update.html", gin.H{
				"errMessage": err.Error(),
			})
			return
		}
		// c.JSON(200,gin.H{
		// 	"note":note,
		// })
		c.Redirect(http.StatusMovedPermanently, "/notes/note-ui")
	}
}

func (n *NotesController) DeleteNotes() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		noteId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = n.notesService.DeleteNotesService(noteId)
		if err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Note Deleted successfully!!!",
		})

	}
}

func (n *NotesController) GetNote() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		noteId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}

		note, err := n.notesService.GetNoteService(noteId)
		if err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"note": note,
		})
	}
}

func (n *NotesController) GetNoteUI() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		noteId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.HTML(http.StatusOK, "note-update.html", gin.H{
				"errMessage": err.Error(),
			})
			return
		}

		note, err := n.notesService.GetNoteService(noteId)
		if err != nil {
			c.HTML(http.StatusOK, "note-update.html", gin.H{
				"errMessage": err.Error(),
			})
			return
		}

		c.HTML(http.StatusOK, "note-update.html", gin.H{
			"note": note,
		})
		return
	}
}
