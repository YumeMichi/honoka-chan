package config

import (
	"encoding/json"
	"honoka-chan/utils"
	"os"
	"strconv"
	"time"
)

type AppConfigs struct {
	AppName   string    `json:"app_name"`
	Settings  Settings  `json:"settings"`
	UserPrefs UserPrefs `json:"user_prefs"`
}

type Settings struct {
	SifCdnServer string `json:"sif_cdn_server"`
	AsCdnServer  string `json:"as_cdn_server"`
}

type UserPrefs struct {
	Name           string `json:"name"`            // 用户名
	Level          int    `json:"level"`           // 用户等级
	ExpNumerator   int    `json:"exp_numerator"`   // Exp 分子
	ExpDenominator int    `json:"exp_denominator"` // Exp 分母
	GameCoin       int    `json:"game_coin"`       // 游戏金币
	SnsCoin        int    `json:"sns_coin"`        // 游戏爱心
	EnergyMax      int    `json:"energy_max"`      // 体力上限
	OverMaxEnergy  int    `json:"over_max_energy"` // 实际体力，为 0 时与 EnergyMax 一致
	InviteCode     string `json:"invite_code"`     // 用户 ID
}

func DefaultConfigs() *AppConfigs {
	return &AppConfigs{
		AppName: "honoka-chan",
		Settings: Settings{
			SifCdnServer: "http://192.168.1.123/static",
			AsCdnServer:  "http://192.168.1.123/static",
		},
		UserPrefs: UserPrefs{
			Name:           "梦路 @bilibili",
			Level:          1028,
			ExpNumerator:   1089696,
			ExpDenominator: 1207185,
			GameCoin:       112124104,
			SnsCoin:        0,
			EnergyMax:      417,
			OverMaxEnergy:  0,
			InviteCode:     "377385143",
		},
	}
}

func Load(p string) *AppConfigs {
	if !utils.PathExists(p) {
		_ = DefaultConfigs().Save(p)
	}
	c := AppConfigs{}
	err := json.Unmarshal([]byte(utils.ReadAllText(p)), &c)
	if err != nil {
		_ = os.Rename(p, p+".backup"+strconv.FormatInt(time.Now().Unix(), 10))
		_ = DefaultConfigs().Save(p)
	}
	c = AppConfigs{}
	_ = json.Unmarshal([]byte(utils.ReadAllText(p)), &c)
	return &c
}

func (c *AppConfigs) Save(p string) error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	utils.WriteAllText(p, string(data))
	return nil
}
