package main

import (
	"honoka-chan/config"
	"honoka-chan/handler"
	"honoka-chan/middleware"
	"honoka-chan/sifcap"
	"honoka-chan/tools"
	"honoka-chan/xclog"

	"github.com/gin-gonic/gin"
)

func init() {
	xclog.Init(config.Conf.Log.LogDir, "", config.Conf.Log.LogLevel, config.Conf.Log.LogSave)
}

func main() {
	tools.AnalysisApi1Data("assets/api1.json")
	tools.AnalysisApi2Data("assets/api2.json")
	tools.AnalysisApi3Data("assets/api3.json")

	if config.Conf.SifCap.Enabled {
		sifcap.Start()
	} else {
		// router
		r := gin.Default()
		r.Use(middleware.KlabHeader)

		r.GET("/webview.php/announce/index", handler.AnnounceIndexHandler)

		r.POST("/main.php/login/authkey", handler.AuthKeyHandler)
		r.POST("/main.php/login/login", handler.LoginHandler)
		r.POST("/main.php/user/userInfo", handler.UserInfoHandler)
		r.POST("/main.php/gdpr/get", handler.GdprHandler)
		r.POST("/main.php/personalnotice/get", handler.PersonalNoticeHandler)
		r.POST("/main.php/tos/tosCheck", handler.TosCheckHandler)
		r.POST("/main.php/download/batch", handler.DownloadBatchHandler)
		r.POST("/main.php/download/event", handler.DownloadEventHandler)
		r.POST("/main.php/lbonus/execute", handler.LBonusExecuteHandler)
		r.POST("/main.php/api", handler.ApiHandler)
		r.POST("/main.php/announce/checkState", handler.AnnounceCheckStateHandler)
		r.POST("/main.php/scenario/startup", handler.ScenarioStartupHandler)
		r.POST("/main.php/user/setNotificationToken", handler.SetNotificationTokenHandler)

		r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}
}
