package tools

func init() {
	GenApi1Data()
	GenApi2Data()
	GenApi3Data()
	LoadApi1Data("assets/api1.json")
	LoadApi2Data("assets/api2.json")
	LoadApi3Data("assets/api3.json")
	// ListUnitData()
	// go SyncNotesList()
}
