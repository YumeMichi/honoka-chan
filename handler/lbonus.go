package handler

import (
	"honoka-chan/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LBonusExecuteHandler(ctx *gin.Context) {
	ctx.Header("Server-Version", "97.4.6")
	ctx.Header("user_id", "5802913")
	ctx.Header("authorize", "consumerKey=lovelive_test&timeStamp=1678640521&version=1.1&token=bS5G6TKTsw0aGxVQz8JWJTx8Tf73H0bF9Bq1PEw3UaxoEoUG8GcrrzaEbjOwEQJTrThgHpBlbwnMRl9ITGw1&nonce=8&requestTimeStamp=1678640521")
	ctx.Header("X-Message-Sign", "2UDMPncUOcUnQdLPCSuZIrRWK9yRgASUzCL3YiMpox/IUDNrHn7BRr/eo6iMQ1TlDTnNqzHj/ZP/8m4rpgvlDFHg16nlPUQJe0hqF9Ck3tjeHfT7wbwV75deoqIlPNONS5u3eCI8ZlhRf9VeDUBxRoVDRiq4xqiwjOILvmPAb1w=")
	ctx.String(http.StatusOK, resp.LBonus)
}
