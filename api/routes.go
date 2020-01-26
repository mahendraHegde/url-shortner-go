package api

func (server *Server) InitRotes() {
	server.RouterGroup.POST("/", server.ShortenUrl)
	server.RouterGroup.GET("/:url", server.Redirect)
}
