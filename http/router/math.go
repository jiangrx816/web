package router

import (
	"github.com/gin-gonic/gin"
	"web/http/handler/math_handler"
	"web/middleware"
)

func Math(r *gin.RouterGroup) {

	prefixRouter := r.Group("v1").Use(middleware.GlobalMiddleware())

	mathHandler := math_handler.NewMathHandler()
	{
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/math/compute", mathHandler.ApiMakeCompute)
	}
}
