package main

import (
	"app/controller"
)

func main() {
	router := controller.GetRouter()
	router.Run(":8080")
}
