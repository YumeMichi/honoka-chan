package tools

import (
	"fmt"
	"time"
)

type UserData struct {
	ID            int    `xorm:"id pk autoincr"`
	Phone         string `xorm:"phone"`
	Password      string `xorm:"password"`
	Autokey       string `xorm:"autokey"`
	Ticket        string `xorm:"ticket"`
	UserID        int    `xorm:"userid"`
	LastLoginTime int64  `xorm:"last_login_time"`
}

type UserPref struct {
	ID               int    `xorm:"id pk autoincr"`
	UserID           int    `xorm:"user_id"`
	AwardID          int    `xorm:"award_id"`
	BackgroundID     int    `xorm:"background_id"`
	UnitOwningUserID int    `xorm:"unit_owning_user_id"`
	UserName         string `xorm:"user_name"`
	UserLevel        int    `xorm:"user_level"`
	UserDesc         string `xorm:"user_desc"`
	UpdateTime       int64  `xorm:"update_time"`
}

func InitUserData(userId int) {
	userList := []UserData{}
	if userId != 0 {
		err := UserEng.Table("users").Where("userid = ?", userId).Find(&userList)
		CheckErr(err)
	} else {
		err := UserEng.Table("users").Asc("id").Find(&userList)
		CheckErr(err)
	}

	session := UserEng.NewSession()
	defer session.Close()

	for _, user := range userList {
		// 检查用户配置
		exists, err := UserEng.Table("user_preference_m").Where("user_id = ?", user.UserID).Exist()
		CheckErr(err)

		if !exists {
			// 默认中心成员（）
			var oId int
			_, err = MainEng.Table("common_unit_m").Cols("unit_owning_user_id").Where("unit_id = ?", 31).Get(&oId)
			CheckErr(err)
			fmt.Println("Center UnitOwningUserID:", oId)
			userPref := UserPref{
				UserID:           user.UserID,
				AwardID:          1, // 音乃木坂学生
				BackgroundID:     1, // 初始背景
				UnitOwningUserID: oId,
				UserName:         "音乃木坂学生",
				UserLevel:        1,
				UserDesc:         "你好。",
				UpdateTime:       time.Now().Unix(),
			}
			_, err = UserEng.Table("user_preference_m").Insert(&userPref)
			CheckErr(err)
			fmt.Println("UserPref ID", userPref.ID)
		} else {
			userPref := UserPref{}
			_, err = UserEng.Table("user_preference_m").Where("user_id = ?", user.UserID).Get(&userPref)
			CheckErr(err)
		}

		// 检查用户卡组配置
		exists, err = UserEng.Table("user_deck_m").Where("user_id = ?", user.UserID).Asc("deck_id").Exist()
		CheckErr(err)
		fmt.Println("UserDeck exists:", exists)

		if !exists {
			userDeck := UserDeckData{
				DeckID:     1,
				MainFlag:   1,
				DeckName:   "队伍A",
				UserID:     user.UserID,
				InsertDate: time.Now().Unix(),
			}

			if err = session.Begin(); err != nil {
				panic(err)
			}

			// 默认队伍
			_, err = session.Table("user_deck_m").Insert(&userDeck)
			if err != nil {
				session.Rollback()
				panic(err)
			}
			userDeckId := userDeck.ID
			fmt.Println("New UserDeck:", userDeckId)

			// 默认卡组
			unitIds := []int{}
			err = MainEng.Table("unit_m").Cols("unit_id").Where("album_series_id = ?", 615).Find(&unitIds)
			if err != nil {
				session.Rollback()
				panic(err)
			}

			unitData := []UnitData{}
			err = MainEng.Table("common_unit_m").In("unit_id", unitIds).Find(&unitData)
			if err != nil {
				session.Rollback()
				panic(err)
			}
			// fmt.Println(unitData)

			position := 1
			for _, unit := range unitData {
				unitDeckData := UnitDeckData{
					UserDeckID:       userDeckId,
					UnitOwningUserID: unit.UnitOwningUserID,
					UnitID:           unit.UnitID,
					Position:         position,
					Level:            100,
					LevelLimitID:     2,
					DisplayRank:      2,
					Love:             1000,
					UnitSkillLevel:   8,
					IsRankMax:        true,
					IsLoveMax:        true,
					IsLevelMax:       true,
					IsSigned:         unit.IsSigned,
					BeforeLove:       1000,
					MaxLove:          1000,
					InsertData:       time.Now().Unix(),
				}
				_, err = session.Table("deck_unit_m").Insert(&unitDeckData)
				if err != nil {
					session.Rollback()
					panic(err)
				}
				fmt.Println("New DeckUnit:", unitDeckData.ID)

				position++
			}

			if err = session.Commit(); err != nil {
				panic(err)
			}
		}
	}
}
