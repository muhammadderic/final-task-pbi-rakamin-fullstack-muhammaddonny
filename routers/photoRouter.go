package routers

import (
	"github.com/gin-gonic/gin"

	"photomanagerapp/controllers"
)

func PhotoRouter(router *gin.Engine) {
	photoGroupRouter := router.Group("photos")

	photoGroupRouter.GET("/", controllers.PhotoShow)
	photoGroupRouter.POST("/", controllers.PhotoCreate)
	photoGroupRouter.PUT("/:photoId", controllers.PhotoUpdate)
	photoGroupRouter.DELETE("/:photoId", controllers.PhotoDelete)
}
