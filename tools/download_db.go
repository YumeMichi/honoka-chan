package tools

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func GenDownloadDb() {
	// Create table
	// sql := `CREATE TABLE "main"."download_db" (
	// 	"id" integer NOT NULL,
	// 	"pkg_type" integer,
	// 	"pkg_id" integer,
	// 	"pkg_order" integer,
	// 	"pkg_size" integer,
	// 	PRIMARY KEY ("id")
	//   )`
	fileLists, err := os.ReadDir("F:/sif_dl/list_CN_Android")
	CheckErr(err)
	for _, v := range fileLists {
		if v.IsDir() {
			panic(err)
		}
		fileList := "F:/sif_dl/list_CN_Android/" + v.Name()
		fileStat, err := os.Stat(fileList)
		CheckErr(err)
		pkgSize := fileStat.Size()
		fileInfo := strings.Split(strings.ReplaceAll(v.Name(), ".zip", ""), "_")
		pkgType, pkgId, pkgOrder := fileInfo[0], fileInfo[1], fileInfo[2]
		fmt.Printf("%s - %s - %s - %d\n", pkgType, pkgId, pkgOrder, pkgSize)

		stmt, err := MainEng.DB().Prepare("INSERT INTO download_db(pkg_type,pkg_id,pkg_order,pkg_size) VALUES (?,?,?,?)")
		CheckErr(err)
		defer stmt.Close()

		res, err := stmt.Exec(pkgType, pkgId, pkgOrder, pkgSize)
		CheckErr(err)

		id, err := res.LastInsertId()
		CheckErr(err)
		fmt.Println("LastInsertId:", id)
	}
}
