package cache

import (
	"github.com/chenyahui/gin-cache/persist"
	rxRedis "github.com/jiangrx816/gopkg/redis"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
	"web/common"
)

func HttpCache() error {
	if !viper.GetBool("http.switch") {
		return nil
	}
	//判断redis是连接成功
	if rxRedis.ClientDefault("web") == nil {
		log.Printf("%+v", "redis server not connect, http-cache failed.")
		return errors.New("redis server not connect, http-cache failed.")
	}
	log.Printf("%+v", "http-cache started.")
	common.GVA_HTTP_CACHE = persist.NewRedisStore(rxRedis.ClientDefault("web"))
	return nil
}
