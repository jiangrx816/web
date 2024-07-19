package router

import (
	"github.com/gin-gonic/gin"
	"web/http/handler/api_handler"
	"web/middleware"
)

func Api(r *gin.RouterGroup) {

	prefixRouter := r.Group("v2").Use(middleware.GlobalMiddleware())

	homeHandler := api_handler.NewApiHandler()
	{
		prefixRouter.GET("/home", homeHandler.Index)
	}
}
