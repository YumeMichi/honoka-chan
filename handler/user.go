package handler

import (
	"honoka-chan/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetNotificationTokenHandler(ctx *gin.Context) {
	ctx.Header("Server-Version", "97.4.6")
	ctx.Header("user_id", "3241988")
	ctx.Header("authorize", "consumerKey=lovelive_test&timeStamp=1678891478&version=1.1&token=cTMEyAfIDErcKUyCgBM6ZuJv8UBRgSzG4z4MfqYhR7xSHBYhIA9ofaVKtSefeiP2LTKIfbnCfE5dppYw8Af&nonce=18&requestTimeStamp=1678891478")
	ctx.Header("X-Message-Sign", "bNZj/YjRADDccb5vcF/NHI9Kin3bkM3ECBQdxttXvCBqzoSBX6yWuEn9Fsjx+Yp3g2D9CcONZzqJlAvXbatGHJkbClSXuomLXVOcNmQYgidjyUvC7CceoSvCbL8U4Ge12tyGd8V2EMHVZfxqPKdHsJSGaOFbUpmo7wAhKVfuEjg=")
	ctx.String(http.StatusOK, resp.NotificationToken)
}
