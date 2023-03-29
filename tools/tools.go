package tools

func init() {
	AnalysisApi1Data("assets/api1.json")
	AnalysisApi2Data("assets/api2.json")
	AnalysisApi3Data("assets/api3.json")
	// ListUnitData()
	go SyncNotesList()
}
