package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {}

func (b BaseController) ResSuccessJson(h gin.H, c *gin.Context) {
	c.Abort()
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"message": "操作成功",
		"data": h,
	})
}

func (b BaseController) ResFailJson(message string, code int, c *gin.Context) {
	c.Abort()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"message": message,
		"data": "",
	})
}