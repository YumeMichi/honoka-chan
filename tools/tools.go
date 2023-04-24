package tools

import (
	"honoka-chan/config"

	"xorm.io/xorm"
)

var (
	MainEng *xorm.Engine
	UserEng *xorm.Engine
)

func init() {
	MainEng = config.MainEng
	UserEng = config.UserEng

	// go SyncNotesList()
	// GenDownloadDb()
	// GenCommonUnitData()
	InitUserData(0)
	// ProcessAccessoryData()
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
