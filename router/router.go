package router

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/config"
	"honoka-chan/encrypt"
	"honoka-chan/handler"
	"honoka-chan/middleware"
	"honoka-chan/model"
	"honoka-chan/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

var (
	sessionKey = "12345678123456781234567812345678"
)

func SifRouter(r *gin.Engine) {
	// Static
	r.Static("/static", "static")

	var files []string
	filepath.Walk("static/templates", func(path string, info os.FileInfo, err error) error {
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
			body, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				panic(err)
			}
			defer ctx.Request.Body.Close()

			var mask string
			req := gjson.Parse(string(body))
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

			loginBody := strings.ReplaceAll(utils.ReadAllText("assets/login.json"), "SESSION_KEY", newKey64)
			resp := SignResp(ctx.GetString("ep"), loginBody, config.StartUpKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/bootstrap/fetchBootstrap", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/bootstrap.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/billing/fetchBillingHistory", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/billing.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/notice/fetchNotice", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/notice.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/asset/getPackUrl", func(ctx *gin.Context) {
			body, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				panic(err)
			}
			defer ctx.Request.Body.Close()
			// fmt.Println(string(body))

			req := []model.AsReq{}
			err = json.Unmarshal(body, &req)
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
			packageList := []string{}
			urlList := []string{}

			err = json.Unmarshal([]byte(utils.ReadAllText("assets/packages.json")), &packageList)
			if err != nil {
				panic(err)
			}

			err = json.Unmarshal([]byte(utils.ReadAllText("assets/urls.json")), &urlList)
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
			respUrls := []string{}
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

			resp := []model.AsResp{}
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
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/updateCardNewFlag.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/bootstrap/getClearedPlatformAchievement", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/getClearedPlatformAchievement.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/live/fetchLiveMusicSelect", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/fetchLiveMusicSelect.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/liveMv/start", func(ctx *gin.Context) {
			resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/liveMvStart.json"), sessionKey)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, resp)
		})
		s.POST("/liveMv/saveDeck", func(ctx *gin.Context) {
			body, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				panic(err)
			}
			defer ctx.Request.Body.Close()
			// fmt.Println(string(body))

			req := []model.AsReq{}
			err = json.Unmarshal(body, &req)
			if err != nil {
				panic(err)
			}

			reqBody, ok := req[0].(map[string]interface{})
			if !ok {
				panic("Assertion failed!")
			}

			reqB, err := json.Marshal(reqBody)
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

			userLiveMvDeckCustomByID := []model.UserLiveMvDeckCustomByID{}
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

			userMemberByMemberID := []model.UserMemberByMemberID{}
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
	}
}
