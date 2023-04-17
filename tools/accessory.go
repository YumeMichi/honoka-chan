package tools

import "fmt"

type AccessoryData struct {
	AccessoryId int `xorm:"accessory_id"`
	SmileMax    int `xorm:"smile_max"`
	PureMax     int `xorm:"pure_max"`
	CoolMax     int `xorm:"cool_max"`
}

type CommonAccessoryData struct {
	AccessoryOwningUserId int `xorm:"accessory_owning_user_id pk autoincr"`
	AccessoryId           int `xorm:"accessory_id"`
	Exp                   int `xorm:"exp"`
}

func ProcessAccessoryData() {
	var accessoryData []AccessoryData
	err := MainEng.Table("accessory_m").Where("max_level = 8 AND rarity = 4").Cols("accessory_id,smile_max,pure_max,cool_max").Find(&accessoryData)
	CheckErr(err)

	for _, data := range accessoryData {
		var exp int
		exists, err := MainEng.Table("accessory_level_m").Where("accessory_id = ?", data.AccessoryId).Select("MAX(next_exp)").Get(&exp)
		CheckErr(err)
		if !exists {
			fmt.Println("LOL")
			continue
		}
		commonData := CommonAccessoryData{
			AccessoryId: data.AccessoryId,
			Exp:         exp,
		}
		// fmt.Println(commonData)
		// 每种饰品生成9个
		for i := 0; i < 9; i++ {
			affected, err := MainEng.Table("common_accessory_m").Insert(&commonData)
			CheckErr(err)
			fmt.Printf("Time: %d, Aff: %d\n", i, affected)

			commonData.AccessoryOwningUserId += i + 1
		}
	}
}
