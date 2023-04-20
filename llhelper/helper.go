package llhelper

import (
	"encoding/json"
	"fmt"
	"honoka-chan/config"
	"honoka-chan/tools"
	"honoka-chan/utils"
	"os"
)

var (
	unitFile = "assets/unit.sd"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	LoadData()
}

func LoadData() {
	_, err := os.Stat(unitFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := Data{}
	err = json.Unmarshal([]byte(utils.ReadAllText(unitFile)), &data)
	CheckErr(err)

	session := config.UserEng.NewSession()
	defer session.Close()
	if err = session.Begin(); err != nil {
		session.Rollback()
		panic(err)
	}

	for _, team := range data.Team {
		fmt.Println(team.Cardid)
		if team.Cardid == 0 {
			continue
		}
		var unitId, unitExp, unitRarity, unitHp, unitSigned int
		exists, err := config.MainEng.Table("common_unit_m").Join("LEFT", "unit_m", "common_unit_m.unit_id = unit_m.unit_id").
			Where("unit_m.unit_number = ?", team.Cardid).
			Cols("common_unit_m.unit_id,common_unit_m.exp,unit_m.rarity,common_unit_m.max_hp,common_unit_m.is_signed").
			Get(&unitId, &unitExp, &unitRarity, &unitHp, &unitSigned)
		CheckErr(err)

		if !exists {
			panic("no such unit")
		}

		if unitRarity != 4 {
			panic("only support UR")
		}

		var diffExp, diffSmile, diffPure, diffCool int
		_, err = config.MainEng.Table("unit_level_limit_pattern_m").Where("unit_level_limit_id = 1 AND unit_level = 350").
			Cols("next_exp,smile_diff,pure_diff,cool_diff").Get(&diffExp, &diffSmile, &diffPure, &diffCool)
		CheckErr(err)

		isSigned := false
		if unitSigned == 1 {
			isSigned = true
		}

		var skillExp int
		if team.Skilllevel != 8 {
			skillExp = 0
		} else {
			skillExp = 29900
		}

		unitData := tools.UnitData{
			UserID:                      1681205441, // Replace with your own
			UnitID:                      unitId,
			Exp:                         unitExp + diffExp,
			NextExp:                     0,
			Level:                       350,
			MaxLevel:                    350,
			LevelLimitID:                1,
			Rank:                        2,
			MaxRank:                     2,
			Love:                        1000,
			MaxLove:                     1000,
			UnitSkillExp:                skillExp,
			UnitSkillLevel:              team.Skilllevel,
			MaxHp:                       unitHp,
			UnitRemovableSkillCapacity:  8,
			FavoriteFlag:                false,
			DisplayRank:                 2,
			IsRankMax:                   true,
			IsLoveMax:                   true,
			IsLevelMax:                  true,
			IsSigned:                    isSigned,
			IsSkillLevelMax:             true,
			IsRemovableSkillCapacityMax: true,
			InsertDate:                  "2023-04-20 10:00:00",
		}

		_, err = session.Table("user_unit_m").Insert(&unitData)
		if err != nil {
			session.Rollback()
			panic(err)
		}
		fmt.Println("Insert ID:", unitData.UnitOwningUserID)

		if err = session.Commit(); err != nil {
			session.Rollback()
			panic(err)
		}
	}
}
