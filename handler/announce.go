package handler

import (
	"honoka-chan/config"
	"honoka-chan/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AnnounceIndexHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello, world!")
}

func AnnounceCheckStateHandler(ctx *gin.Context) {
	ctx.Header("Server-Version", config.Conf.Server.VersionNumber)
	ctx.Header("user_id", "3241988")
	ctx.Header("authorize", "consumerKey=lovelive_test&timeStamp=1679360075&version=1.1&token=4JM3w3i6FDXr4lHN4K2Q5ys7Fn56QMT5ytIzPLCZ0ItuouoraRlfCYPlnzqsFsE1v9Phd66Cw4bmjcKoxE52gd&nonce=12&requestTimeStamp=1679360075")
	ctx.Header("X-Message-Sign", "CGfQ1rVjIRXxdq1zPyfp8aHDyekfCJ6PwzZ+a6JZkXJdyzGZXnFK6ZZ9lup5KLYLXKzSOmbzeOUbTV6lYdE5G+hoP+oVSpnKDGt1NKdoy5IeJihfXVbNjM5AC8NF7d8oMGk/zTVsVqHHENCX4bbMMIQuFP+K+iw2SiRRD1dbZPw=")
	ctx.String(http.StatusOK, resp.AnnounceState)
}
