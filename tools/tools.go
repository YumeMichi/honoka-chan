package tools

import (
	"encoding/json"
	"honoka-chan/database"
	"honoka-chan/model"
	"honoka-chan/utils"
)

func AnalysisApi1Data(path string) {
	apiData := utils.ReadAllText(path)
	if apiData != "" {
		var obj model.Response
		err := json.Unmarshal([]byte(apiData), &obj)
		if err != nil {
			panic(err)
		}

		var data interface{}
		err = json.Unmarshal(obj.ResponseData, &data)
		if err != nil {
			panic(err)
		}
		// resultType := reflect.TypeOf(data)
		// if resultType.Kind() == reflect.Map {
		// 	data = data.(map[string]interface{})
		// }
		result := data.([]interface{})
		for k, v := range result {
			m, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}

			key := ""
			switch k {
			case 0:
				key = "live_status_result"
			case 1:
				key = "live_list_result"
			case 2:
				key = "unit_list_result"
			case 3:
				key = "unit_deck_result"
			case 4:
				key = "unit_support_result"
			case 5:
				key = "owning_equip_result"
			case 6:
				key = "costume_list_result"
			case 7:
				key = "album_unit_result"
			case 8:
				key = "scenario_status_result"
			case 9:
				key = "subscenario_status_result"
			case 10:
				key = "event_scenario_result"
			case 11:
				key = "multi_unit_scenario_result"
			case 12:
				key = "product_result"
			case 13:
				key = "banner_result"
			case 14:
				key = "item_marquee_result"
			case 15:
				key = "user_intro_result"
			case 16:
				key = "special_cutin_result"
			case 17:
				key = "award_result"
			case 18:
				key = "background_result"
			case 19:
				key = "stamp_result"
			case 20:
				key = "exchange_point_result"
			case 21:
				key = "live_se_result"
			case 22:
				key = "live_icon_result"
			case 23:
				key = "item_list_result"
			case 24:
				key = "marathon_result"
			case 25:
				key = "challenge_result"
			}

			if key != "" {
				database.RedisCli.HSet(database.RedisCtx, "temp_dataset", key, string(m))
			}
		}
	}
}

func AnalysisApi2Data(path string) {
	apiData := utils.ReadAllText(path)
	if apiData != "" {
		var obj model.Response
		err := json.Unmarshal([]byte(apiData), &obj)
		if err != nil {
			panic(err)
		}

		var data interface{}
		err = json.Unmarshal(obj.ResponseData, &data)
		if err != nil {
			panic(err)
		}
		// resultType := reflect.TypeOf(data)
		// if resultType.Kind() == reflect.Map {
		// 	data = data.(map[string]interface{})
		// }
		result := data.([]interface{})
		for k, v := range result {
			m, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}

			key := ""
			switch k {
			case 0:
				key = "login_topinfo_result"
			case 1:
				key = "login_topinfo_once_result"
			case 2:
				key = "unit_accessory_result"
			case 3:
				key = "museum_result"
			}

			if key != "" {
				database.RedisCli.HSet(database.RedisCtx, "temp_dataset", key, string(m))
			}
		}
	}
}

func AnalysisApi3Data(path string) {
	apiData := utils.ReadAllText(path)
	if apiData != "" {
		var obj model.Response
		err := json.Unmarshal([]byte(apiData), &obj)
		if err != nil {
			panic(err)
		}

		var data interface{}
		err = json.Unmarshal(obj.ResponseData, &data)
		if err != nil {
			panic(err)
		}
		// resultType := reflect.TypeOf(data)
		// if resultType.Kind() == reflect.Map {
		// 	data = data.(map[string]interface{})
		// }
		result := data.([]interface{})
		for k, v := range result {
			m, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}

			key := ""
			switch k {
			case 0:
				key = "profile_livecnt_result"
			case 1:
				key = "profile_card_ranking_result"
			case 2:
				key = "profile_info_result"
			}

			if key != "" {
				database.RedisCli.HSet(database.RedisCtx, "temp_dataset", key, string(m))
			}
		}
	}
}
