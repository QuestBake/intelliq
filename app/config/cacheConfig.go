package config

import (
	"intelliq/app/common"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack"
)

//CacheConnect connects to redis server
func CacheConnect(router *gin.Engine) {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			common.CACHE_DOMAIN: common.CACHE_PORT,
		},
		//	DB:       10,
		//		Password: "appPwd",
	})
	store := &cache.Codec{
		Redis: ring,
		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}
	if store != nil {
		log.Info("Successfully connected to Redis at ", common.CACHE_DOMAIN, common.CACHE_PORT)
		router.Use(addRedisToContext(store))
	}
}

func addRedisToContext(cacheStore *cache.Codec) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(common.CACHE_STORE_KEY, cacheStore)
	}
}
