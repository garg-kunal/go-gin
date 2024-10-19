package middleware

import (
	"go-tutorial/internal/utils"
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)


func CheckMiddleware(c *gin.Context){

	headers:=c.GetHeader("Authorization");

	fmt.Println(headers);

	if headers == ""{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"error":"Headers not provided",
		})
		return;
	}

	token :=strings.Split(headers," ")

	if len(token) <2 {
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"error":"Token not provided",
		})
		return;
	}

	data,err:= utils.TokenCheck(token[1])
    fmt.Println(data)
	if err!=nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"error":"Claims not matched!!!",
		})
		return;
	}

	c.Next();

}