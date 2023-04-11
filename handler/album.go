package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/config"
	"honoka-chan/database"
	"honoka-chan/encrypt"
	"honoka-chan/model"
	"honoka-chan/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func AlbumSeriesAllHandler(ctx *gin.Context) {
	reqTime := time.Now().Unix()

	authorizeStr := ctx.Request.Header["Authorize"]
	authToken, err := utils.GetAuthorizeToken(authorizeStr)
	if err != nil {
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}
	userId := ctx.Request.Header[http.CanonicalHeaderKey("User-ID")]
	if len(userId) == 0 {
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}

	if !database.MatchTokenUid(authToken, userId[0]) {
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}

	nonce, err := utils.GetAuthorizeNonce(authorizeStr)
	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}
	nonce++

	respTime := time.Now().Unix()
	newAuthorizeStr := fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", respTime, authToken, nonce, userId[0], reqTime)
	// fmt.Println(newAuthorizeStr)

	var albumIds []int
	err = MainEng.Table("album_series_m").Select("album_series_id").Find(&albumIds)
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

	resp := model.AlbumSeriesResp{
		ResponseData: albumSeriesAllRes,
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}
	respb, err := json.Marshal(resp)
	CheckErr(err)
	// fmt.Println(string(respb))

	xms := encrypt.RSA_Sign_SHA1(respb, "privatekey.pem")
	xms64 := base64.RawStdEncoding.EncodeToString(xms)

	ctx.Header("Server-Version", config.Conf.Server.VersionNumber)
	ctx.Header("user_id", userId[0])
	ctx.Header("authorize", newAuthorizeStr)
	ctx.Header("X-Message-Sign", xms64)

	ctx.String(http.StatusOK, string(respb))
}
