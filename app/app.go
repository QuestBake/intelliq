package main

import (
	"github.com/gin-gonic/gin"

	"intelliq/app/approuter"
	"intelliq/app/config"
	"intelliq/app/logger"
	"intelliq/app/security"
)

var router *gin.Engine

func main() {
	logger.InitLogger(config.Conf.Get("log.log_file").(string),
		config.Conf.Get("log.log_max_bytes").(int64), config.Conf.Get("log.log_backup_count").(int))
	var configErr error
	config.Conf, configErr = config.LoadConfig("app_config.toml")
	if configErr != nil {
		logger.Logger.Error("Error while reading config file!!")
	}
	defer logger.Logger.Close()
	router = gin.Default()
	if router != nil {
		config.DBConnect()
		config.CacheConnect(router)
		security.EnableSecurity(router)
		approuter.AddRouters(router)
		//router.Run(common.APP_PORT)
		router.RunTLS(config.Conf.Get("app.app_port").(string), config.Conf.Get("security.ssl_certificate_path").(string),
			config.Conf.Get("security.ssl_key_filepath").(string))
	} else {
		logger.Logger.Error("Router Failed")
	}
}
