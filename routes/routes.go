package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api.POST("/login", func(c *gin.Context) {
		// TODO: Implement login
	})
	api.GET("/authping", func(c *gin.Context) {
		//TODO: Implement authping
	})

}
