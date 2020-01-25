package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mahendrahegde/url-shortner-golang/api/routes"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "github.com/mahendrahegde/url-shortner-golang/docs" 
)

type Server struct {
}
// @title  API Docs
// @version V1
// @description This is a sample url shortner service.
// @BasePath /v1
func (server *Server) Start() {
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router := gin.Default()
	routerV1 := router.Group("/v1")
	routes.InitRotes(routerV1)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.Run()
}
