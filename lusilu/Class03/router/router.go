package router

import (
	"github.com/gin-gonic/gin"
	"httpcache/lusilu/Class03/controller/cache"
	"httpcache/lusilu/Class03/middleware"
)

type Router struct {
	controller *cache.Controller
}

func New() *Router {
	return &Router{controller: cache.New()}
}

func (r *Router) Install(mux *gin.Engine) {
	group := mux.Group("/api")
	group.POST("/user/register", r.controller.Register)
	group.POST("/user/login", r.controller.Login)
	group.POST("/user/logout", middleware.Middleware(), r.controller.Logout)
	group.POST("/user/delete", middleware.Middleware(), r.controller.Delete)

}
