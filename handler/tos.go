package handler

import (
	"honoka-chan/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TosCheckHandler(ctx *gin.Context) {
	ctx.Header("user_id", "5802913")
	ctx.Header("authorize", "consumerKey=lovelive_test&timeStamp=1678640520&version=1.1&token=bS5G6TKTsw0aGxVQz8JWJTx8Tf73H0bF9Bq1PEw3UaxoEoUG8GcrrzaEbjOwEQJTrThgHpBlbwnMRl9ITGw1&nonce=6&requestTimeStamp=1678640520")
	ctx.Header("X-Message-Sign", "lHGAUrSt27AAUkeB+WJCSLjs9xZdyNnDRlZVZ0VRXFhqDnVBag0lwSQEUQoedh3/FyIHazbfjuTJw9REsgSJwX05I7GW7KoqPIYNoRLiLgjM6Y7MOZWrpVdipWl0q7IPS3mKyL8ye6sRZ8TfWx8cFfnohp3U7FnvrAEWdh1h71I=")
	ctx.String(http.StatusOK, resp.TosCheck)
}
