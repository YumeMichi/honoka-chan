package handler

import (
	"honoka-chan/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GdprHandler(ctx *gin.Context) {
	ctx.Header("Server-Version", "97.4.6")
	ctx.Header("user_id", "5802913")
	ctx.Header("authorize", "consumerKey=lovelive_test&timeStamp=1678640520&version=1.1&token=bS5G6TKTsw0aGxVQz8JWJTx8Tf73H0bF9Bq1PEw3UaxoEoUG8GcrrzaEbjOwEQJTrThgHpBlbwnMRl9ITGw1&nonce=4&requestTimeStamp=1678640520")
	ctx.Header("X-Message-Sign", "nGH0cQ34z9D3QnTGFDe2r2WMBGfTzYx+5oaJJbvYCqMESTDAQWlxq8X73OBLgdCkEIMuIAiRyF0z+0MQEVohDL4l6nDgcVDDJztCXP/W5ZXZh1wNgGhHZIDrwboNjsg1acq0+phBAiBEQQt6HipEdGRQh5fhAhhA717ns/C4iUI=")
	ctx.String(http.StatusOK, resp.Gdpr)
}
