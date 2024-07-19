package router

import (
	"github.com/gin-gonic/gin"
	"web/http/handler/poetry_handler"
	"web/middleware"
)

func Poetry(r *gin.RouterGroup) {

	prefixRouter := r.Group("v2").Use(middleware.GlobalMiddleware())

	poetryHandler := poetry_handler.NewPoetryHandler()
	{
		//获取朝代列表
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/poem/dynasty/list", poetryHandler.ApiDynastyList)

		//获取作者列表
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/poem/author/list", poetryHandler.ApiAuthorList)

		//获取作者详情
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/poem/author/info", poetryHandler.ApiAuthorInfo)

		//集合类别列表
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/poem/kind/list", poetryHandler.ApiKindList)

		//类别集合列表
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/poem/collection/list", poetryHandler.ApiCollectionList)

		//名句列表
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/poem/quotes/list", poetryHandler.ApiQuotesList)

		//名句详情
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/poem/quotes/info", poetryHandler.ApiQuotesInfo)

		//集合作品列表
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/poem/works/list", poetryHandler.ApiWorksList)

		//作品详情
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/poem/works/info", poetryHandler.ApiWorksInfo)

		//名句搜索
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/poem/search/quotes", poetryHandler.ApiSearchQuotes)

		//古诗搜索
		prefixRouter.Use(middleware.CheckWechatMiddleware()).GET("/poem/search/works", poetryHandler.ApiSearchWorks)
	}
}
