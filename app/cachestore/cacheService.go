package cachestore

import (
	"fmt"
	"intelliq/app/common"
	utility "intelliq/app/common"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache"
)

var skipURIS = []string{"login", "forgot"}

//RegisterRequestValidation validates for valid sesionId for each request
func RegisterRequestValidation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestURL := ctx.Request.RequestURI
		for _, uri := range skipURIS {
			if strings.Contains(requestURL, uri) {
				ctx.Next()
				return
			}
		}
		sessionID := ctx.Request.Header.Get(common.REQUEST_SESSION_ID_KEY)
		if len(sessionID) == 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden,
				common.GetErrorResponse(common.MSG_USER_SESSION_ERROR))
			return
		}
		store, _ := ctx.Get(common.CACHE_STORE_KEY)
		if store == nil {
			panic("NO REDIS INSTANCE RUNNING...")
		} else {
			cacheStore := store.(*cache.Codec)
			isSessionOK := cacheStore.Exists(sessionID)
			if !isSessionOK {
				ctx.AbortWithStatusJSON(http.StatusForbidden,
					common.GetErrorResponse(common.MSG_USER_AUTH_ERROR))
				return
			}
		}
		ctx.Next()
		if !strings.Contains(requestURL, "logout") {
			RefreshCache(ctx, sessionID)
		}
	}
}

//SetCache sets value in cache
func SetCache(ctx *gin.Context, key string, val interface{}, minutes int) {
	store, _ := ctx.Get(common.CACHE_STORE_KEY)
	if store == nil {
		panic("NO REDIS INSTANCE RUNNING...")
	} else {
		cacheStore := store.(*cache.Codec)
		cacheStore.Set(&cache.Item{
			Key:        key,
			Object:     val,
			Expiration: time.Duration(minutes) * time.Minute,
		})
	}
}

//GetCache gets value from cache
func GetCache(ctx *gin.Context, key string) interface{} {
	store, _ := ctx.Get(common.CACHE_STORE_KEY)
	if store == nil {
		panic("NO REDIS INSTANCE RUNNING...")
	} else {
		cacheStore := store.(*cache.Codec)
		var value interface{}
		err := cacheStore.Get(key, &value)
		if err != nil {
			fmt.Println("KEY DOESNT EXISITS : ", key)
			return nil
		}
		return value
	}
}

//CheckCache checks for key in cache
func CheckCache(ctx *gin.Context, key string) bool {
	store, _ := ctx.Get(common.CACHE_STORE_KEY)
	if store == nil {
		panic("NO REDIS INSTANCE RUNNING...")
	} else {
		cacheStore := store.(*cache.Codec)
		return cacheStore.Exists(key)
	}
}

//RemoveCache removes key-value from cache
func RemoveCache(ctx *gin.Context, key string) error {
	store, _ := ctx.Get(common.CACHE_STORE_KEY)
	if store == nil {
		panic("NO REDIS INSTANCE RUNNING...")
	} else {
		cacheStore := store.(*cache.Codec)
		return cacheStore.Delete(key)
	}
}

//RefreshCache resets session timeout value in cache
func RefreshCache(ctx *gin.Context, key string) {
	store, _ := ctx.Get(common.CACHE_STORE_KEY)
	if store == nil {
		panic("NO REDIS INSTANCE RUNNING...")
	} else {
		val := GetCache(ctx, key)
		if val != nil {
			SetCache(ctx, key, val,
				common.CACHE_SESSION_TIMEOUT)
		} else {
			fmt.Println("Cache Refresh Failed :=> ", key)
		}
	}
}

//GenerateSessionID generates unique sessionID for user
func GenerateSessionID(ctx *gin.Context) string {
	attempts := 2
	for {
		sessionID := utility.GenerateUUID()
		if len(sessionID) > 0 {
			if !CheckCache(ctx, sessionID) {
				return sessionID
			}
		} else {
			if attempts == 0 {
				return ""
			}
			attempts--
		}
	}
}
