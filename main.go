package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	if err := router.Run(); err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}
}
