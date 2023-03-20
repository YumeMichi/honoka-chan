package handler

import (
	"honoka-chan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LBonusExecuteHandler(ctx *gin.Context) {
	ctx.Header("Server-Version", "97.4.6")
	ctx.Header("user_id", "3241988")
	ctx.Header("authorize", "consumerKey=lovelive_test&timeStamp=1679236701&version=1.1&token=cHPoOHP5dAs2dh30EkOW8FndO07xlpKHrDRdVOtT7Whlo1opiEMXSwk1JJdAFd4cSeKQvGVRwH2Z7sFh1gnz3gd&nonce=9&requestTimeStamp=1679236698")
	ctx.Header("X-Message-Sign", "pV5H3WfRtj2zpDYuYwt9BuB8jMUJiXrGXQbsemJA+8sX7/c9s4mnbMFTKDD3cxK1mSeCLNhJVtR1M6QKZVgbCjyQGSVPR1EG1cTumR9T5LFF6ighJWV7EEYxbeYgJjAEcjVHOgB3d2hy7SK7u4oCatEgXhbJMQYGV5lH2gdwEpw=")
	ctx.String(http.StatusOK, utils.ReadAllText("assets/lbonus.json"))
}
