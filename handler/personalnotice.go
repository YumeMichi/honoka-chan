package handler

import (
	"honoka-chan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PersonalNoticeHandler(ctx *gin.Context) {
	ctx.Header("Server-Version", "97.4.6")
	ctx.Header("user_id", "3241988")
	ctx.Header("authorize", "consumerKey=lovelive_test&timeStamp=1679236701&version=1.1&token=cHPoOHP5dAs2dh30EkOW8FndO07xlpKHrDRdVOtT7Whlo1opiEMXSwk1JJdAFd4cSeKQvGVRwH2Z7sFh1gnz3gd&nonce=5&requestTimeStamp=1679236697")
	ctx.Header("X-Message-Sign", "mTvY9EUM4LAomFxHQPINslVF8KBJ/nWZvCmVzYYyFln+M23/T05cuT/E6FUt9ExwmGMFy6TqtbZwcoomFWaEm38uJH2nQy/3RDjS0L26AsyFOHDIUOK11a4qHxv309sRjb04KhckTmzJERTooCnRturTYcNYet0g01vz2Geu4Ew=")
	ctx.String(http.StatusOK, utils.ReadAllText("assets/personalnotice.json"))
}
