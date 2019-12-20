package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginTest(c *gin.Context)  {
	fmt.Println("调用login方法...")
	username := c.Param("username")
	password := c.Param("password")
	fmt.Println(username)
	fmt.Println(password)
	if username=="admin" && password=="123456"{
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg": "ok",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg": "err",
		})
	}
}
