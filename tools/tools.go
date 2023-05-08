package tools

func init() {
	InitUserData(0)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
