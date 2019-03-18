package config

import (
	"intelliq/app/cachestore"
	"intelliq/app/common"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
)

//CacheConnect connects to redis server
func CacheConnect(router *gin.Engine) {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"localhost": ":6379",
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
		router.Use(addRedisToContext(store))
		router.Use(cachestore.RegisterRequestValidation())
	}
}

func addRedisToContext(cacheStore *cache.Codec) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(common.CACHE_STORE_KEY, cacheStore)
	}
}
