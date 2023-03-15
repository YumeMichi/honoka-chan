package handler

import (
	"honoka-chan/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserInfoHandler(ctx *gin.Context) {
	ctx.Header("Server-Version", "97.4.6")
	ctx.Header("user_id", "5802913")
	ctx.Header("authorize", "consumerKey=lovelive_test&timeStamp=1678640520&version=1.1&token=bS5G6TKTsw0aGxVQz8JWJTx8Tf73H0bF9Bq1PEw3UaxoEoUG8GcrrzaEbjOwEQJTrThgHpBlbwnMRl9ITGw1&nonce=3&requestTimeStamp=1678640520")
	ctx.Header("X-Message-Sign", "oa9sisoLJILWOKto4CAfb1jJwzoozpkpIAhnXl3s6m5jy75Zyb8tM+Jds6rqytnJt/labC/hFH0dhna/qqGbnTeLq6zUwHrTf8crz4uwtVi8qOBwvSbtA1dgbiKJr9raAflgJOQj6cft7XGP4bdkYfbuNxPS9gFhK+MG7b8S3Z8=")
	ctx.String(http.StatusOK, resp.UserInfo)
}
