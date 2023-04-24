package tools

import (
	"encoding/json"
	"fmt"
	"honoka-chan/model"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type UnitData struct {
	UnitOwningUserID            int    `xorm:"unit_owning_user_id pk autoincr" json:"unit_owning_user_id"`
	UserID                      int    `xorm:"user_id" json:"-"`
	UnitID                      int    `xorm:"unit_id" json:"unit_id"`
	Exp                         int    `xorm:"exp" json:"exp"`
	NextExp                     int    `xorm:"next_exp" json:"next_exp"`
	Level                       int    `xorm:"level" json:"level"`
	MaxLevel                    int    `xorm:"max_level" json:"max_level"`
	LevelLimitID                int    `xorm:"level_limit_id" json:"level_limit_id"`
	Rank                        int    `xorm:"rank" json:"rank"`
	MaxRank                     int    `xorm:"max_rank" json:"max_rank"`
	Love                        int    `xorm:"love" json:"love"`
	MaxLove                     int    `xorm:"max_love" json:"max_love"`
	UnitSkillExp                int    `xorm:"unit_skill_exp" json:"unit_skill_exp"`
	UnitSkillLevel              int    `xorm:"unit_skill_level" json:"unit_skill_level"`
	MaxHp                       int    `xorm:"max_hp" json:"max_hp"`
	UnitRemovableSkillCapacity  int    `xorm:"unit_removable_skill_capacity" json:"unit_removable_skill_capacity"`
	FavoriteFlag                bool   `xorm:"favorite_flag" json:"favorite_flag"`
	DisplayRank                 int    `xorm:"display_rank" json:"display_rank"`
	IsRankMax                   bool   `xorm:"is_rank_max" json:"is_rank_max"`
	IsLoveMax                   bool   `xorm:"is_love_max" json:"is_love_max"`
	IsLevelMax                  bool   `xorm:"is_level_max" json:"is_level_max"`
	IsSigned                    bool   `xorm:"is_signed" json:"is_signed"`
	IsSkillLevelMax             bool   `xorm:"is_skill_level_max" json:"is_skill_level_max"`
	IsRemovableSkillCapacityMax bool   `xorm:"is_removable_skill_capacity_max" json:"is_removable_skill_capacity_max"`
	InsertDate                  string `xorm:"insert_date" json:"insert_date"`
}

type UserDeckData struct {
	ID         int    `xorm:"id pk autoincr"`
	DeckID     int    `xorm:"deck_id"`
	MainFlag   int    `xorm:"main_flag"`
	DeckName   string `xorm:"deck_name"`
	UserID     int    `xorm:"user_id"`
	InsertDate int64  `xorm:"insert_date"`
}

type UnitDeckData struct {
	ID               int   `xorm:"id pk autoincr" json:"-"`
	UserDeckID       int   `xorm:"user_deck_id" json:"-"`
	UnitOwningUserID int   `xorm:"unit_owning_user_id" json:"unit_owning_user_id"`
	UnitID           int   `xorm:"unit_id" json:"unit_id"`
	Position         int   `xorm:"position" json:"position"`
	Level            int   `xorm:"level" json:"level"`
	LevelLimitID     int   `xorm:"level_limit_id" json:"level_limit_id"`
	DisplayRank      int   `xorm:"display_rank" json:"display_rank"`
	Love             int   `xorm:"love" json:"love"`
	UnitSkillLevel   int   `xorm:"unit_skill_level" json:"unit_skill_level"`
	IsRankMax        bool  `xorm:"is_rank_max" json:"is_rank_max"`
	IsLoveMax        bool  `xorm:"is_love_max" json:"is_love_max"`
	IsLevelMax       bool  `xorm:"is_level_max" json:"is_level_max"`
	IsSigned         bool  `xorm:"is_signed" json:"is_signed"`
	BeforeLove       int   `xorm:"before_love" json:"before_love"`
	MaxLove          int   `xorm:"max_love" json:"max_love"`
	InsertData       int64 `xorm:"insert_date" json:"-"`
}

type AlbumSeries struct {
	UnitID                        int `json:"unit_id"`
	AlbumSeriesID                 int `json:"album_series_id"`
	Rarity                        int `json:"rarity"`
	RankMin                       int `json:"rank_min"`
	RankMax                       int `json:"rank_max"`
	HpMax                         int `json:"hp_max"`
	DefaultRemovableSkillCapacity int `json:"default_removable_skill_capacity"`
}

func GenCommonUnitData2() {
	sql := `SELECT unit_id,album_series_id,rarity,rank_min,rank_max,hp_max,default_removable_skill_capacity FROM unit_m WHERE unit_id NOT IN (SELECT unit_id FROM unit_m WHERE unit_type_id IN (SELECT unit_type_id FROM unit_type_m WHERE image_button_asset IS NULL AND (background_color = 'dcdbe3' AND original_attribute_id IS NULL) OR (unit_type_id IN (10,110,127,128,129) OR unit_type_id BETWEEN 131 AND 140))) ORDER BY unit_number ASC`
	rows, err := MainEng.QueryInterface(sql)
	CheckErr(err)

	rb, err := json.Marshal(rows)
	CheckErr(err)
	// fmt.Println(string(rb))

	// 默认卡组
	userDeck := UserDeckData{
		DeckID:     1,
		MainFlag:   1,
		DeckName:   "Default deck",
		UserID:     1,
		InsertDate: time.Now().Unix(),
	}
	_, err = UserEng.Table("user_deck_m").Insert(&userDeck)
	CheckErr(err)
	userDeckId := userDeck.ID
	fmt.Println("userDeckId:", userDeckId)

	position := 1

	albums := []AlbumSeries{}
	if err = json.Unmarshal(rb, &albums); err != nil {
		panic(err)
	}

	startTime := time.Now().Add(-time.Hour * 24 * 30)
	for _, v := range albums {
		unitData := UnitData{}
		startTime = startTime.Add(time.Second * 3)

		unitData.UnitID = v.UnitID
		unitData.NextExp = 0
		unitData.Rank = v.RankMin
		unitData.MaxRank = v.RankMax
		unitData.MaxHp = v.HpMax
		unitData.DisplayRank = 2
		unitData.UnitSkillLevel = 8
		unitData.FavoriteFlag = true
		unitData.IsRankMax = true
		unitData.IsLevelMax = true
		unitData.IsLoveMax = true
		unitData.IsSkillLevelMax = true
		unitData.IsRemovableSkillCapacityMax = true
		unitData.InsertDate = startTime.Local().Format("2006-01-02 03:04:05")
		if v.Rarity != 4 {
			unitData.LevelLimitID = 0
			unitData.UnitRemovableSkillCapacity = v.DefaultRemovableSkillCapacity
			unitData.IsSigned = false
			switch v.Rarity {
			case 1:
				// N
				unitData.Exp = 8000
				unitData.Level = 40
				unitData.MaxLevel = 40
				unitData.Love = 50
				unitData.MaxLove = 50
				unitData.UnitSkillExp = 0
				unitData.UnitSkillLevel = 0
				unitData.FavoriteFlag = false
				unitData.IsRankMax = false
				unitData.IsLevelMax = false
				unitData.IsLoveMax = false
				unitData.IsSkillLevelMax = false
				unitData.IsRemovableSkillCapacityMax = false
			case 2:
				// R
				unitData.Exp = 13500
				unitData.Level = 60
				unitData.MaxLevel = 60
				unitData.Love = 200
				unitData.MaxLove = 200
				unitData.UnitSkillExp = 490
			case 3:
				// SR
				unitData.Exp = 36800
				unitData.Level = 80
				unitData.MaxLevel = 80
				unitData.Love = 500
				unitData.MaxLove = 500
				unitData.UnitSkillExp = 4900
			case 5:
				// SSR
				unitData.Exp = 56657
				unitData.Level = 90
				unitData.MaxLevel = 90
				unitData.Love = 750
				unitData.MaxLove = 750
				unitData.UnitSkillExp = 12700
			}
		} else {
			// UR
			unitData.Exp = 79700
			unitData.Level = 100
			unitData.MaxLevel = 100
			unitData.LevelLimitID = 1
			unitData.Love = 1000
			unitData.MaxLove = 1000
			unitData.UnitSkillExp = 29900
			unitData.UnitRemovableSkillCapacity = 8

			stmt, err := MainEng.DB().Prepare("SELECT COUNT(*) AS ct FROM unit_sign_asset_m WHERE unit_id = ?")
			CheckErr(err)
			defer stmt.Close()

			var ct int
			err = stmt.QueryRow(v.UnitID).Scan(&ct)
			CheckErr(err)

			if ct > 0 {
				unitData.IsSigned = true
			} else {
				unitData.IsSigned = false
			}
		}

		_, err = MainEng.Table("common_unit_m").Insert(&unitData)
		CheckErr(err)
		// fmt.Println("Inserted ID:", unitData.UnitOwningUserID)

		// 默认卡组
		if v.AlbumSeriesID == 615 { // 仆光套
			unitDeckData := UnitDeckData{
				UserDeckID:       userDeckId,
				UnitOwningUserID: unitData.UnitOwningUserID,
				UnitID:           v.UnitID,
				Position:         position,
				Level:            100,
				LevelLimitID:     1,
				DisplayRank:      2,
				Love:             1000,
				UnitSkillLevel:   8,
				IsRankMax:        true,
				IsLoveMax:        true,
				IsLevelMax:       true,
				IsSigned:         unitData.IsSigned,
				BeforeLove:       1000,
				MaxLove:          1000,
				InsertData:       time.Now().Unix(),
			}

			// 入库
			_, err = UserEng.Table("deck_unit_m").Insert(&unitDeckData)
			CheckErr(err)
			deckUnitId := unitDeckData.ID
			fmt.Println("deckUnitId:", deckUnitId)

			//
			position++
		}
	}
}

func GenCommonUnitData3() {
	uId := 1681205441 // Hardcode for now
	userDeck := []UserDeckData{}
	err := UserEng.Table("user_deck_m").Where("user_id = ?", uId).Asc("deck_id").Find(&userDeck)
	CheckErr(err)

	unitDeckInfo := []model.UnitDeckInfoRes{}
	for _, deck := range userDeck {
		deckUnit := []UnitDeckData{}
		err = UserEng.Table("deck_unit_m").Where("user_deck_id = ?", deck.ID).Asc("position").Find(&deckUnit)
		CheckErr(err)

		oUids := []model.UnitOwningUserIds{}
		for _, unit := range deckUnit {
			oUids = append(oUids, model.UnitOwningUserIds{
				Position:         unit.Position,
				UnitOwningUserID: unit.UnitOwningUserID,
			})
		}

		mainFlag := false
		if deck.MainFlag == 1 {
			mainFlag = true
		}
		unitDeckInfo = append(unitDeckInfo, model.UnitDeckInfoRes{
			UnitDeckID:        deck.DeckID,
			MainFlag:          mainFlag,
			DeckName:          deck.DeckName,
			UnitOwningUserIds: oUids,
		})
	}

	b, err := json.Marshal(unitDeckInfo)
	CheckErr(err)
	fmt.Println(string(b))
}

func GenCommonUnitData() {
}
