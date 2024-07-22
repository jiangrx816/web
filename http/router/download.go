package router

import (
	"github.com/gin-gonic/gin"
	"web/http/handler/download_handler"
	"web/middleware"
)

func Download(r *gin.RouterGroup) {

	prefixRouter := r.Group("v2").Use(middleware.GlobalMiddleware())

	downloadHandler := download_handler.NewDownloadHandler()
	{
		//展示播放按钮
		prefixRouter.Use().GET("/download/album/list", downloadHandler.ApiAlbumList)
	}
}
