package handler

import (
	"honoka-chan/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PersonalNoticeHandler(ctx *gin.Context) {
	ctx.Header("Server-Version", "97.4.6")
	ctx.Header("user_id", "5802913")
	ctx.Header("authorize", "consumerKey=lovelive_test&timeStamp=1678640520&version=1.1&token=bS5G6TKTsw0aGxVQz8JWJTx8Tf73H0bF9Bq1PEw3UaxoEoUG8GcrrzaEbjOwEQJTrThgHpBlbwnMRl9ITGw1&nonce=5&requestTimeStamp=1678640520")
	ctx.Header("X-Message-Sign", "c/J6CJPZZeGk44VLWkgCfamhFFKlLheu4e2ga2KrEl6DkVyotQTDte39RGMJS+kd5qRg/nSh/QdRGEfMLIeXB2xj3TT/UB1a03wV6slr39B+Xd2ip8BiuDKxmqpw3ESdK25WdlY+fMXbXc4RkkrleqElME7jw+VZ2SiwIgiedNg=")
	ctx.String(http.StatusOK, resp.Gdpr)
}
