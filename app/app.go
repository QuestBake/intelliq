package main

import (
	"github.com/gin-gonic/gin"

	"intelliq/app/approuter"
	"intelliq/app/common"
	"intelliq/app/config"
	"intelliq/app/logger"
	"intelliq/app/security"
)

var router *gin.Engine

func main() {
	logger.InitLogger(common.LOG_FILE, common.LOG_MAX_BYTES, common.LOG_BACKUP_COUNT)
	defer logger.Logger.Close()
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
		logger.Logger.Error("Router Failed")
	}
}
