package main

import (
	"honoka-chan/config"
	"honoka-chan/handler"
	_ "honoka-chan/llhelper"
	"honoka-chan/middleware"
	_ "honoka-chan/tools"
	"honoka-chan/xclog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	xclog.Init(config.Conf.Log.LogDir, "", config.Conf.Log.LogLevel, config.Conf.Log.LogSave)
}

func main() {
	// Router
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("static/*.tmpl")

	// /
	r.Any("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello, world!")
	})

	// Private APIs
	v1 := r.Group("v1")
	{
		v1.GET("/basic/getcode", handler.GetCodeHandler)
		v1.POST("/basic/getcode", handler.GetCodeHandler)
		v1.POST("/account/active", handler.ActiveHandler)
		v1.POST("/basic/publickey", handler.PublicKeyHandler)
		v1.POST("/basic/handshake", handler.HandshakeHandler)
		v1.POST("/account/initialize", handler.InitializeHandler)
		v1.POST("/account/login", handler.AccountLoginHandler)
		v1.POST("/account/loginauto", handler.LoginAutoHandler)
		v1.POST("/basic/loginarea", handler.LoginAreaHandler)
		v1.POST("/account/reportRole", handler.ReportRoleHandler)
		v1.POST("/basic/getProductList", handler.GetProductListHandler)
		v1.POST("/guest/status", handler.GuestStatusHandler)
	}
	r.GET("/agreement/all", handler.AgreementHandler)
	r.GET("/integration/appReport/initialize", handler.ReportApp)
	r.POST("/report/ge/app", handler.ReportLog)
	// Private APIs

	// Server APIs
	m := r.Group("main.php").Use(middleware.Common)
	{
		m.POST("/album/seriesAll", middleware.ParseMultipartForm, handler.AlbumSeriesAllHandler)
		m.POST("/announce/checkState", middleware.ParseMultipartForm, handler.AnnounceCheckStateHandler)
		m.POST("/api", middleware.ParseMultipartForm, handler.ApiHandler)
		m.POST("/award/set", handler.AwardSet)
		m.POST("/background/set", handler.BackgroundSet)
		m.POST("/download/additional", handler.DownloadAdditionalHandler)
		m.POST("/download/batch", handler.DownloadBatchHandler)
		m.POST("/download/event", handler.DownloadEventHandler)
		m.POST("/download/getUrl", handler.DownloadUrlHandler)
		m.POST("/download/update", handler.DownloadUpdateHandler)
		m.POST("/event/eventList", middleware.ParseMultipartForm, handler.EventListHandler)
		m.POST("/gdpr/get", middleware.ParseMultipartForm, handler.GdprHandler)
		m.POST("/lbonus/execute", handler.LBonusExecuteHandler)
		m.POST("/live/gameover", handler.GameOverHandler)
		m.POST("/live/partyList", handler.PartyListHandler)
		m.POST("/live/play", middleware.ParseMultipartForm, handler.PlayLiveHandler)
		m.POST("/live/preciseScore", middleware.ParseMultipartForm, handler.PlayScoreHandler)
		m.POST("/live/reward", middleware.ParseMultipartForm, handler.PlayRewardHandler)
		m.POST("/login/authkey", middleware.AuthKey, handler.AuthKey)
		m.POST("/login/login", middleware.Login, handler.Login)
		m.POST("/multiunit/scenarioStartup", handler.MultiUnitStartUpHandler)
		m.POST("/museum/info", middleware.ParseMultipartForm, handler.MuseumInfo)
		m.POST("/notice/noticeFriendGreeting", middleware.ParseMultipartForm, handler.NoticeFriendGreetingHandler)
		m.POST("/notice/noticeFriendVariety", middleware.ParseMultipartForm, handler.NoticeFriendVarietyHandler)
		m.POST("/notice/noticeUserGreetingHistory", handler.NoticeUserGreetingHandler)
		m.POST("/payment/productList", middleware.ParseMultipartForm, handler.ProductListHandler)
		m.POST("/personalnotice/get", middleware.ParseMultipartForm, handler.PersonalNoticeHandler)
		m.POST("/profile/profileRegister", handler.ProfileRegister)
		m.POST("/scenario/reward", handler.ScenarioRewardHandler)
		m.POST("/scenario/startup", handler.ScenarioStartupHandler)
		m.POST("/subscenario/reward", handler.SubScenarioStartupHandler)
		m.POST("/subscenario/startup", handler.SubScenarioStartupHandler)
		m.POST("/tos/tosCheck", middleware.ParseMultipartForm, handler.TosCheckHandler)
		m.POST("/unit/deck", handler.SetDeckHandler)
		m.POST("/unit/deckName", handler.SetDeckNameHandler)
		m.POST("/unit/favorite", handler.SetDisplayRankHandler)
		m.POST("/unit/removableSkillEquipment", handler.RemoveSkillEquip)
		m.POST("/unit/setDisplayRank", handler.SetDisplayRankHandler)
		m.POST("/unit/wearAccessory", handler.WearAccessory)
		m.POST("/user/changeName", handler.ChangeNameHandler)
		m.POST("/user/changeNavi", handler.ChangeNaviHandler)
		m.POST("/user/setNotificationToken", handler.SetNotificationTokenHandler)
		m.POST("/user/userInfo", middleware.ParseMultipartForm, handler.UserInfoHandler)
	}
	r.GET("/webview.php/announce/index", handler.AnnounceIndexHandler)
	// Server APIs

	// Web
	// Manga
	r.GET("/manga", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "manga.tmpl", gin.H{})
	})

	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
