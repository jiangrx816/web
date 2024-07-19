package router

import (
	"github.com/gin-gonic/gin"
	"time"
	"web/cache"
	"web/common"
)

func routerCache(hour int64) (res gin.HandlerFunc) {
	return cache.CacheByRequestURI(common.GVA_HTTP_CACHE, time.Duration(hour)*time.Hour)
}
