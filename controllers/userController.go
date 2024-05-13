package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"photomanagerapp/database"
	"photomanagerapp/models"
)

func UserRegister(c *gin.Context) {
	// Define request struct to capture user registration data
	var request struct {
		Username string
		Email    string
		Password string
	}

	// Bind incoming request data to the struct
	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed request!",
		})
		return
	}

	// Validate required fields are not empty
	if request.Username == "" || request.Email == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username, email, or password cannot be empty!",
		})
		return
	}

	// Validate password length
	if len(request.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password's length minimum 6 characters!",
		})
		return
	}

	// Hash the password using bcrypt for secure storage
	encrypt, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to encrypt the password!",
		})
		return
	}

	// Create a new user object with validated and hashed data
	newUser := models.User{Username: request.Username, Email: request.Email, Password: string(encrypt)}

	// Insert the new user data into the database
	res := database.DB.Create(&newUser)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Registration Failed!",
		})
		return
	}

	// Respond with success message on successful registration
	c.JSON(http.StatusOK, gin.H{
		"message": "Registration Successfull!",
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
