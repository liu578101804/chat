package controllers

import "github.com/gin-gonic/gin"

func RegisterController(group gin.RouterGroup)  {

	group.GET("/api/chat", NewChatController().WxChat)

}
