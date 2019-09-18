package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liu578101804/chat/web/controllers"
)

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("views/*")
	router.Static("assets", "assets")

	controllers.RegisterController(router.RouterGroup)

	router.Run(":8081")
}