package api


func (server *Server) InitRotes() {
	server.RouterGroup.POST("/", server.ShortenUrlController)
}
