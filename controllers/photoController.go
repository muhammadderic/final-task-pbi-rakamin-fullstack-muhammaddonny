package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PhotoShow(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Show a photo",
	})
}

func PhotoCreate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create a photo",
	})
}

func PhotoEdit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Edit a photo",
	})
}

func PhotoDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete a photo",
	})
}
