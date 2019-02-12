package main

import (
	"intelliq/app/approuter"
	"intelliq/app/config"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	approuter.AddRouters(router)
	config.Connect()
	router.Run()
}
