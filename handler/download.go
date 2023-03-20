package handler

import (
	"honoka-chan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DownloadEventHandler(ctx *gin.Context) {
	ctx.Header("user_id", "3241988")
	ctx.Header("authorize", "consumerKey=lovelive_test&timeStamp=1679236701&version=1.1&token=cHPoOHP5dAs2dh30EkOW8FndO07xlpKHrDRdVOtT7Whlo1opiEMXSwk1JJdAFd4cSeKQvGVRwH2Z7sFh1gnz3gd&nonce=8&requestTimeStamp=1679236698")
	ctx.Header("X-Message-Sign", "rFJdTYrGPvG9V/IESgiAZu4qod4VgVTe5/oj8/5fA2mzXhJkKAD7nE8IOI9MQmjwECFYRc4KLPjtqsxuosloIjxTk6pHiAPtGvzkajqBivbbaBIB2OdIZQarqMRURN5M0TGLOC0vzO9xP5fYKJzMHSskJsmMdunPWLkUz3eqoZU=")
	ctx.String(http.StatusOK, utils.ReadAllText("assets/event.json"))
}
