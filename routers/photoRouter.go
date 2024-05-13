package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PhotoRouter(router *gin.Engine) {
	photoGroupRouter := router.Group("photos")

	photoGroupRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Show a photo",
		})
	})
	photoGroupRouter.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Create a photo",
		})
	})
	photoGroupRouter.PUT("/:photoId", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Edit a photo",
		})
	})
	photoGroupRouter.DELETE("/:photoId", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete a photo",
		})
	})
}
