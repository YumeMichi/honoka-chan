package handler

import (
	"encoding/json"
	"fmt"
	"honoka-chan/model"
	"honoka-chan/resp"
	"honoka-chan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiHandler(ctx *gin.Context) {
	// fmt.Println(c.PostForm("request_data"))
	var formdata []model.SifApi
	err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &formdata)
	if err != nil {
		fmt.Println(err)
		return
	}
	v1 := false
	for _, v := range formdata {
		// fmt.Println(v)
		if v.Module == "live" {
			v1 = true
			break
		}
	}
	ctx.Header("Server-Version", "97.4.6")
	ctx.Header("user_id", "5802913")
	if v1 {
		// 登录后第一次请求API
		ctx.Header("authorize", "consumerKey=lovelive_test&timeStamp=1678640521&version=1.1&token=bS5G6TKTsw0aGxVQz8JWJTx8Tf73H0bF9Bq1PEw3UaxoEoUG8GcrrzaEbjOwEQJTrThgHpBlbwnMRl9ITGw1&nonce=9&requestTimeStamp=1678640521")
		// ctx.Header("X-Message-Sign", "bNuSClKqt20FoGduZJI4pB1Y8xUwrrarvfsq0soqU5U97x7kGLESNoXSQVZcFfa1Eo4kgntEktokmDHzCbxpsYFvrD1mhbn++UOmcwXXCQRdxbbfhTt7MVfXbcqXAuFKAfkE37n9dkn1U0RnNt5U4m3mbRhLYT5B16ZcPGIPn/E=")
		ctx.String(http.StatusOK, utils.ReadAllText("data/test2.json"))
	} else {
		// 登录后第二次请求API
		ctx.Header("authorize", "consumerKey=lovelive_test&timeStamp=1678640523&version=1.1&token=bS5G6TKTsw0aGxVQz8JWJTx8Tf73H0bF9Bq1PEw3UaxoEoUG8GcrrzaEbjOwEQJTrThgHpBlbwnMRl9ITGw1&nonce=10&requestTimeStamp=1678640523")
		// ctx.Header("X-Message-Sign", "w6+DlFx4DcmaoNXoLu71PH9sOOQeGFX/x0aDayt0qpHa+Uw+MeA0EJossW3/OgZvZgcFxBK2tKrHZJxRrguwCp0lpaI6/onWrYhK9xZeAW33nrsuWWT52v4wyNPY236xGecrDrs9R0nmTOuxElQEKqdFeZYL/JZiuuxvUbwxMy8=")
		ctx.String(http.StatusOK, resp.Api2)
	}
}
