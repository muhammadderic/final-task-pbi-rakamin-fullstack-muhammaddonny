package routers

import (
	"github.com/gin-gonic/gin"

	"photomanagerapp/controllers"
)

func UserRouter(router *gin.Engine) {
	userGroupRouter := router.Group("users")

	userGroupRouter.POST("/register", controllers.UserRegister)
	userGroupRouter.POST("/login", controllers.UserLogIn)
	userGroupRouter.POST("/logout", controllers.UserLogOut)
	userGroupRouter.PUT("/:userId", controllers.UserUpdate)
	userGroupRouter.DELETE("/:userId", controllers.UserDelete)
}
