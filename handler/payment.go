package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/encrypt"
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

func ProductList(ctx *gin.Context) {
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
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}
