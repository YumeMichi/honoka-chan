package tools

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"honoka-chan/model"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func ListUnitData() {
	db, err := sql.Open("sqlite3", "assets/unit.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sql := `SELECT unit_id,rarity,rank_min,rank_max,hp_max,default_removable_skill_capacity FROM unit_m WHERE unit_id NOT IN (SELECT unit_id FROM unit_m WHERE unit_type_id IN (SELECT unit_type_id FROM unit_type_m WHERE image_button_asset IS NULL AND (background_color = 'dcdbe3' AND original_attribute_id IS NULL) OR (unit_type_id IN (10,110,127,128,129) OR unit_type_id BETWEEN 131 AND 140))) ORDER BY unit_number ASC;`
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}

	unitsData := []model.Active{}
	oId := 3071290948
	sdt := time.Now().Add(-time.Hour * 24 * 30)

	for rows.Next() {
		unitData := model.Active{}
		var uid, rit, rank, max_rank, hp_max, skill_capacity int
		err = rows.Scan(&uid, &rit, &rank, &max_rank, &hp_max, &skill_capacity)
		if err != nil {
			fmt.Println(err)
			continue
		}
		oId++
		sdt = sdt.Add(time.Second * 3)

		unitData.UnitOwningUserID = oId
		unitData.UnitID = uid
		unitData.Rank = rank
		unitData.MaxRank = max_rank
		unitData.MaxHp = hp_max
		unitData.DisplayRank = 2
		unitData.InsertDate = sdt.Local().Format("2006-01-02 03:04:05")
		if rit != 4 {
			unitData.Exp = 0
			unitData.Level = 1
			unitData.LevelLimitID = 1
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
			unitData.Exp = 79700
			unitData.NextExp = 0
			unitData.Level = 100
			unitData.MaxLevel = 100
			unitData.LevelLimitID = 0
			unitData.Love = 1000
			unitData.MaxLove = 1000
			unitData.UnitSkillExp = 29900
			unitData.UnitSkillLevel = 8
			unitData.UnitRemovableSkillCapacity = 8
			unitData.FavoriteFlag = true
			unitData.IsRankMax = true
			unitData.IsLevelMax = true
			unitData.IsLoveMax = true
			unitData.IsSkillLevelMax = true
			unitData.IsRemovableSkillCapacityMax = true

			rs, err := db.Query("SELECT COUNT(*) AS ct FROM unit_sign_asset_m WHERE unit_id = ?", uid)
			if err != nil {
				panic(err)
			}

			ct := 0
			for rs.Next() {
				err = rs.Scan(&ct)
				if err != nil {
					panic(err)
				}
			}
			if ct > 0 {
				unitData.IsSigned = true
			} else {
				unitData.IsSigned = false
			}
		}

		unitsData = append(unitsData, unitData)
	}

	data, err := json.Marshal(unitsData)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
