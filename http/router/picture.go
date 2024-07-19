package router

import (
	"github.com/gin-gonic/gin"
	"web/http/handler/picture_handler"
	"web/middleware"
)

func Picture(r *gin.RouterGroup) {

	prefixRouter := r.Group("v2").Use(middleware.GlobalMiddleware())

	pictureHandler := picture_handler.NewPictureHandler()
	{
		//展示播放按钮
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/show/play", pictureHandler.ApiShowPlay)
		//获取栏目展示
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/getBookListNav", pictureHandler.ApiBookNavList)
		//汉语绘本
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/chinese/getList", pictureHandler.ApiChineseBookList)
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/chinese/getBookInfo", pictureHandler.ApiChineseBookInfo)
		//英语绘本
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/english/getList", pictureHandler.ApiEnglishBookList)
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/english/getBookInfo", pictureHandler.ApiEnglishBookInfo)
		//古诗绘本
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/poetry/getList", pictureHandler.ApiPoetryBookList)
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/poetry/getBookInfo", pictureHandler.ApiPoetryBookInfo)
		//专辑绘本
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/chinese/getAlbumList", pictureHandler.ApiAlbumBookIndex)
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/chinese/getAlbumListInfo", pictureHandler.ApiAlbumBookList)
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/chinese/getAlbumInfo", pictureHandler.ApiAlbumBookInfo)
	}
}
