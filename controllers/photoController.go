package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"photomanagerapp/database"
	"photomanagerapp/models"
)

func PhotoShow(c *gin.Context) {
	// Retrieve user information from context
	user, _ := c.Get("user")

	// Define an empty slice to hold photo data
	var photos []models.Photo

	// Find photos based on the authorized user's ID
	database.DB.Where("user_id = ?", user.(models.User).ID).Find(&photos)

	// Respond with a success message and the list of photos for the user
	c.JSON(http.StatusOK, gin.H{
		"photos": photos,
	})
}

func PhotoCreate(c *gin.Context) {
	user, _ := c.Get("user")

	var request struct {
		Title    string
		Caption  string
		PhotoUrl string
	}

	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed request!",
		})
		return
	}

	// Create a new Photo object with data and authorized user's ID
	photo := models.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoUrl: request.PhotoUrl,
		UserID:   user.(models.User).ID,
	}

	// Insert the new photo data into the database
	res := database.DB.Create(&photo)

	// Check for errors during photo creation
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Create photo failed!",
		})
		return
	}

	// Fetch the newly created photo data by its ID
	var createdPhoto models.Photo
	database.DB.First(&createdPhoto, photo.ID)

	// Respond with success message and the created photo information
	c.JSON(http.StatusOK, gin.H{
		"photo": createdPhoto,
	})
}

func PhotoUpdate(c *gin.Context) {
	user, _ := c.Get("user")

	// Get photo ID from request path parameter
	id := c.Param("photoId")

	var request struct {
		Title    string
		Caption  string
		PhotoUrl string
	}

	// Bind request data to the struct
	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed request!",
		})
		return
	}

	// Fetch photo data based on the ID
	var photo models.Photo
	retrieve := database.DB.First(&photo, id)

	// Check if photo exists and matches authorized user's photo
	if retrieve.Error != nil || photo.UserID != user.(models.User).ID {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Photo not found!",
		})
		return
	}

	// Load location time zone
	jakartaLocation, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error loading Jakarta timezone!",
		})
		return
	}

	// Update photo data with new values and formatted Updated_At
	res := database.DB.Model(&photo).Updates(models.Photo{
		Title:      request.Title,
		Caption:    request.Caption,
		PhotoUrl:   request.PhotoUrl,
		Updated_At: time.Now().In(jakartaLocation).Format("2006-01-02 15:04:05"),
	})

	// Check if update operation was successful
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Update photo failed!",
		})
		return
	}

	// 1Fetch the updated photo data
	var updatedPhoto models.Photo
	database.DB.First(&updatedPhoto, photo.ID)

	// Respond with success message and updated photo information
	c.JSON(http.StatusOK, gin.H{
		"photo": updatedPhoto,
	})
}

func PhotoDelete(c *gin.Context) {
	user, _ := c.Get("user")

	id := c.Param("photoId")

	var photo models.Photo
	retrieve := database.DB.First(&photo, id)

	// Check if photo exists and matches authorized user's photo
	if retrieve.Error != nil || photo.UserID != user.(models.User).ID {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Photo not found!",
		})
		return
	}

	// Delete the photo data from the database
	res := database.DB.Delete(&photo)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Delete photo failed!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success! Photo has been deleted",
	})
}
