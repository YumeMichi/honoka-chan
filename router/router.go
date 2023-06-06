package router

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/config"
	"honoka-chan/database"
	"honoka-chan/encrypt"
	"honoka-chan/handler"
	"honoka-chan/middleware"
	"honoka-chan/model"
	"honoka-chan/utils"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"xorm.io/xorm"
)

var (
	sessionKey = "12345678123456781234567812345678"

	mEng *xorm.Engine
)

func init() {
	mEng = config.MainEng
}

func SifRouter(r *gin.Engine) {
	// Static
	r.Static("/static", "static")

	var files []string
	_ = filepath.Walk("static/templates", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})
	r.LoadHTMLFiles(files...)

	// session
	store := cookie.NewStore([]byte("llsif"))
	r.Use(sessions.Sessions("llsif", store))

	// /
	r.Any("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello, world!")
	})

	// Private APIs
	v1 := r.Group("v1")
	{
		v1.GET("/basic/getcode", handler.GetCode)
		v1.POST("/account/active", handler.Active)
		v1.POST("/account/initialize", handler.Initialize)
		v1.POST("/account/loginauto", handler.LoginAuto)
		v1.POST("/account/login", handler.AccountLogin)
		v1.POST("/account/reportRole", handler.ReportRole)
		v1.POST("/basic/getcode", handler.GetCode)
		v1.POST("/basic/getProductList", handler.GetProductList)
		v1.POST("/basic/handshake", handler.Handshake)
		v1.POST("/basic/loginarea", handler.LoginArea)
		v1.POST("/basic/publickey", handler.PublicKey)
		v1.POST("/guest/status", handler.GuestStatus)
	}
	r.GET("/agreement/all", handler.Agreement)
	r.GET("/integration/appReport/initialize", handler.ReportApp)
	r.POST("/report/ge/app", handler.ReportLog)
	// Private APIs

	// Server APIs
	m := r.Group("main.php").Use(middleware.Common)
	{
		m.POST("/album/seriesAll", middleware.ParseMultipartForm, handler.AlbumSeriesAll)
		m.POST("/announce/checkState", middleware.ParseMultipartForm, handler.AnnounceCheckState)
		m.POST("/api", middleware.ParseMultipartForm, handler.Api)
		m.POST("/award/set", handler.AwardSet)
		m.POST("/background/set", handler.BackgroundSet)
		m.POST("/download/additional", handler.DownloadAdditional)
		m.POST("/download/batch", handler.DownloadBatch)
		m.POST("/download/event", handler.DownloadEvent)
		m.POST("/download/getUrl", handler.DownloadUrl)
		m.POST("/download/update", handler.DownloadUpdate)
		m.POST("/event/eventList", middleware.ParseMultipartForm, handler.EventList)
		m.POST("/gdpr/get", middleware.ParseMultipartForm, handler.Gdpr)
		m.POST("/lbonus/execute", handler.LBonusExecute)
		m.POST("/live/gameover", handler.GameOver)
		m.POST("/live/partyList", handler.PartyList)
		m.POST("/live/play", middleware.ParseMultipartForm, handler.PlayLive)
		m.POST("/live/preciseScore", middleware.ParseMultipartForm, handler.PlayScore)
		m.POST("/live/reward", middleware.ParseMultipartForm, handler.PlayReward)
		m.POST("/login/authkey", middleware.AuthKey, handler.AuthKey)
		m.POST("/login/login", middleware.Login, handler.Login)
		m.POST("/multiunit/scenarioStartup", handler.MultiUnitStartUp)
		m.POST("/museum/info", middleware.ParseMultipartForm, handler.MuseumInfo)
		m.POST("/notice/noticeFriendGreeting", middleware.ParseMultipartForm, handler.NoticeFriendGreeting)
		m.POST("/notice/noticeFriendVariety", middleware.ParseMultipartForm, handler.NoticeFriendVariety)
		m.POST("/notice/noticeUserGreetingHistory", handler.NoticeUserGreeting)
		m.POST("/payment/productList", middleware.ParseMultipartForm, handler.ProductList)
		m.POST("/personalnotice/get", middleware.ParseMultipartForm, handler.PersonalNotice)
		m.POST("/profile/profileRegister", handler.ProfileRegister)
		m.POST("/scenario/reward", handler.ScenarioReward)
		m.POST("/scenario/startup", handler.ScenarioStartup)
		m.POST("/subscenario/reward", handler.SubScenarioStartup)
		m.POST("/subscenario/startup", handler.SubScenarioStartup)
		m.POST("/tos/tosCheck", middleware.ParseMultipartForm, handler.TosCheck)
		m.POST("/unit/deck", handler.SetDeck)
		m.POST("/unit/deckName", handler.SetDeckName)
		m.POST("/unit/favorite", handler.SetDisplayRank)
		m.POST("/unit/removableSkillEquipment", handler.RemoveSkillEquip)
		m.POST("/unit/setDisplayRank", handler.SetDisplayRank)
		m.POST("/unit/wearAccessory", handler.WearAccessory)
		m.POST("/user/changeName", handler.ChangeName)
		m.POST("/user/changeNavi", handler.ChangeNavi)
		m.POST("/user/setNotificationToken", handler.SetNotificationToken)
		m.POST("/user/userInfo", middleware.ParseMultipartForm, handler.UserInfo)
	}
	r.GET("/webview.php/announce/index", handler.AnnounceIndex)
	// Server APIs

	// Manga
	r.GET("/manga", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "common/manga.html", gin.H{})
	})

	// WebUI
	w := r.Group("admin").Use(middleware.WebAuth)
	{
		w.GET("/index", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "admin/index.html", gin.H{
				"url": strings.Split(ctx.Request.URL.String(), "?")[0],
			})
		})
		w.GET("/login", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "admin/login.html", gin.H{})
		})
		w.POST("/login", handler.WebLogin)
		w.GET("/logout", handler.WebLogout)
		w.GET("/card", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "admin/card.html", gin.H{
				"menu": 1,
				"url":  strings.Split(ctx.Request.URL.String(), "?")[0],
			})
		})
		w.GET("/upload", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "admin/upload.html", gin.H{
				"menu": 1,
				"url":  strings.Split(ctx.Request.URL.String(), "?")[0],
			})
		})
		w.POST("/upload", handler.Upload)
	}
}

func AsRouter(r *gin.Engine) {
	r.Static("/llas-dev/static", "static/llas-dev/static")
	s := r.Group("ep3071").Use(middleware.CommonAs)
	{
		s.POST("/login/login", func(ctx *gin.Context) {
			reqBody := ctx.GetString("reqBody")

			var mask string
			req := gjson.Parse(reqBody)
			req.ForEach(func(key, value gjson.Result) bool {
				if value.Get("mask").String() != "" {
					mask = value.Get("mask").String()
					return false
				}
				return true
			})
			// fmt.Println("Request data:", req.String())
			// fmt.Println("Mask:", mask)

			mask64, err := base64.StdEncoding.DecodeString(mask)
			if err != nil {
				panic(err)
			}
			randomBytes := encrypt.RSA_DecryptOAEP(mask64, "privatekey.pem")
			// fmt.Println("Random Bytes:", randomBytes)

			newKey := utils.SliceXor(randomBytes, []byte(sessionKey))
			newKey64 := base64.StdEncoding.EncodeToString(newKey)

			loginBody := strings.ReplaceAll(utils.ReadAllText("assets/as/login.json"), "SESSION_KEY", newKey64)
			resp := SignResp(ctx.GetString("ep"), loginBody, config.StartUpKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/bootstrap/fetchBootstrap", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/fetchBootstrap.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/billing/fetchBillingHistory", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/fetchBillingHistory.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/notice/fetchNotice", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/fetchNotice.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/asset/getPackUrl", func(ctx *gin.Context) {
			reqBody := ctx.GetString("reqBody")

			var req []model.AsReq
			err := json.Unmarshal([]byte(reqBody), &req)
			if err != nil {
				panic(err)
			}
			// fmt.Println(req)

			packBody, ok := req[0].(map[string]interface{})
			if !ok {
				panic("Assertion failed!")
			}
			// fmt.Println(packBody)

			packNames, ok := packBody["pack_names"].([]interface{})
			if !ok {
				panic("Assertion failed!")
			}

			// 生成更新包 map
			var packageList []string
			var urlList []string

			err = json.Unmarshal([]byte(utils.ReadAllText("assets/as/packages.json")), &packageList)
			if err != nil {
				panic(err)
			}

			err = json.Unmarshal([]byte(utils.ReadAllText("assets/as/urls.json")), &urlList)
			if err != nil {
				panic(err)
			}

			if len(packageList) != len(urlList) {
				fmt.Println("File size not match!")
				return
			}

			packageUrls := map[string]string{}
			for k, p := range packageList {
				packageUrls[p] = urlList[k]
			}

			// Response
			var respUrls []string
			for _, pack := range packNames {
				packName, ok := pack.(string)
				if !ok {
					panic("Assertion failed!")
				}
				// fmt.Println(packageUrls[packName])
				respUrls = append(respUrls, packageUrls[packName])
			}

			urlResp := model.PackUrlRespBody{
				UrlList: respUrls,
			}

			var resp []model.AsResp
			resp = append(resp, time.Now().UnixMilli()) // 时间戳
			resp = append(resp, config.MasterVersion)   // 版本号
			resp = append(resp, 0)                      // 固定值
			resp = append(resp, urlResp)                // 数据体

			mm, err := json.Marshal(resp)
			if err != nil {
				panic(err)
			}
			// fmt.Println(string(mm))

			signBody := mm[1 : len(mm)-1]
			// fmt.Println(string(signBody))

			ep := strings.ReplaceAll(ctx.Request.URL.String(), "/ep3071", "")
			// fmt.Println(ep)

			sign := encrypt.HMAC_SHA1_Encrypt([]byte(ep+" "+string(signBody)), []byte(sessionKey))

			resp = append(resp, sign)
			mm, err = json.Marshal(resp)
			if err != nil {
				panic(err)
			}
			// fmt.Println(string(mm))

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, string(mm))
		})
		s.POST("/card/updateCardNewFlag", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/updateCardNewFlag.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/bootstrap/getClearedPlatformAchievement", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/getClearedPlatformAchievement.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/live/fetchLiveMusicSelect", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/fetchLiveMusicSelect.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/liveMv/start", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/liveMvStart.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/liveMv/saveDeck", func(ctx *gin.Context) {
			reqBody := ctx.GetString("reqBody")

			var req []model.AsReq
			err := json.Unmarshal([]byte(reqBody), &req)
			if err != nil {
				panic(err)
			}

			body, ok := req[0].(map[string]interface{})
			if !ok {
				panic("Assertion failed!")
			}

			reqB, err := json.Marshal(body)
			if err != nil {
				panic(err)
			}
			// fmt.Println(string(reqB))

			saveReq := model.LiveSaveDeckReq{}
			err = json.Unmarshal(reqB, &saveReq)
			if err != nil {
				panic(err)
			}
			// fmt.Println(saveReq)

			userLiveMvDeckInfo := model.UserLiveMvDeckInfo{
				LiveMasterID: saveReq.LiveMasterID,
			}

			memberIds := map[int]int{}
			for k, v := range saveReq.MemberMasterIDByPos {
				if k%2 == 0 {
					memberId := saveReq.MemberMasterIDByPos[k+1]
					memberIds[v] = memberId

					switch v {
					case 1:
						userLiveMvDeckInfo.MemberMasterID1 = memberId
					case 2:
						userLiveMvDeckInfo.MemberMasterID2 = memberId
					case 3:
						userLiveMvDeckInfo.MemberMasterID3 = memberId
					case 4:
						userLiveMvDeckInfo.MemberMasterID4 = memberId
					case 5:
						userLiveMvDeckInfo.MemberMasterID5 = memberId
					case 6:
						userLiveMvDeckInfo.MemberMasterID6 = memberId
					case 7:
						userLiveMvDeckInfo.MemberMasterID7 = memberId
					case 8:
						userLiveMvDeckInfo.MemberMasterID8 = memberId
					case 9:
						userLiveMvDeckInfo.MemberMasterID9 = memberId
					case 10:
						userLiveMvDeckInfo.MemberMasterID10 = memberId
					case 11:
						userLiveMvDeckInfo.MemberMasterID11 = memberId
					case 12:
						userLiveMvDeckInfo.MemberMasterID12 = memberId
					}
				}
			}
			// fmt.Println(memberIds)

			suitIds := map[int]int{}
			for k, v := range saveReq.SuitMasterIDByPos {
				if k%2 == 0 {
					suitId := saveReq.SuitMasterIDByPos[k+1]
					suitIds[v] = suitId

					switch v {
					case 1:
						userLiveMvDeckInfo.SuitMasterID1 = suitId
					case 2:
						userLiveMvDeckInfo.SuitMasterID2 = suitId
					case 3:
						userLiveMvDeckInfo.SuitMasterID3 = suitId
					case 4:
						userLiveMvDeckInfo.SuitMasterID4 = suitId
					case 5:
						userLiveMvDeckInfo.SuitMasterID5 = suitId
					case 6:
						userLiveMvDeckInfo.SuitMasterID6 = suitId
					case 7:
						userLiveMvDeckInfo.SuitMasterID7 = suitId
					case 8:
						userLiveMvDeckInfo.SuitMasterID8 = suitId
					case 9:
						userLiveMvDeckInfo.SuitMasterID9 = suitId
					case 10:
						userLiveMvDeckInfo.SuitMasterID10 = suitId
					case 11:
						userLiveMvDeckInfo.SuitMasterID11 = suitId
					case 12:
						userLiveMvDeckInfo.SuitMasterID12 = suitId
					}
				}
			}
			// fmt.Println(suitIds)

			var userLiveMvDeckCustomByID []model.UserLiveMvDeckCustomByID
			userLiveMvDeckCustomByID = append(userLiveMvDeckCustomByID, saveReq.LiveMasterID)
			userLiveMvDeckCustomByID = append(userLiveMvDeckCustomByID, userLiveMvDeckInfo)
			// fmt.Println(userLiveMvDeckCustomByID)

			viewStatusIds := map[int]int{}
			for k, v := range saveReq.ViewStatusByPos {
				if k%2 == 0 {
					viewStatusId := saveReq.ViewStatusByPos[k+1]
					viewStatusIds[v] = viewStatusId
				}
			}
			// fmt.Println(viewStatusIds)

			var userMemberByMemberID []model.UserMemberByMemberID
			for k, v := range memberIds {
				userMemberByMemberID = append(userMemberByMemberID, v)
				userMemberByMemberID = append(userMemberByMemberID, model.UserMemberInfo{
					MemberMasterID:           v,
					CustomBackgroundMasterID: 103506600,
					SuitMasterID:             suitIds[k],
					LovePoint:                0,
					LovePointLimit:           999999,
					LoveLevel:                1,
					ViewStatus:               viewStatusIds[k],
					IsNew:                    true,
				})
			}
			// fmt.Println(userMemberByMemberID)

			saveResp := model.LiveSaveDeckResp{
				UserModel: model.UserModel{
					UserStatus:                                              CommonUserStatus(),
					UserMemberByMemberID:                                    userMemberByMemberID,
					UserCardByCardID:                                        []interface{}{},
					UserSuitBySuitID:                                        []interface{}{},
					UserLiveDeckByID:                                        []interface{}{},
					UserLivePartyByID:                                       []interface{}{},
					UserLessonDeckByID:                                      []interface{}{},
					UserLiveMvDeckByID:                                      []interface{}{},
					UserLiveMvDeckCustomByID:                                userLiveMvDeckCustomByID,
					UserLiveDifficultyByDifficultyID:                        []interface{}{},
					UserStoryMainByStoryMainID:                              []interface{}{},
					UserStoryMainSelectedByStoryMainCellID:                  []interface{}{},
					UserVoiceByVoiceID:                                      []interface{}{},
					UserEmblemByEmblemID:                                    []interface{}{},
					UserGachaTicketByTicketID:                               []interface{}{},
					UserGachaPointByPointID:                                 []interface{}{},
					UserLessonEnhancingItemByItemID:                         []interface{}{},
					UserTrainingMaterialByItemID:                            []interface{}{},
					UserGradeUpItemByItemID:                                 []interface{}{},
					UserCustomBackgroundByID:                                []interface{}{},
					UserStorySideByID:                                       []interface{}{},
					UserStoryMemberByID:                                     []interface{}{},
					UserCommunicationMemberDetailBadgeByID:                  []interface{}{},
					UserStoryEventHistoryByID:                               []interface{}{},
					UserRecoveryLpByID:                                      []interface{}{},
					UserRecoveryApByID:                                      []interface{}{},
					UserMissionByMissionID:                                  []interface{}{},
					UserDailyMissionByMissionID:                             []interface{}{},
					UserWeeklyMissionByMissionID:                            []interface{}{},
					UserInfoTriggerBasicByTriggerID:                         []interface{}{},
					UserInfoTriggerCardGradeUpByTriggerID:                   []interface{}{},
					UserInfoTriggerMemberGuildSupportItemExpiredByTriggerID: []interface{}{},
					UserInfoTriggerMemberLoveLevelUpByTriggerID:             []interface{}{},
					UserAccessoryByUserAccessoryID:                          []interface{}{},
					UserAccessoryLevelUpItemByID:                            []interface{}{},
					UserAccessoryRarityUpItemByID:                           []interface{}{},
					UserUnlockScenesByEnum:                                  []interface{}{},
					UserSceneTipsByEnum:                                     []interface{}{},
					UserRuleDescriptionByID:                                 []interface{}{},
					UserExchangeEventPointByID:                              []interface{}{},
					UserSchoolIdolFestivalIDRewardMissionByID:               []interface{}{},
					UserGpsPresentReceivedByID:                              []interface{}{},
					UserEventMarathonByEventMasterID:                        []interface{}{},
					UserEventMiningByEventMasterID:                          []interface{}{},
					UserEventCoopByEventMasterID:                            []interface{}{},
					UserLiveSkipTicketByID:                                  []interface{}{},
					UserStoryEventUnlockItemByID:                            []interface{}{},
					UserEventMarathonBoosterByID:                            []interface{}{},
					UserReferenceBookByID:                                   []interface{}{},
					UserReviewRequestProcessFlowByID:                        []interface{}{},
					UserRankExpByID:                                         []interface{}{},
					UserShareByID:                                           []interface{}{},
					UserTowerByTowerID:                                      []interface{}{},
					UserRecoveryTowerCardUsedCountItemByRecoveryTowerCardUsedCountItemMasterID: []interface{}{},
					UserStoryLinkageByID:             []interface{}{},
					UserSubscriptionStatusByID:       []interface{}{},
					UserStoryMainPartDigestMovieByID: []interface{}{},
					UserMemberGuildByID:              []interface{}{},
					UserMemberGuildSupportItemByID:   []interface{}{},
					UserDailyTheaterByDailyTheaterID: []interface{}{},
					UserPlayListByID:                 []interface{}{},
				},
			}
			respB, err := json.Marshal(saveResp)
			if err != nil {
				panic(err)
			}
			// fmt.Println(string(b))

			resp := SignResp(ctx.GetString("ep"), string(respB), sessionKey)
			// fmt.Println(resp)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/communicationMember/fetchCommunicationMemberDetail", func(ctx *gin.Context) {
			reqBody := ctx.GetString("reqBody")
			var memberId string
			gjson.Parse(reqBody).ForEach(func(key, value gjson.Result) bool {
				if value.Get("member_id").String() != "" {
					memberId = value.Get("member_id").String()
					return false
				}
				return true
			})

			signBody := strings.ReplaceAll(utils.ReadAllText("assets/as/fetchCommunicationMemberDetail.json"), `"MEMBER_ID"`, memberId)
			resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/navi/tapLovePoint", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/tapLovePoint.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/communicationMember/updateUserCommunicationMemberDetailBadge", func(ctx *gin.Context) {
			reqBody := ctx.GetString("reqBody")
			var memberMasterId string
			gjson.Parse(reqBody).ForEach(func(key, value gjson.Result) bool {
				if value.Get("member_master_id").String() != "" {
					memberMasterId = value.Get("member_master_id").String()
					return false
				}
				return true
			})

			signBody := strings.ReplaceAll(utils.ReadAllText("assets/as/updateUserCommunicationMemberDetailBadge.json"), `"MEMBER_MASTER_ID"`, memberMasterId)
			resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/communicationMember/updateUserLiveDifficultyNewFlag", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/updateUserLiveDifficultyNewFlag.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/communicationMember/finishUserStorySide", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/finishUserStorySide.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/communicationMember/finishUserStoryMember", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/finishUserStoryMember.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/communicationMember/setTheme", func(ctx *gin.Context) {
			reqBody := ctx.GetString("reqBody")
			var memberMasterId, suitMasterId, backgroundMasterId string
			gjson.Parse(reqBody).ForEach(func(key, value gjson.Result) bool {
				if value.Get("member_master_id").String() != "" {
					memberMasterId = value.Get("member_master_id").String()
					suitMasterId = value.Get("suit_master_id").String()
					backgroundMasterId = value.Get("custom_background_master_id").String()
					return false
				}
				return true
			})

			signBody := strings.ReplaceAll(utils.ReadAllText("assets/as/setTheme.json"), `"MEMBER_MASTER_ID"`, memberMasterId)
			signBody = strings.ReplaceAll(signBody, `"BACKGROUND_MASTER_ID"`, backgroundMasterId)
			signBody = strings.ReplaceAll(signBody, `"SUIT_MASTER_ID"`, suitMasterId)
			resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/userProfile/fetchProfile", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/fetchProfile.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/emblem/fetchEmblem", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/fetchEmblem.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/emblem/activateEmblem", func(ctx *gin.Context) {
			reqBody := ctx.GetString("reqBody")
			var emblemId string
			gjson.Parse(reqBody).ForEach(func(key, value gjson.Result) bool {
				if value.Get("emblem_master_id").String() != "" {
					emblemId = value.Get("emblem_master_id").String()
					return false
				}
				return true
			})
			signBody := strings.ReplaceAll(utils.ReadAllText("assets/as/activateEmblem.json"), `"EMBLEM_ID"`, emblemId)
			resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/navi/saveUserNaviVoice", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/saveUserNaviVoice.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/liveDeck/saveDeckAll", func(ctx *gin.Context) {
			reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0]
			// fmt.Println(reqBody.String())

			req := model.AsSaveDeckReq{}
			decoder := json.NewDecoder(strings.NewReader(reqBody.String()))
			decoder.UseNumber()
			err := decoder.Decode(&req)
			CheckErr(err)
			// fmt.Println("Raw:", req.SquadDict)

			if req.CardWithSuit[1] == 0 {
				req.CardWithSuit[1] = req.CardWithSuit[0]
			}
			if req.CardWithSuit[3] == 0 {
				req.CardWithSuit[3] = req.CardWithSuit[2]
			}
			if req.CardWithSuit[5] == 0 {
				req.CardWithSuit[5] = req.CardWithSuit[4]
			}
			if req.CardWithSuit[7] == 0 {
				req.CardWithSuit[7] = req.CardWithSuit[6]
			}
			if req.CardWithSuit[9] == 0 {
				req.CardWithSuit[9] = req.CardWithSuit[8]
			}
			if req.CardWithSuit[11] == 0 {
				req.CardWithSuit[11] = req.CardWithSuit[10]
			}
			if req.CardWithSuit[13] == 0 {
				req.CardWithSuit[13] = req.CardWithSuit[12]
			}
			if req.CardWithSuit[15] == 0 {
				req.CardWithSuit[15] = req.CardWithSuit[14]
			}
			if req.CardWithSuit[17] == 0 {
				req.CardWithSuit[17] = req.CardWithSuit[16]
			}

			deckInfo := model.AsDeckInfo{
				UserLiveDeckID: req.DeckID,
				Name: model.AsDeckName{
					DotUnderText: "队伍名称变了?",
				},
				CardMasterID1: req.CardWithSuit[0],
				CardMasterID2: req.CardWithSuit[2],
				CardMasterID3: req.CardWithSuit[4],
				CardMasterID4: req.CardWithSuit[6],
				CardMasterID5: req.CardWithSuit[8],
				CardMasterID6: req.CardWithSuit[10],
				CardMasterID7: req.CardWithSuit[12],
				CardMasterID8: req.CardWithSuit[14],
				CardMasterID9: req.CardWithSuit[16],
				SuitMasterID1: req.CardWithSuit[1],
				SuitMasterID2: req.CardWithSuit[3],
				SuitMasterID3: req.CardWithSuit[5],
				SuitMasterID4: req.CardWithSuit[7],
				SuitMasterID5: req.CardWithSuit[9],
				SuitMasterID6: req.CardWithSuit[11],
				SuitMasterID7: req.CardWithSuit[13],
				SuitMasterID8: req.CardWithSuit[15],
				SuitMasterID9: req.CardWithSuit[17],
			}
			// fmt.Println(deckInfo)

			deckInfoRes := []model.AsResp{}
			deckInfoRes = append(deckInfoRes, req.DeckID)
			deckInfoRes = append(deckInfoRes, deckInfo)

			partyInfoRes := []model.AsResp{}
			for k, v := range req.SquadDict {
				if k%2 == 0 {
					partyId, err := v.(json.Number).Int64()
					CheckErr(err)
					// fmt.Println("Party ID:", partyId)

					rDictInfo, err := json.Marshal(req.SquadDict[k+1])
					CheckErr(err)

					dictInfo := model.AsDeckSquadDict{}
					if err = json.Unmarshal(rDictInfo, &dictInfo); err != nil {
						panic(err)
					}
					// fmt.Println("Party Info:", dictInfo)

					roleIds := []int{}
					err = mEng.Table("m_card").
						Where("id IN (?,?,?)", dictInfo.CardMasterIds[0], dictInfo.CardMasterIds[1], dictInfo.CardMasterIds[2]).
						Cols("role").Find(&roleIds)
					CheckErr(err)
					// fmt.Println("roleIds:", roleIds)

					var partyIcon int
					var partyName string
					// 脑残逻辑部分
					exists, err := mEng.Table("m_live_party_name").
						Where("role_1 = ? AND role_2 = ? AND role_3 = ?", roleIds[0], roleIds[1], roleIds[2]).
						Cols("name,live_party_icon").Get(&partyName, &partyIcon)
					CheckErr(err)
					if !exists {
						exists, err = mEng.Table("m_live_party_name").
							Where("role_1 = ? AND role_2 = ? AND role_3 = ?", roleIds[0], roleIds[2], roleIds[1]).
							Cols("name,live_party_icon").Get(&partyName, &partyIcon)
						CheckErr(err)
						if !exists {
							exists, err = mEng.Table("m_live_party_name").
								Where("role_1 = ? AND role_2 = ? AND role_3 = ?", roleIds[1], roleIds[0], roleIds[2]).
								Cols("name,live_party_icon").Get(&partyName, &partyIcon)
							CheckErr(err)
							if !exists {
								exists, err = mEng.Table("m_live_party_name").
									Where("role_1 = ? AND role_2 = ? AND role_3 = ?", roleIds[1], roleIds[2], roleIds[0]).
									Cols("name,live_party_icon").Get(&partyName, &partyIcon)
								CheckErr(err)
								if !exists {
									exists, err = mEng.Table("m_live_party_name").
										Where("role_1 = ? AND role_2 = ? AND role_3 = ?", roleIds[2], roleIds[0], roleIds[1]).
										Cols("name,live_party_icon").Get(&partyName, &partyIcon)
									CheckErr(err)
									if !exists {
										exists, err = mEng.Table("m_live_party_name").
											Where("role_1 = ? AND role_2 = ? AND role_3 = ?", roleIds[2], roleIds[1], roleIds[0]).
											Cols("name,live_party_icon").Get(&partyName, &partyIcon)
										CheckErr(err)
										if !exists {
											panic("Fuck you!")
										}
									}
								}
							}
						}
					}

					var realPartyName string
					_, err = mEng.Table("m_dictionary").Where("id = ?", strings.ReplaceAll(partyName, "k.", "")).Cols("message").Get(&realPartyName)
					CheckErr(err)

					partyInfo := model.AsPartyInfo{
						PartyID:        int(partyId),
						UserLiveDeckID: req.DeckID,
						Name: model.AsPartyName{
							DotUnderText: realPartyName,
						},
						IconMasterID:     partyIcon,
						CardMasterID1:    dictInfo.CardMasterIds[0],
						CardMasterID2:    dictInfo.CardMasterIds[1],
						CardMasterID3:    dictInfo.CardMasterIds[2],
						UserAccessoryID1: dictInfo.UserAccessoryIds[0],
						UserAccessoryID2: dictInfo.UserAccessoryIds[1],
						UserAccessoryID3: dictInfo.UserAccessoryIds[2],
					}
					// fmt.Println(partyInfo)

					partyInfoRes = append(partyInfoRes, partyId)
					partyInfoRes = append(partyInfoRes, partyInfo)
				}
			}

			// fmt.Println(deckInfoRes)
			// fmt.Println(partyInfoRes)

			m1, err := json.Marshal(deckInfoRes)
			CheckErr(err)

			m2, err := json.Marshal(partyInfoRes)
			CheckErr(err)

			respBody := utils.ReadAllText("assets/as/saveDeckAll.json")
			respBody = strings.ReplaceAll(respBody, `"DECK_INFO"`, string(m1))
			respBody = strings.ReplaceAll(respBody, `"PARTY_INFO"`, string(m2))
			// fmt.Println(respBody)

			resp := SignResp(ctx.GetString("ep"), respBody, sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/livePartners/fetch", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/fetchLivePartners.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/liveDeck/fetchLiveDeckSelect", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/fetchLiveDeckSelect.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/live/start", func(ctx *gin.Context) {
			reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0]
			// fmt.Println(reqBody.String())

			liveStartReq := model.AsLiveStartReq{}
			if err := json.Unmarshal([]byte(reqBody.String()), &liveStartReq); err != nil {
				panic(err)
			}
			// fmt.Println(liveStartReq)

			var cardInfo string
			partnerResp := gjson.Parse(utils.ReadAllText("assets/as/fetchLivePartners.json")).Get("partner_select_state.live_partners")
			partnerResp.ForEach(func(k, v gjson.Result) bool {
				userId := v.Get("user_id").Int()
				if userId == int64(liveStartReq.PartnerUserID) {
					v.Get("card_by_category").ForEach(func(kk, vv gjson.Result) bool {
						if vv.IsObject() {
							cardId := vv.Get("card_master_id").Int()
							if cardId == int64(liveStartReq.PartnerCardMasterID) {
								cardInfo = vv.String()
								// fmt.Println(cardInfo)
								return false
							}
						}
						return true
					})
					return false
				}
				return true
			})

			// 保存请求包因为 /live/finish 接口的响应包里有部分字段不在该接口的请求包里
			liveId := strconv.Itoa(int(time.Now().UnixNano()))
			err := database.LevelDb.Put([]byte("live_"+liveId), []byte(reqBody.String()))
			CheckErr(err)

			liveDifficultyId := strconv.Itoa(liveStartReq.LiveDifficultyID)
			liveNotes := utils.ReadAllText("assets/as/notes/" + liveDifficultyId + ".json")
			if liveNotes == "" {
				panic("歌曲情报信息不存在！")
			}

			liveStartResp := utils.ReadAllText("assets/as/liveStart.json")
			liveStartResp = strings.ReplaceAll(liveStartResp, `"LIVE_ID"`, liveId)
			liveStartResp = strings.ReplaceAll(liveStartResp, `"DECK_ID"`, strconv.Itoa(liveStartReq.DeckID))
			liveStartResp = strings.ReplaceAll(liveStartResp, `"LIVE_NOTES"`, liveNotes)
			liveStartResp = strings.ReplaceAll(liveStartResp, `"LIVE_PARTNER"`, cardInfo)
			liveStartResp = strings.ReplaceAll(liveStartResp, `"LIVE_DIFFICULTY_ID"`, liveDifficultyId)

			resp := SignResp(ctx.GetString("ep"), liveStartResp, sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/live/finish", func(ctx *gin.Context) {
			reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0]
			// fmt.Println(reqBody.String())

			var cardMasterId, maxVolt, skillCount, appealCount int64
			liveFinishReq := gjson.Parse(reqBody.String())
			liveFinishReq.Get("live_score.card_stat_dict").ForEach(func(key, value gjson.Result) bool {
				if value.IsObject() {
					volt := value.Get("got_voltage").Int()
					if volt > maxVolt {
						maxVolt = volt

						cardMasterId = value.Get("card_master_id").Int()
						skillCount = value.Get("skill_triggered_count").Int()
						appealCount = value.Get("appeal_count").Int()
					}
				}
				return true
			})

			mvpInfo := model.AsMvpInfo{
				CardMasterID:        cardMasterId,
				GetVoltage:          maxVolt,
				SkillTriggeredCount: skillCount,
				AppealCount:         appealCount,
			}
			mvp, err := json.Marshal(mvpInfo)
			CheckErr(err)
			// fmt.Println("mvpInfo:", mvpInfo)

			liveId := liveFinishReq.Get("live_id").String()
			res, err := database.LevelDb.Get([]byte("live_" + liveId))
			CheckErr(err)

			liveStartReq := model.AsLiveStartReq{}
			if err := json.Unmarshal(res, &liveStartReq); err != nil {
				panic(err)
			}
			// fmt.Println("liveStartReq:", liveStartReq)

			partnerInfo := model.AsLivePartnerInfo{
				LastPlayedAt:                        time.Now().Unix(),
				RecommendCardMasterID:               liveStartReq.PartnerCardMasterID,
				RecommendCardLevel:                  1,
				IsRecommendCardImageAwaken:          true,
				IsRecommendCardAllTrainingActivated: true,
				IsNew:                               false,
				FriendApprovedAt:                    nil,
				RequestStatus:                       3,
				IsRequestPending:                    false,
			}
			partnerResp := gjson.Parse(utils.ReadAllText("assets/as/fetchLivePartners.json")).Get("partner_select_state.live_partners")
			partnerResp.ForEach(func(k, v gjson.Result) bool {
				userId := v.Get("user_id").Int()
				if userId == int64(liveStartReq.PartnerUserID) {
					partnerInfo.UserID = int(userId)
					partnerInfo.Name.DotUnderText = v.Get("name.dot_under_text").String()
					partnerInfo.Rank = int(v.Get("rank").Int())
					partnerInfo.EmblemID = int(v.Get("emblem_id").Int())
					partnerInfo.IntroductionMessage.DotUnderText = v.Get("introduction_message.dot_under_text").String()
				}
				return true
			})
			partner, err := json.Marshal(partnerInfo)
			CheckErr(err)
			// fmt.Println(partnerInfo)

			liveResultResp := model.AsLiveResultAchievementStatus{
				ClearCount:       1,
				GotVoltage:       liveFinishReq.Get("live_score.current_score").Int(),
				RemainingStamina: liveFinishReq.Get("live_score.remaining_stamina").Int(),
			}
			liveResult, err := json.Marshal(liveResultResp)
			CheckErr(err)
			// fmt.Println(liveResultResp)

			liveFinishResp := utils.ReadAllText("assets/as/liveFinish.json")
			liveFinishResp = strings.ReplaceAll(liveFinishResp, `"LIVE_DIFFICULTY_ID"`, strconv.Itoa(liveStartReq.LiveDifficultyID))
			liveFinishResp = strings.ReplaceAll(liveFinishResp, `"DECK_ID"`, strconv.Itoa(liveStartReq.DeckID))
			liveFinishResp = strings.ReplaceAll(liveFinishResp, `"MVP_INFO"`, string(mvp))
			liveFinishResp = strings.ReplaceAll(liveFinishResp, `"PARTNER_INFO"`, string(partner))
			liveFinishResp = strings.ReplaceAll(liveFinishResp, `"LIVE_RESULT"`, string(liveResult))
			liveFinishResp = strings.ReplaceAll(liveFinishResp, `"LIVE_VOLTAGE"`, liveFinishReq.Get("live_score.current_score").String())
			// fmt.Println(liveFinishResp)

			resp := SignResp(ctx.GetString("ep"), liveFinishResp, sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
	}
}
