package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ScenarioStartupHandler(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	ctx.Header("X-Powered-By", "KLab Native APP Platform")
	ctx.Header("server_version", "20120129")
	ctx.Header("Server-Version", "97.4.6")
	ctx.Header("user_id", "5802913")
	ctx.Header("version_up", "0")
	ctx.Header("status_code", "200")
	ctx.Header("authorize", "consumerKey=lovelive_test&timeStamp=1678786443&version=1.1&token=2gjbstqCHAlIv600Pk0AJN5uvfPYgLvfSKB418sMIoclOy1M9UP7LUC1CG14tPicVzOq2MBOXLvpmNRwyWIi9Qf&nonce=17&requestTimeStamp=1678786443")
	ctx.Header("X-Message-Sign", "L+cK2fiaDxkRwJCmT6lpMu9RM/emmUye32RtS8C36I7OpBhGiZLbiPN30dPXITWDIUjUyPsVCqvLVzEy484Q+NHvQfDEUuS5SnKOtoNxW9Zk6kjyckjJKFIbsdH61cbb4pfzmSptcaByZ6ieNIpbHTzsvpu54JFOXb84g859Pxs=")

	res := `{"response_data":{"scenario_id":1,"scenario_adjustment":50,"server_timestamp":1678785161},"release_info":[],"status_code":200}`
	ctx.String(http.StatusOK, res)
}
