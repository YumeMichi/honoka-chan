package handler

import (
	"honoka-chan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserInfoHandler(ctx *gin.Context) {
	ctx.Header("Server-Version", "97.4.6")
	ctx.Header("user_id", "3241988")
	ctx.Header("authorize", "consumerKey=lovelive_test&timeStamp=1679236700&version=1.1&token=cHPoOHP5dAs2dh30EkOW8FndO07xlpKHrDRdVOtT7Whlo1opiEMXSwk1JJdAFd4cSeKQvGVRwH2Z7sFh1gnz3gd&nonce=3&requestTimeStamp=1679236697")
	ctx.Header("X-Message-Sign", "mGcmcvrm22pt5zJw1CJxe/6Y9R4aZPb6znja4jxioWHlinjWU5nEq61qyslW0bX6uVWTifz17eSDdidlJHccusbQaKXOyoFfbQC1hpArv97b3RJGrnDK7iShPOTz3+mwYiUhtXrJ3oohRGH1siEG0G3H4pSK3JHnAbPlF84cR4w=")
	ctx.String(http.StatusOK, utils.ReadAllText("assets/userinfo.json"))
}
