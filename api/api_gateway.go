package api

import "github.com/gin-gonic/gin"

func NewGinEngine() *gin.Engine {
	r := gin.Default()

	r.GET("/", index)

	routes := r.Group("/api/v1")
	routes.POST("/box", createBox)
	routes.GET("/box", findBox)

	return r
}
