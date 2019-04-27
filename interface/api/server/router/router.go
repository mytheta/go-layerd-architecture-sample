package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mytheta/go-layerd-architecture-sample/interface/api/server/handler"
)

type Router interface {
	RouterEngine() *gin.Engine
}

type router struct {
	userHandler handler.UserHandler
}

func NewRouter(h handler.UserHandler) Router {
	return &router{h}

}

// RouterEngine はgin.Engineを生成します
func (r *router) RouterEngine() *gin.Engine {

	e := gin.New()

	// Global middleware
	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	// health check
	e.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	api := e.Group("/v1")
	r.apiRouter(api)

	return e
}
