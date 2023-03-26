package handler

import (
	"encoding/base64"
	"encoding/json"
	"errors"
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
		var res string
		var err error
		// fmt.Println(v)
		// fmt.Println(v.Module, v.Action)

		switch v.Module {
		case "login":
			if v.Action == "topInfo" {
				// fmt.Println("topInfo")
				res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "login_topinfo_result").Result()
			} else if v.Action == "topInfoOnce" {
				// fmt.Println("topInfoOnce")
				res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "login_topinfo_once_result").Result()
			}
		case "live":
			if v.Action == "liveStatus" {
				// fmt.Println("liveStatus")
				res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "live_status_result").Result()
			} else if v.Action == "schedule" {
				// fmt.Println("schedule")
				res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "live_list_result").Result()
			}
		case "unit":
			switch v.Action {
			case "unitAll":
				// fmt.Println("unitAll")
				res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "unit_list_result").Result()
			case "deckInfo":
				// fmt.Println("deckInfo")
				res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "unit_deck_result").Result()
			case "supporterAll":
				// fmt.Println("supporterAll")
				res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "unit_support_result").Result()
			case "removableSkillInfo":
				// fmt.Println("removableSkillInfo")
				res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "owning_equip_result").Result()
			case "accessoryAll":
				// fmt.Println("accessoryAll")
				res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "unit_accessory_result").Result()
			}
		case "costume":
			// fmt.Println("costumeList")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "costume_list_result").Result()
		case "album":
			// fmt.Println("albumAll")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "album_unit_result").Result()
		case "scenario":
			// fmt.Println("scenarioStatus")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "scenario_status_result").Result()
		case "subscenario":
			// fmt.Println("subscenarioStatus")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "subscenario_status_result").Result()
		case "eventscenario":
			// fmt.Println("status")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "event_scenario_result").Result()
		case "multiunit":
			// fmt.Println("multiunitscenarioStatus")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "multi_unit_scenario_result").Result()
		case "payment":
			// fmt.Println("productList")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "product_result").Result()
		case "banner":
			// fmt.Println("bannerList")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "banner_result").Result()
		case "notice":
			// fmt.Println("noticeMarquee")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "item_marquee_result").Result()
		case "user":
			// fmt.Println("getNavi")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "user_intro_result").Result()
		case "navigation":
			// fmt.Println("specialCutin")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "special_cutin_result").Result()
		case "award":
			// fmt.Println("awardInfo")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "award_result").Result()
		case "background":
			// fmt.Println("backgroundInfo")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "background_result").Result()
		case "stamp":
			// fmt.Println("stampInfo")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "stamp_result").Result()
		case "exchange":
			// fmt.Println("owningPoint")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "exchange_point_result").Result()
		case "livese":
			// fmt.Println("liveseInfo")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "live_se_result").Result()
		case "liveicon":
			// fmt.Println("liveiconInfo")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "live_icon_result").Result()
		case "item":
			// fmt.Println("list")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "item_list_result").Result()
		case "marathon":
			// fmt.Println("marathonInfo")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "marathon_result").Result()
		case "challenge":
			// fmt.Println("challengeInfo")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "challenge_result").Result()
		case "museum":
			// fmt.Println("info")
			res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "museum_result").Result()
		case "profile":
			if v.Action == "liveCnt" {
				// fmt.Println("liveCnt")
				res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "profile_livecnt_result").Result()
			} else if v.Action == "cardRanking" {
				// fmt.Println("cardRanking")
				res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "profile_card_ranking_result").Result()
			} else if v.Action == "profileInfo" {
				// fmt.Println("profileInfo")
				res, err = database.RedisCli.HGet(database.RedisCtx, "temp_dataset", "profile_info_result").Result()
			}
		default:
			// fmt.Println("Fuck you!")
			err = errors.New("invalid option")
		}

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
