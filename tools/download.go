package tools

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func GenDownloadDb() {
	// Create table
	// CREATE TABLE "download_m" (
	// 	"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	// 	"pkg_type" integer,
	// 	"pkg_id" integer,
	// 	"pkg_order" integer,
	// 	"pkg_size" integer,
	// 	"pkg_os" TEXT
	//   );
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
		fmt.Printf("Android: %s - %s - %s - %d\n", pkgType, pkgId, pkgOrder, pkgSize)

		stmt, err := MainEng.DB().Prepare("INSERT INTO download_m(pkg_type,pkg_id,pkg_order,pkg_size,pkg_os) VALUES (?,?,?,?,?)")
		CheckErr(err)
		defer stmt.Close()

		res, err := stmt.Exec(pkgType, pkgId, pkgOrder, pkgSize, "Android")
		CheckErr(err)

		id, err := res.LastInsertId()
		CheckErr(err)
		fmt.Println("LastInsertId:", id)
	}

	fileLists, err = os.ReadDir("F:/sif_dl/list_CN_iOS")
	CheckErr(err)
	for _, v := range fileLists {
		if v.IsDir() {
			panic(err)
		}
		fileList := "F:/sif_dl/list_CN_iOS/" + v.Name()
		fileStat, err := os.Stat(fileList)
		CheckErr(err)
		pkgSize := fileStat.Size()
		fileInfo := strings.Split(strings.ReplaceAll(v.Name(), ".zip", ""), "_")
		pkgType, pkgId, pkgOrder := fileInfo[0], fileInfo[1], fileInfo[2]
		fmt.Printf("iOS: %s - %s - %s - %d\n", pkgType, pkgId, pkgOrder, pkgSize)

		stmt, err := MainEng.DB().Prepare("INSERT INTO download_m(pkg_type,pkg_id,pkg_order,pkg_size,pkg_os) VALUES (?,?,?,?,?)")
		CheckErr(err)
		defer stmt.Close()

		res, err := stmt.Exec(pkgType, pkgId, pkgOrder, pkgSize, "iOS")
		CheckErr(err)

		id, err := res.LastInsertId()
		CheckErr(err)
		fmt.Println("LastInsertId:", id)
	}
}
