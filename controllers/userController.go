package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Register a user",
	})
}

func UserLogIn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Log in a user",
	})
}

func UserLogOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Log out a user",
	})
}

func UserUpdate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update a user",
	})
}

func UserDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete a user",
	})
}
