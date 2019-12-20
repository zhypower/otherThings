package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadRouters(router *gin.Engine) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Status": 0,
			"data":   "Hello World!",
		})
	})

	//router.POST("/login", controllers.LoginTest)
	router.POST("/login", func(c *gin.Context) {
		fmt.Println("调用login方法...")
		var username string = c.Query("username")
		var password string = c.Query("password")
		if username=="admin" && password=="123456"{
			fmt.Println("返回ok...")
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg": "ok",
			})
		} else {
			fmt.Println("返回err...")
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg": "err",
			})
		}
	})
}
