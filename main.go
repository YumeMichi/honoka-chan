package main

import (
	"honoka-chan/config"
	"honoka-chan/router"
	_ "honoka-chan/tools"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// Router
	r := gin.Default()
	router.SifRouter(r)
	router.AsRouter(r)

	r.Run(":" + config.Conf.Settings.ServerPort)
}
