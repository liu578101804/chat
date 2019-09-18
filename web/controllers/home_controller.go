package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct {
	BaseController
}


func (this *HomeController)GetHome(c *gin.Context) {

	c.HTML(http.StatusOK,"index.html",nil)
}

func NewHomeController() *HomeController {
	return &HomeController{}
}