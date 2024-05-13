package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"photomanagerapp/routers"
)

func main() {
	router := gin.Default()

	routers.UserRouter(router)
	routers.PhotoRouter(router)

	if err := router.Run(); err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}
