package main

import (
	"gin_admin/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	core.RunWindowsServer()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	router.Run(":8080")

}
