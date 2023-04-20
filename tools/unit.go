package tools

import (
	"encoding/json"
	"fmt"
	"honoka-chan/model"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func ListUnitData() {
	sql := `SELECT unit_id,rarity,rank_min,rank_max,hp_max,default_removable_skill_capacity FROM unit_m WHERE unit_id NOT IN (SELECT unit_id FROM unit_m WHERE unit_type_id IN (SELECT unit_type_id FROM unit_type_m WHERE image_button_asset IS NULL AND (background_color = 'dcdbe3' AND original_attribute_id IS NULL) OR (unit_type_id IN (10,110,127,128,129) OR unit_type_id BETWEEN 131 AND 140))) ORDER BY unit_number ASC;`
	rows, err := MainEng.DB().Query(sql)
	CheckErr(err)
	defer rows.Close()

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
		unitData.NextExp = 0
		unitData.Rank = rank
		unitData.MaxRank = max_rank
		unitData.MaxHp = hp_max
		unitData.DisplayRank = 2
		unitData.UnitSkillLevel = 8
		unitData.FavoriteFlag = true
		unitData.IsRankMax = true
		unitData.IsLevelMax = true
		unitData.IsLoveMax = true
		unitData.IsSkillLevelMax = true
		unitData.IsRemovableSkillCapacityMax = true
		unitData.InsertDate = sdt.Local().Format("2006-01-02 03:04:05")
		if rit != 4 {
			unitData.LevelLimitID = 0
			unitData.UnitRemovableSkillCapacity = skill_capacity
			unitData.IsSigned = false
			switch rit {
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

			var count int
			err = stmt.QueryRow(uid).Scan(&count)
			CheckErr(err)

			if count > 0 {
				unitData.IsSigned = true
			} else {
				unitData.IsSigned = false
			}
		}

		unitsData = append(unitsData, unitData)
	}

	data, err := json.Marshal(unitsData)
	CheckErr(err)
	fmt.Println(string(data))
}
