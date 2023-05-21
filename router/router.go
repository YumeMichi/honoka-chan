package router

import (
	"encoding/base64"
	"fmt"
	"honoka-chan/config"
	"honoka-chan/encrypt"
	"honoka-chan/handler"
	"honoka-chan/middleware"
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

	s := r.Group("ep3071")
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
			fmt.Println("Request data:", req.String())
			fmt.Println("Mask:", mask)

			mask64, err := base64.StdEncoding.DecodeString(mask)
			if err != nil {
				panic(err)
			}
			randomBytes := encrypt.RSA_DecryptOAEP(mask64, "privatekey.pem")
			fmt.Println("Random Bytes:", randomBytes)

			newKey := utils.SliceXor(randomBytes, []byte(sessionKey))
			newKey64 := base64.StdEncoding.EncodeToString(newKey)

			loginBody := strings.ReplaceAll(utils.ReadAllText("data/login_body.txt"), "SESSION_KEY", newKey64)
			signBody := fmt.Sprintf("%d,\"%s\",0,%s", time.Now().UnixMilli(), config.MasterVersion, loginBody)

			ep := strings.ReplaceAll(ctx.Request.URL.String(), "/ep3071", "")
			fmt.Println(ep)

			sign := encrypt.HMAC_SHA1_Encrypt([]byte(ep+" "+signBody), []byte(config.StartUpKey))
			fmt.Println(sign)

			res := fmt.Sprintf("[%s,\"%s\"]", signBody, sign)
			// fmt.Println(res)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, res)
		})
		s.POST("/bootstrap/fetchBootstrap", func(ctx *gin.Context) {
			body, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				panic(err)
			}
			defer ctx.Request.Body.Close()
			fmt.Println(string(body))

			ep := strings.ReplaceAll(ctx.Request.URL.String(), "/ep3071", "")
			fmt.Println(ep)

			bootstrapBody := utils.ReadAllText("data/bootstrap_body.txt")
			signBody := fmt.Sprintf("%d,\"%s\",0,%s", time.Now().UnixMilli(), config.MasterVersion, bootstrapBody)
			sign := encrypt.HMAC_SHA1_Encrypt([]byte(ep+" "+signBody), []byte(sessionKey))
			fmt.Println(sign)

			res := fmt.Sprintf("[%s,\"%s\"]", signBody, sign)
			// fmt.Println(res)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, res)
		})
		s.POST("/billing/fetchBillingHistory", func(ctx *gin.Context) {
			body, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				panic(err)
			}
			defer ctx.Request.Body.Close()
			fmt.Println(string(body))

			ep := strings.ReplaceAll(ctx.Request.URL.String(), "/ep3071", "")
			fmt.Println(ep)

			bootstrapBody := utils.ReadAllText("data/bill_body.txt")
			signBody := fmt.Sprintf("%d,\"%s\",0,%s", time.Now().UnixMilli(), config.MasterVersion, bootstrapBody)
			sign := encrypt.HMAC_SHA1_Encrypt([]byte(ep+" "+signBody), []byte(sessionKey))
			fmt.Println(sign)

			res := fmt.Sprintf("[%s,\"%s\"]", signBody, sign)
			// fmt.Println(res)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, res)
		})
		s.POST("/notice/fetchNotice", func(ctx *gin.Context) {
			body, err := io.ReadAll(ctx.Request.Body)
			if err != nil {
				panic(err)
			}
			defer ctx.Request.Body.Close()
			fmt.Println(string(body))

			ep := strings.ReplaceAll(ctx.Request.URL.String(), "/ep3071", "")
			fmt.Println(ep)

			bootstrapBody := utils.ReadAllText("data/notice_body.txt")
			signBody := fmt.Sprintf("%d,\"%s\",0,%s", time.Now().UnixMilli(), config.MasterVersion, bootstrapBody)
			sign := encrypt.HMAC_SHA1_Encrypt([]byte(ep+" "+signBody), []byte(sessionKey))
			fmt.Println(sign)

			res := fmt.Sprintf("[%s,\"%s\"]", signBody, sign)
			// fmt.Println(res)

			ctx.Header("Content-Type", "application/json")
			ctx.String(http.StatusOK, res)
		})
	}
}
