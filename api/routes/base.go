package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mahendrahegde/url-shortner-golang/api/controllers"
)

func InitRotes(router *gin.RouterGroup) {
	router.POST("/", controllers.ShortenUrl)
}
