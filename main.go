package main

import (
	"honoka-chan/handler"
	_ "honoka-chan/llhelper"
	"honoka-chan/middleware"
	_ "honoka-chan/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {
	// Router
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("static/templates/**/*.html")

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

	// Web
	// Manga
	r.GET("/manga", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "common/manga.html", gin.H{})
	})

	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
