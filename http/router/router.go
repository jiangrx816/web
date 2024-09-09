package router

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/jiangrx816/gopkg/log"
)

func All() func(r *gin.Engine) {
	return func(r *gin.Engine) {

		// panic日志
		r.Use(ginzap.RecoveryWithZap(log.Sugar().Desugar(), true))
		r.MaxMultipartMemory = 10 << 20 // 10MB

		prefixRouter := r.Group("/")

		// 默认的Api路由
		Api(prefixRouter)

		// 绘本
		Picture(prefixRouter)

		// Market
		Market(prefixRouter)

		// 古诗库
		Poetry(prefixRouter)

		// 下载库
		Download(prefixRouter)

		// 数学
		Math(prefixRouter)
	}
}
