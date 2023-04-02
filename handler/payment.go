package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/config"
	"honoka-chan/database"
	"honoka-chan/encrypt"
	"honoka-chan/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductResp struct {
	ResponseData ProductData   `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}
type RestrictionInfo struct {
	Restricted bool `json:"restricted"`
}
type UnderAgeInfo struct {
	BirthSet    bool        `json:"birth_set"`
	HasLimit    bool        `json:"has_limit"`
	LimitAmount interface{} `json:"limit_amount"`
	MonthUsed   int         `json:"month_used"`
}
type ProductData struct {
	RestrictionInfo  RestrictionInfo `json:"restriction_info"`
	UnderAgeInfo     UnderAgeInfo    `json:"under_age_info"`
	SnsProductList   []interface{}   `json:"sns_product_list"`
	ProductList      []interface{}   `json:"product_list"`
	SubscriptionList []interface{}   `json:"subscription_list"`
	ShowPointShop    bool            `json:"show_point_shop"`
	ServerTimestamp  int64           `json:"server_timestamp"`
}

func ProductListHandler(ctx *gin.Context) {
	reqTime := time.Now().Unix()

	authorizeStr := ctx.Request.Header["Authorize"]
	authToken, err := utils.GetAuthorizeToken(authorizeStr)
	if err != nil {
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}
	userId := ctx.Request.Header[http.CanonicalHeaderKey("User-ID")]
	if len(userId) == 0 {
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}

	if !database.MatchTokenUid(authToken, userId[0]) {
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}

	nonce, err := utils.GetAuthorizeNonce(authorizeStr)
	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}
	nonce++

	respTime := time.Now().Unix()
	newAuthorizeStr := fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", respTime, authToken, nonce, userId[0], reqTime)
	// fmt.Println(newAuthorizeStr)

	prodReesp := ProductResp{
		ResponseData: ProductData{
			RestrictionInfo: RestrictionInfo{
				Restricted: false,
			},
			UnderAgeInfo: UnderAgeInfo{
				BirthSet:    true,
				HasLimit:    false,
				LimitAmount: nil,
				MonthUsed:   0,
			},
			SnsProductList:   []interface{}{},
			ProductList:      []interface{}{},
			SubscriptionList: []interface{}{},
			ShowPointShop:    true,
			ServerTimestamp:  time.Now().Unix(),
		},
		ReleaseInfo: []interface{}{},
		StatusCode:  200,
	}
	resp, err := json.Marshal(prodReesp)
	if err != nil {
		panic(err)
	}
	xms := encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")
	xms64 := base64.RawStdEncoding.EncodeToString(xms)

	ctx.Header("Server-Version", config.Conf.Server.VersionNumber)
	ctx.Header("user_id", userId[0])
	ctx.Header("authorize", newAuthorizeStr)
	ctx.Header("X-Message-Sign", xms64)
	ctx.String(http.StatusOK, string(resp))
}
