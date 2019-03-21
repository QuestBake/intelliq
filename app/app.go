package main

import (
	"github.com/gin-gonic/gin"

	"intelliq/app/approuter"
	"intelliq/app/common"
	"intelliq/app/config"
	"intelliq/app/security"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	if router != nil {
		config.DBConnect()
		config.CacheConnect(router)
		security.EnableSecurity(router)
		approuter.AddRouters(router)
		//router.Run(common.APP_PORT)
		router.RunTLS(common.APP_PORT, common.SSL_CERT_FILEPATH,
			common.SSL_KEY_FILEPATH)
	} else {
		panic("Router Failed")
	}
}
