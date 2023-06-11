package routes

import (
	"github.com/TomyPY/FinTracker/cmd/api/handlers"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rgb *gin.RouterGroup
}

func NewRouter(eng *gin.Engine) Router {
	return &router{
		eng: eng,
	}
}
func (r *router) MapRoutes() {

	r.buildHealthRoutes()

}

func (r *router) buildHealthRoutes() {
	r.eng.GET("/ping", handlers.PingHandler())
}
