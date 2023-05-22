package main

import (
	"honoka-chan/router"
	_ "honoka-chan/tools"

	"github.com/gin-gonic/gin"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	// eng, err := xorm.NewEngine("sqlite", "assets/masterdata.db")
	// CheckErr(err)
	// err = eng.Ping()
	// CheckErr(err)

	// cardIds := []int{}
	// err = eng.Table("m_card").Cols("id").OrderBy("id ASC").Find(&cardIds)
	// CheckErr(err)

	// jsonStr := "["
	// for _, card := range cardIds {
	// 	cardInfo := model.AsCardInfo{
	// 		CardMasterID:               card,
	// 		Level:                      1,
	// 		Exp:                        0,
	// 		LovePoint:                  0,
	// 		IsFavorite:                 false,
	// 		IsAwakening:                true,
	// 		IsAwakeningImage:           true,
	// 		IsAllTrainingActivated:     false,
	// 		TrainingActivatedCellCount: 0,
	// 		MaxFreePassiveSkill:        1,
	// 		Grade:                      0,
	// 		TrainingLife:               0,
	// 		TrainingAttack:             0,
	// 		TrainingDexterity:          0,
	// 		ActiveSkillLevel:           1,
	// 		PassiveSkillALevel:         1,
	// 		PassiveSkillBLevel:         1,
	// 		PassiveSkillCLevel:         1,
	// 		AdditionalPassiveSkill1ID:  0,
	// 		AdditionalPassiveSkill2ID:  0,
	// 		AdditionalPassiveSkill3ID:  0,
	// 		AdditionalPassiveSkill4ID:  0,
	// 		AcquiredAt:                 int(time.Now().Unix()),
	// 		IsNew:                      false,
	// 	}
	// 	m, err := json.Marshal(cardInfo)
	// 	CheckErr(err)

	// 	jsonStr += fmt.Sprintf("%d,%s,", card, string(m))
	// }
	// jsonStr = strings.TrimRight(jsonStr, ",")
	// jsonStr += "]"
	// fmt.Println(jsonStr)

	// packageList := []string{}
	// urlList := []string{}

	// err := json.Unmarshal([]byte(utils.ReadAllText("data/packages.json")), &packageList)
	// CheckErr(err)

	// err = json.Unmarshal([]byte(utils.ReadAllText("data/urls.json")), &urlList)
	// CheckErr(err)

	// if len(packageList) != len(urlList) {
	// 	fmt.Println("File size not match!")
	// 	return
	// }

	// packageUrls := map[string]string{}
	// for k, p := range packageList {
	// 	packageUrls[p] = urlList[k]
	// }

	// // fmt.Println(packageUrls)

	// // packUrlBody := model.PackUrlBody{
	// // 	PackNames: []string{
	// // 		"cxh9rj",
	// // 		"ib3f4p",
	// // 		"iwjqao",
	// // 	},
	// // }

	// // hash := "asdfgh"

	// // asReq := []model.AsReq{}
	// // asReq = append(asReq, packUrlBody)
	// // asReq = append(asReq, hash)
	// // mm, err := json.Marshal(asReq)
	// // CheckErr(err)
	// // fmt.Println(string(mm))

	// jsonStr := utils.ReadAllText("data/api1.json")
	// req := []model.AsReq{}
	// err = json.Unmarshal([]byte(jsonStr), &req)
	// CheckErr(err)
	// // fmt.Println(req)

	// packBody, ok := req[0].(map[string]interface{})
	// if !ok {
	// 	panic("Assertion failed!")
	// }
	// // fmt.Println(packBody)

	// packNames, ok := packBody["pack_names"].([]interface{})
	// if !ok {
	// 	panic("Assertion failed!")
	// }

	// respUrls := []string{}
	// for _, pack := range packNames {
	// 	packName, ok := pack.(string)
	// 	if !ok {
	// 		panic("Assertion failed!")
	// 	}
	// 	fmt.Println(packageUrls[packName])
	// 	respUrls = append(respUrls, packageUrls[packName])
	// }

	// urlResp := model.PackUrlRespBody{
	// 	UrlList: respUrls,
	// }

	// resp := []model.AsResp{}
	// resp = append(resp, time.Now().UnixMilli()) // 时间戳
	// resp = append(resp, config.MasterVersion)   // 版本号
	// resp = append(resp, 0)                      // 固定值
	// resp = append(resp, urlResp)                // 数据体

	// mm, err := json.Marshal(resp)
	// CheckErr(err)
	// // fmt.Println(string(mm))

	// signBody := mm[1 : len(mm)-1]
	// fmt.Println(string(signBody))

	// sessionKey := "12345678123456781234567812345678"
	// sign := encrypt.HMAC_SHA1_Encrypt(signBody, []byte(sessionKey))

	// resp = append(resp, sign)
	// mm, err = json.Marshal(resp)
	// CheckErr(err)
	// fmt.Println(string(mm))
}

func main() {
	// Gin
	gin.SetMode(gin.ReleaseMode)

	// Router
	r := gin.Default()

	// SIF
	router.SifRouter(r)

	// AS
	router.AsRouter(r)

	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
