package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/config"
	"honoka-chan/database"
	"honoka-chan/encrypt"
	"honoka-chan/model"
	"honoka-chan/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ApiHandler(ctx *gin.Context) {
	reqTime := time.Now().Unix()
	// fmt.Println(ctx.PostForm("request_data"))
	var formdata []model.SifApi
	err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &formdata)
	if err != nil {
		fmt.Println(err)
		return
	}
	var results []interface{}
	for _, v := range formdata {
		var res []byte
		var key string
		var err error
		// fmt.Println(v)
		// fmt.Println(v.Module, v.Action)

		switch v.Module {
		case "login":
			if v.Action == "topInfo" {
				// fmt.Println("topInfo")
				key = "login_topinfo_result"
			} else if v.Action == "topInfoOnce" {
				// fmt.Println("topInfoOnce")
				key = "login_topinfo_once_result"
			}
		case "live":
			if v.Action == "liveStatus" {
				// fmt.Println("liveStatus")
				key = "live_status_result"
			} else if v.Action == "schedule" {
				// fmt.Println("schedule")
				key = "live_list_result"
			}
		case "unit":
			switch v.Action {
			case "unitAll":
				// fmt.Println("unitAll")
				key = "unit_list_result"
			case "deckInfo":
				// fmt.Println("deckInfo")
				key = "unit_deck_result"
			case "supporterAll":
				// fmt.Println("supporterAll")
				key = "unit_support_result"
			case "removableSkillInfo":
				// fmt.Println("removableSkillInfo")
				key = "owning_equip_result"
			case "accessoryAll":
				// fmt.Println("accessoryAll")
				key = "unit_accessory_result"
			}
		case "costume":
			// fmt.Println("costumeList")
			key = "costume_list_result"
		case "album":
			// fmt.Println("albumAll")
			key = "album_unit_result"
		case "scenario":
			// fmt.Println("scenarioStatus")
			key = "scenario_status_result"
		case "subscenario":
			// fmt.Println("subscenarioStatus")
			key = "subscenario_status_result"
		case "eventscenario":
			// fmt.Println("status")
			key = "event_scenario_result"
		case "multiunit":
			// fmt.Println("multiunitscenarioStatus")
			key = "multi_unit_scenario_result"
		case "payment":
			// fmt.Println("productList")
			key = "product_result"
		case "banner":
			// fmt.Println("bannerList")
			key = "banner_result"
		case "notice":
			// fmt.Println("noticeMarquee")
			key = "item_marquee_result"
		case "user":
			// fmt.Println("getNavi")
			key = "user_intro_result"
		case "navigation":
			// fmt.Println("specialCutin")
			key = "special_cutin_result"
		case "award":
			// fmt.Println("awardInfo")
			key = "award_result"
		case "background":
			// fmt.Println("backgroundInfo")
			key = "background_result"
		case "stamp":
			// fmt.Println("stampInfo")
			key = "stamp_result"
		case "exchange":
			// fmt.Println("owningPoint")
			key = "exchange_point_result"
		case "livese":
			// fmt.Println("liveseInfo")
			key = "live_se_result"
		case "liveicon":
			// fmt.Println("liveiconInfo")
			key = "live_icon_result"
		case "item":
			// fmt.Println("list")
			key = "item_list_result"
		case "marathon":
			// fmt.Println("marathonInfo")
			key = "marathon_result"
		case "challenge":
			// fmt.Println("challengeInfo")
			key = "challenge_result"
		case "museum":
			// fmt.Println("info")
			key = "museum_result"
		case "profile":
			if v.Action == "liveCnt" {
				// fmt.Println("liveCnt")
				key = "profile_livecnt_result"
			} else if v.Action == "cardRanking" {
				// fmt.Println("cardRanking")
				key = "profile_card_ranking_result"
			} else if v.Action == "profileInfo" {
				// fmt.Println("profileInfo")
				key = "profile_info_result"
			}
		default:
			// fmt.Println("Fuck you!")
			// err = errors.New("invalid option")
		}

		res, err = database.LevelDb.Get([]byte(key))
		if err != nil {
			panic(err)
		}

		var result interface{}
		err = json.Unmarshal([]byte(res), &result)
		if err != nil {
			panic(err)
		}
		results = append(results, result)
	}
	// fmt.Println(results)
	b, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	rp := model.Response{
		ResponseData: b,
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}
	b, err = json.Marshal(rp)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(b))

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

	xms := encrypt.RSA_Sign_SHA1(b, "privatekey.pem")
	xms64 := base64.RawStdEncoding.EncodeToString(xms)

	ctx.Header("Server-Version", config.Conf.Server.VersionNumber)
	ctx.Header("user_id", userId[0])
	ctx.Header("authorize", newAuthorizeStr)
	ctx.Header("X-Message-Sign", xms64)
	ctx.String(http.StatusOK, string(b))
}
