package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success111...",
		})
	})

	r.Run(":8080")
}
