package main

import (
	"app/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := controller.GetRouter()
	router.Run(":8080")
}
