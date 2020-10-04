package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.POST("/load", load)
	r.GET("/get", get)

	r.Run(":8080")
}
