package handler

import (
	"honoka-chan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GdprHandler(ctx *gin.Context) {
	ctx.Header("Server-Version", "97.4.6")
	ctx.Header("user_id", "3241988")
	ctx.Header("authorize", "consumerKey=lovelive_test&timeStamp=1679236700&version=1.1&token=cHPoOHP5dAs2dh30EkOW8FndO07xlpKHrDRdVOtT7Whlo1opiEMXSwk1JJdAFd4cSeKQvGVRwH2Z7sFh1gnz3gd&nonce=4&requestTimeStamp=1679236697")
	ctx.Header("X-Message-Sign", "El2L/coKTfbHF8shto7SjxnukwicN1BtIMhvD21uKy+9ISwH0G2fI4aCemKn77o54H2zv+mVN4osDK0N0Zi86lGyx0rFXMlQ75D7XKX5KbWAhHcgn+W8t6tk2R0PVUyEeo1gtHgjNauT1asNK+PDJ7h3WWINPFgfVSnldCUnYLk=")
	ctx.String(http.StatusOK, utils.ReadAllText("assets/gdpr.json"))
}
