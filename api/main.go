package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liu578101804/chat/api/controllers"
)

func main() {

	router := gin.Default()

	controllers.RegisterController(router.RouterGroup)

	router.Run(":8082")
}

