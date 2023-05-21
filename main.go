package main

import (
	"honoka-chan/router"
	_ "honoka-chan/tools"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin
	gin.SetMode(gin.ReleaseMode)

	// Router
	r := gin.Default()

	// SIF
	router.SifRouter(r)

	// AS
	router.AsRouter(r)

	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
