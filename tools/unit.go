package tools

import (
	"database/sql"
	"fmt"
	"honoka-chan/model"

	"github.com/goccy/go-json"
	_ "github.com/mattn/go-sqlite3"
)

func ListUnitData() {
	db, err := sql.Open("sqlite3", "assets/unit.db")
	if err != nil {
		panic(err)
	}

	sql := `SELECT unit_id,rarity,rank_min,rank_max,hp_max,default_removable_skill_capacity FROM unit_m WHERE unit_id NOT IN (SELECT unit_id FROM unit_m WHERE unit_type_id IN (SELECT unit_type_id FROM unit_type_m WHERE image_button_asset IS NULL AND (background_color = 'dcdbe3' AND original_attribute_id IS NULL) OR (unit_type_id IN (10,110,127,128,129) OR unit_type_id BETWEEN 131 AND 140)));`
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	unitsData := []model.Active{}
	oId := 3071290948

	for rows.Next() {
		unitData := model.Active{}
		var uid, rit, rank, max_rank, hp_max, skill_capacity int
		err = rows.Scan(&uid, &rit, &rank, &max_rank, &hp_max, &skill_capacity)
		if err != nil {
			fmt.Println(err)
			continue
		}
		oId++

		unitData.UnitOwningUserID = oId
		unitData.UnitID = uid
		unitData.LevelLimitID = 1
		unitData.Rank = rank
		unitData.MaxRank = max_rank
		unitData.MaxHp = hp_max
		unitData.DisplayRank = 2
		unitData.InsertDate = "2023-03-17 21:35:20"
		if rit != 4 {
			unitData.Exp = 0
			unitData.Level = 1
			unitData.Love = 0
			unitData.UnitSkillExp = 0
			unitData.UnitSkillLevel = 1
			unitData.UnitRemovableSkillCapacity = skill_capacity
			unitData.FavoriteFlag = false
			unitData.IsRankMax = false
			unitData.IsLevelMax = false
			unitData.IsLoveMax = false
			unitData.IsSigned = false
			unitData.IsSkillLevelMax = false
			unitData.IsRemovableSkillCapacityMax = false
			switch rit {
			case 1:
				// N
				unitData.NextExp = 6
				unitData.MaxLevel = 40
				unitData.MaxLove = 50
			case 2:
				// R
				unitData.NextExp = 14
				unitData.MaxLevel = 60
				unitData.MaxLove = 200
			case 3:
				// SR
				unitData.NextExp = 54
				unitData.MaxLevel = 80
				unitData.MaxLove = 500
			case 5:
				// SSR
				unitData.NextExp = 128
				unitData.MaxLevel = 90
				unitData.MaxLove = 750
			}
		} else {
			// UR
			unitData.Exp = 117111 // Lv.120
			unitData.NextExp = 0
			unitData.Level = 120
			unitData.MaxLevel = 120
			unitData.Love = 1000
			unitData.MaxLove = 1000
			unitData.UnitSkillExp = 29900
			unitData.UnitSkillLevel = 8
			unitData.UnitRemovableSkillCapacity = 8
			unitData.FavoriteFlag = true
			unitData.IsRankMax = false
			unitData.IsLevelMax = false
			unitData.IsLoveMax = false
			unitData.IsSigned = true
			unitData.IsSkillLevelMax = true
			unitData.IsRemovableSkillCapacityMax = true
		}

		unitsData = append(unitsData, unitData)
	}

	data, err := json.Marshal(unitsData)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	db.Close()
}
