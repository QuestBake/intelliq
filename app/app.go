package main

import (
	"intelliq/app/approuter"
	"intelliq/app/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	config.DBConnect()
	config.CacheConnect(router)
	approuter.AddRouters(router)
	enableCors()
	router.Run()
}

func enableCors() {
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:           []string{"Origin", "currentrole", "Content-Type", "X-Requested-With", "Accept"},
		ExposeHeaders:          []string{"currentrole"},
		AllowCredentials:       false,
		MaxAge:                 12 * time.Hour,
		AllowBrowserExtensions: true,
	}))
}
