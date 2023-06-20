package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/encrypt"
	"honoka-chan/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AlbumRes struct {
	UnitId int `xorm:"unit_id"`
	Rarity int `xorm:"rarity"`
}

func AlbumSeriesAll(ctx *gin.Context) {
	var albumIds []int
	err := MainEng.Table("album_series_m").Select("album_series_id").Find(&albumIds)
	CheckErr(err)
	// fmt.Println(albumIds)

	albumSeriesAllRes := []model.AlbumSeriesRes{}
	for _, albumId := range albumIds {
		unitList := []AlbumRes{}
		err = MainEng.Table("unit_m").Where("album_series_id = ?", albumId).Cols("unit_id,rarity").Find(&unitList)
		CheckErr(err)

		albumSeriesAll := []model.AlbumResult{}
		for _, unit := range unitList {
			albumSeries := model.AlbumResult{
				UnitID:           unit.UnitId,
				RankMaxFlag:      true,
				LoveMaxFlag:      true,
				RankLevelMaxFlag: true,
				AllMaxFlag:       true,
				TotalLove:        10000,
				FavoritePoint:    1000,
			}

			if unit.Rarity != 4 {
				switch unit.Rarity {
				case 1:
					// N
					albumSeries.HighestLovePerUnit = 50
				case 2:
					// R
					albumSeries.HighestLovePerUnit = 200
				case 3:
					// SR
					albumSeries.HighestLovePerUnit = 500
				case 5:
					// SSR
					albumSeries.HighestLovePerUnit = 750
				}
			} else {
				// UR
				albumSeries.HighestLovePerUnit = 1000

				// IsSigned
				albumSeries.SignFlag = IsSigned(unit.UnitId)
			}

			albumSeriesAll = append(albumSeriesAll, albumSeries)
		}

		albumSeriesAllRes = append(albumSeriesAllRes, model.AlbumSeriesRes{
			SeriesID: albumId,
			UnitList: albumSeriesAll,
		})
	}

	albumResp := model.AlbumSeriesResp{
		ResponseData: albumSeriesAllRes,
		ReleaseInfo:  []any{},
		StatusCode:   200,
	}
	resp, err := json.Marshal(albumResp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}
