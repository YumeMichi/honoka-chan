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
	ctx.Header("user_id", "5802913")
	ctx.Header("authorize", "consumerKey=lovelive_test&timeStamp=1678640534&version=1.1&token=bS5G6TKTsw0aGxVQz8JWJTx8Tf73H0bF9Bq1PEw3UaxoEoUG8GcrrzaEbjOwEQJTrThgHpBlbwnMRl9ITGw1&nonce=11&requestTimeStamp=1678640534")
	ctx.Header("X-Message-Sign", "vP/1uPKsGnArUaIL7lPDgZC4lGPjwVZDvFWiGXFfxrRzq6wQ8Veq3623KCl05/2hTTQnOm3L6bNu377s//R/7SieeQ7L7FNWLptTHoyyqgjZe16CJzyqFGbmm9pxxZDttJ74zkuBO1astYBPxh2+qws4LduZUS0c0VklzP5nNoY=")
	ctx.String(http.StatusOK, resp.AnnounceState)
}
