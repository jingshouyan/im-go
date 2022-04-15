package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tpl", gin.H{
			"Title": "Users",
		})
	})
	r.Run()
}
