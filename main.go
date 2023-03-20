package main

import (
	"honoka-chan/config"
	"honoka-chan/handler"
	"honoka-chan/middleware"
	"honoka-chan/sifcap"
	"honoka-chan/xclog"

	"github.com/gin-gonic/gin"
)

func init() {
	xclog.Init(config.Conf.Log.LogDir, "", config.Conf.Log.LogLevel, config.Conf.Log.LogSave)
}

func main() {
	// test
	// apiData := utils.ReadAllText("data/api1.json")
	// var obj model.Response
	// err := json.Unmarshal([]byte(apiData), &obj)
	// if err != nil {
	// 	panic(err)
	// }

	// var data interface{}
	// err = json.Unmarshal(obj.ResponseData, &data)
	// if err != nil {
	// 	panic(err)
	// }
	// resultType := reflect.TypeOf(data)
	// // fmt.Println(resultType.Kind())
	// if resultType.Kind() == reflect.Map {
	// 	data = data.(map[string]interface{})
	// }
	// result := data.([]interface{})
	// for k, v := range result {
	// 	m, err := json.Marshal(v)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	switch k {
	// 	case 0:
	// 		res := model.LiveStatusResp{}
	// 		err = json.Unmarshal(m, &res)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		// fmt.Println(res.Result.NormalLiveStatusList[0].HiScore)
	// 	case 1:
	// 		res := model.LiveScheduleResp{}
	// 		err = json.Unmarshal(m, &res)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		// fmt.Println(res.Result.LiveList[0].StartDate)
	// 	case 2:
	// 		res := model.UnitAllResp{}
	// 		err = json.Unmarshal(m, &res)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		for _, vv := range res.Result.Active {
	// 			if vv.MaxLevel == 120 {
	// 				fmt.Println(vv)
	// 			}
	// 		}
	// 	}
	// }

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
		r.POST("/main.php/download/event", handler.DownloadEventHandler)
		r.POST("/main.php/lbonus/execute", handler.LBonusExecuteHandler)
		r.POST("/main.php/api", handler.ApiHandler)
		r.POST("/main.php/announce/checkState", handler.AnnounceCheckStateHandler)
		r.POST("/main.php/scenario/startup", handler.ScenarioStartupHandler)
		r.POST("/main.php/user/setNotificationToken", handler.SetNotificationTokenHandler)

		r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	}
}
