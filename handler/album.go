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
	_ "github.com/mattn/go-sqlite3"
)

func AlbumSeriesAllHandler(ctx *gin.Context) {
	var albumIds []int
	err := MainEng.Table("album_series_m").Select("album_series_id").Find(&albumIds)
	CheckErr(err)
	// fmt.Println(albumIds)

	albumSeriesAllRes := []model.AlbumSeriesRes{}
	for _, albumId := range albumIds {
		AlbumStmt, err := MainEng.DB().Prepare("SELECT unit_id,rarity FROM unit_m WHERE album_series_id = ?")
		CheckErr(err)
		defer AlbumStmt.Close()

		rows, err := AlbumStmt.Query(albumId)
		CheckErr(err)

		albumSeriesAll := []model.AlbumResult{}
		for rows.Next() {
			var unitId, rarity int
			err = rows.Scan(&unitId, &rarity)
			CheckErr(err)

			albumSeries := model.AlbumResult{
				UnitID:           unitId,
				RankMaxFlag:      true,
				LoveMaxFlag:      true,
				RankLevelMaxFlag: true,
				AllMaxFlag:       true,
				TotalLove:        10000,
				FavoritePoint:    1000,
			}

			if rarity != 4 {
				switch rarity {
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
				signStmt, err := MainEng.DB().Prepare("SELECT COUNT(*) AS ct FROM unit_sign_asset_m WHERE unit_id = ?")
				CheckErr(err)
				defer signStmt.Close()

				var count int
				err = signStmt.QueryRow(unitId).Scan(&count)
				CheckErr(err)

				if count > 0 {
					albumSeries.SignFlag = true
				} else {
					albumSeries.SignFlag = false
				}
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
		ReleaseInfo:  []interface{}{},
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
