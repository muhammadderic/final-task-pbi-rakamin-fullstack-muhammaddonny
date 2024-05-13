package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	userGroupRouter := router.Group("users")

	userGroupRouter.POST("/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Register a user",
		})
	})
	userGroupRouter.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Log in a user",
		})
	})
	userGroupRouter.POST("/logout", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Log out a user",
		})
	})
	userGroupRouter.PUT("/:userId", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Update a user",
		})
	})
	userGroupRouter.DELETE("/:userId", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Delete a user",
		})
	})
}
