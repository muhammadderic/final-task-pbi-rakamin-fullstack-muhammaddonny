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

	// Check if user exists and password is correct
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

	// Sign the claim with a secret key to generate a JWT token
	token, err := sign.SignedString([]byte(middlewares.SecretKey))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	// Set cookie properties for secure storage (SameSite with Lax mode)
	c.SetSameSite(http.SameSiteLaxMode)

	// Set a cookie named "Authorization" with the generated JWT token and appropriate expiration
	c.SetCookie("Authorization", token, 3600*24, "", "", false, true) // Expires in 24 hours with HttpOnly flag

	// Respond with success message and the generated JWT token
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
		"token":   token,
	})
}

func UserLogOut(c *gin.Context) {
	// Invalidate the "Authorization" cookie
	c.SetCookie("Authorization", "", -1, "", "", false, true)

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully!",
	})
}

func UserUpdate(c *gin.Context) {
	// Retrieve user information from context (assuming middleware extracted it)
	user, _ := c.Get("user")

	// Get user ID from request path parameter
	id := c.Param("userId")

	// Define struct to capture update data
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

	// Fetch user data to be updated (assuming based on ID)
	var updateUser models.User
	retrieve := database.DB.First(&updateUser, id)

	// Check if user exists and matches authorized user
	if retrieve.Error != nil || updateUser.ID != user.(models.User).ID {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found!",
		})
		return
	}

	// Load time zone (optional, based on your current location)
	jakartaLocation, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error loading Jakarta timezone!",
		})
		return
	}

	// Update user data with new values and formatted Updated_At
	res := database.DB.Model(&updateUser).Updates(models.User{
		Username:   request.Username,
		Email:      request.Email,
		Password:   string(encrypt),
		Updated_At: time.Now().In(jakartaLocation).Format("2006-01-02 15:04:05"),
	})

	// Check if update operation was successful
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Update user failed!",
		})
		return
	}

	// Fetch the updated user data
	var updatedUser models.User
	database.DB.First(&updatedUser, updateUser.ID)

	// Respond with success message and updated user information
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":         updatedUser.ID,
			"username":   updatedUser.Username,
			"email":      updatedUser.Email,
			"created_at": updatedUser.Created_At,
			"updated_at": updatedUser.Updated_At,
		},
	})
}

func UserDelete(c *gin.Context) {
	// Retrieve user information from context (assuming middleware extracted it)
	user, _ := c.Get("user")

	// Get user ID from request path parameter
	id := c.Param("userId")

	// Fetch user data to be deleted (assuming based on ID)
	var deleteUser models.User
	retrieve := database.DB.First(&deleteUser, id)

	// Check if user exists and matches authorized user
	if retrieve.Error != nil || deleteUser.ID != user.(models.User).ID {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found!",
		})
		return
	}

	// Delete user's photos (assuming a separate Photo model exists)
	var photo models.Photo
	resPhoto := database.DB.Where("user_id = ?", user.(models.User).ID).Delete(&photo)
	if resPhoto.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Delete user's photos failed!",
		})
		return
	}

	// Delete the user data from the database
	res := database.DB.Delete(&deleteUser)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Delete user failed!",
		})
		return
	}

	// Invalidate the "Authorization" cookie to log out the user
	c.SetCookie("Authorization", "", -1, "", "", false, true)

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted and logged out successfully!",
	})
}
