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

func LBonusExecuteHandler(ctx *gin.Context) {
	weeks := map[string]int{
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
		"Saturday":  6,
		"Sunday":    7,
	}

	// 本月日历
	y, m, d := time.Now().Local().Date()
	cm := m

	d1 := time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
	// fmt.Println(d1)
	// fmt.Println(weeks[d1.Weekday().String()])

	d2 := d1.AddDate(0, 1, -1)
	// fmt.Println(d2)

	weeksList := []model.LbDays{}
	for c := d1; ; c = c.AddDate(0, 0, 1) {
		_, _, rd := c.Date()
		received := false
		if rd <= d {
			received = true
		}
		rw := weeks[c.Weekday().String()]
		weeksList = append(weeksList, model.LbDays{
			Day:               rd,
			DayOfTheWeek:      rw,
			SpecialDay:        false,
			SpecialImageAsset: "",
			Received:          received,
			AdReceived:        false,
			Item: model.LbDayItem{
				ItemID:  4,
				AddType: 3001,
				Amount:  1,
			},
		})
		if c == d2 {
			break
		}
	}

	// 下月日历
	y, m, _ = time.Now().AddDate(0, 1, 0).Date()
	// fmt.Println(y, m, d)

	d1 = time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
	// fmt.Println(d1)
	// fmt.Println(weeks[d1.Weekday().String()])

	d2 = d1.AddDate(0, 1, -1)
	// fmt.Println(d2)

	nextWeeksList := []model.LbDays{}
	for c := d1; ; c = c.AddDate(0, 0, 1) {
		_, _, rd := c.Date()
		rw := weeks[c.Weekday().String()]
		nextWeeksList = append(nextWeeksList, model.LbDays{
			Day:               rd,
			DayOfTheWeek:      rw,
			SpecialDay:        false,
			SpecialImageAsset: "",
			Received:          false,
			AdReceived:        false,
			Item: model.LbDayItem{
				ItemID:  4,
				AddType: 3001,
				Amount:  1,
			},
		})
		if c == d2 {
			break
		}
	}

	LbRes := model.LbResp{
		ResponseData: model.LbRes{
			Sheets: []interface{}{},
			CalendarInfo: model.CalendarInfo{
				CurrentDate: time.Now().Format("2006-01-02 03:04:05"),
				CurrentMonth: model.LbMonth{
					Year:  y,
					Month: int(cm),
					Days:  weeksList,
				},
				NextMonth: model.LbMonth{
					Year:  y,
					Month: int(m),
					Days:  nextWeeksList,
				},
			},
			TotalLoginInfo: model.TotalLoginInfo{
				LoginCount:     2626,
				RemainingCount: 74,
				Reward: []model.Reward{
					{
						ItemID:  5,
						AddType: 1000,
						Amount:  5,
					},
				},
			},
			LicenseLbonusList: []interface{}{},
			ClassSystem: model.LbClassSystem{
				RankInfo: model.LbRankInfo{
					BeforeClassRankID: 10,
					AfterClassRankID:  10,
					RankUpDate:        "2020-02-12 11:57:15",
				},
				CompleteFlag: false,
				IsOpened:     true,
				IsVisible:    true,
			},
			StartDashSheets: []interface{}{},
			EffortPoint: []model.EffortPoint{
				{
					LiveEffortPointBoxSpecID: 5,
					Capacity:                 4000000,
					Before:                   1400116,
					After:                    1400116,
					Rewards:                  []model.Rewards{},
				},
			},
			LimitedEffortBox: []interface{}{},
			MuseumInfo:       model.MuseumInfo{},
			ServerTimestamp:  time.Now().Unix(),
			PresentCnt:       0,
		},
		ReleaseInfo: []interface{}{},
		StatusCode:  200,
	}

	resp, err := json.Marshal(LbRes)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}
