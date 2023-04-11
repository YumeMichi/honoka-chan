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

	// GenApi1Data()
	// GenApi2Data()
	// GenApi3Data()
	// LoadApi1Data("assets/api1.json")
	// LoadApi2Data("assets/api2.json")
	// LoadApi3Data("assets/api3.json")
	// ListUnitData()
	// go SyncNotesList()
	// GenDownloadDb()
	// GenCommonUnitData()
	InitUserData(0)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
