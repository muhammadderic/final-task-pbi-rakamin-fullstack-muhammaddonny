package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"photomanagerapp/database"
	"photomanagerapp/middlewares"
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
	// Define request struct to capture login credentials (email & password)
	var request struct {
		Email    string
		Password string
	}

	// Bind incoming request data to the struct
	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed request",
		})
		return
	}

	// Find user by email in database
	var user models.User
	database.DB.First(&user, "email = ?", request.Email)

	// 4. Check if user exists and password is correct
	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect Email or Password",
		})
		return
	}

	// Create a JWT claim with user ID and expiration time
	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":  user.ID,
		"expired": time.Now().Add(time.Hour * 24).Unix(),
	})

	// 6. Sign the claim with a secret key to generate a JWT token
	token, err := sign.SignedString([]byte(middlewares.SecretKey))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// 7. Set cookie properties for secure storage (SameSite with Lax mode)
	c.SetSameSite(http.SameSiteLaxMode)

	// 8. Set a cookie named "Authorization" with the generated JWT token and appropriate expiration
	c.SetCookie("Authorization", token, 3600*24, "", "", false, true) // Expires in 24 hours with HttpOnly flag

	// 9. Respond with success message and the generated JWT token
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
		"token":   token,
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
