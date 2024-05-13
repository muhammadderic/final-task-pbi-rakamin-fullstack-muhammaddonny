package main

import (
	"fmt"

	"photomanagerapp/routers"

	"github.com/gin-gonic/gin"
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
