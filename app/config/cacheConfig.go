package config

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack"
)

//CacheConnect connects to redis server
func CacheConnect(router *gin.Engine) {
	cacheDomain := Conf.Get("cache.cache_domain").(string)
	cachePort := Conf.Get("cache.cache_port").(string)
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			cacheDomain: cachePort,
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
		log.Info("Successfully connected to Redis at ", cacheDomain, cachePort)
		router.Use(addRedisToContext(store))
	}
}

func addRedisToContext(cacheStore *cache.Codec) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(Conf.Get("cache.cache_store_key").(string), cacheStore)
	}
}
