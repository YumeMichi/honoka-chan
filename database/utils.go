package database

import "fmt"

func MatchTokenUid(token, uid string) bool {
	// ret := LevelDb.List()
	// for k, v := range ret {
	// 	fmt.Println(k, v)
	// 	fmt.Println()
	// }
	res, err := LevelDb.Get([]byte(uid))
	if err != nil {
		fmt.Println(err)
		return false
	}

	return string(res) == token
}
