package providers

import (
	"github.com/gin-gonic/gin"
)

type Server struct{}

func (e *Server) App() *gin.Engine {
	app := gin.New()
	return app
}

func (e *Server) Listen(app *gin.Engine) {
	app.Run()
}
