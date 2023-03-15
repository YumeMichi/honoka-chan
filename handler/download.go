package handler

import (
	"honoka-chan/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DownloadEventHandler(ctx *gin.Context) {
	ctx.Header("user_id", "5802913")
	ctx.Header("authorize", "consumerKey=lovelive_test&timeStamp=1678640521&version=1.1&token=bS5G6TKTsw0aGxVQz8JWJTx8Tf73H0bF9Bq1PEw3UaxoEoUG8GcrrzaEbjOwEQJTrThgHpBlbwnMRl9ITGw1&nonce=7&requestTimeStamp=1678640521")
	ctx.Header("X-Message-Sign", "anBqaPYp7nEprAbxhOouwx8e1dCn1PbmXNplSVtGT3QWNvS+bcMYnqr8QxluM5SIGi3esO8tzHsDzn+Z7tdohFtpkT1zdYQ5qK07eZ8VtyckaRnmx5QLzn7MMflFWr3P/HONS04QwLakoeZyNDCBmxjqOy6QzNsf6e4o9D+21AM=")
	ctx.String(http.StatusOK, resp.Event)
}
