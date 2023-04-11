package tools

import (
	"database/sql"
	"fmt"
	"honoka-chan/utils"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tidwall/gjson"
)

func SyncNotesList() {
	db, err := sql.Open("sqlite3", "assets/main.db")
	CheckErr(err)
	defer func() {
		db.Close()
		fmt.Println("Sync notes list done!")
	}()
	db.SetMaxOpenConns(1)

	sql := `SELECT live_setting_id,notes_setting_asset FROM live_setting_m ORDER BY live_setting_id ASC`
	rows, err := db.Query(sql)
	CheckErr(err)

	liveList := make(map[int]string)

	for rows.Next() {
		var id int
		var asset string
		err = rows.Scan(&id, &asset)
		if err != nil {
			fmt.Println(err)
			continue
		}
		liveList[id] = asset
	}
	rows.Close()

	// fmt.Println(liveList)

	for _, asset := range liveList {
		url := "https://card.niconi.co.ni/live/" + asset
		fmt.Println(url)

		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if res.StatusCode == http.StatusNotFound {
			fmt.Println(res.StatusCode)
			continue
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}
		_ = res.Body.Close()
		// fmt.Println(string(body))

		regex := regexp.MustCompile("var lives = (.*?)\n")
		matchs := regex.FindSubmatch(body)
		match := string(matchs[len(matchs)-1])

		match = strings.ReplaceAll(match, "{\\\"", "{\"")
		match = strings.ReplaceAll(match, "\\\":", "\":")
		match = strings.ReplaceAll(match, ",\\\"", ",\"")
		match = strings.ReplaceAll(match, "notes_list\":\"", "notes_list\":")
		match = strings.ReplaceAll(match, "]\"}]", "]}]")

		var notesList string
		gjson.Parse(match).ForEach(func(key, value gjson.Result) bool {
			notesList = value.Get("notes_list").String()
			return true
		})
		// fmt.Println(notesList)

		if notesList == "" {
			fmt.Println("notes_list is null")
			continue
		}

		notesDir := "./assets/notes"
		_, err = os.Stat(notesDir)
		if err != nil {
			err = os.MkdirAll(notesDir, 0755)
			CheckErr(err)
		}
		utils.WriteAllText(path.Join(notesDir, asset), notesList)

		time.Sleep(time.Second)
	}
}
