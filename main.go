package main

import (
	// "bufio"
	// "database/sql"
	"fmt"
	// _ "github.com/go-sql-driver/mysql"
	"historyKeeper/localHistory"
	"historyKeeper/sqlManager"
	"historyKeeper/utils"
	"io"
	"net/http"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "HELLO WORLD!\n")
}

func main() {
	prepare()
	// http.HandleFunc("/", hello)
	// port := os.Getenv("PORT")
	// http.ListenAndServe(":"+port, nil)
}

func prepare() {
	sqlManager.Prepare()
	// sqlManager.RegisterUser("test", "test")

	dbLatestHistory := sqlManager.FetchLatestUserInfo("rainbow")
	fmt.Println("dbLatestHistory date =  " + dbLatestHistory.Date)
	dbLatestTime, _ := time.Parse(TimeFormat, dbLatestHistory.Date)

	linesHistory := localHistory.FetchLocalHistory()
	for _, oneLineHistory := range linesHistory {
		//dbのuuidに紐づく最新日付より新しいシェル履歴がローカルにあれば追加する
		localLatestTime, _ := time.Parse(TimeFormat, oneLineHistory.Date)

		if localLatestTime.After(dbLatestTime) {
			fmt.Println("dbのuuidに紐づく最新日付より新しいシェル履歴がローカルにあれば追加する")
			fmt.Println("dbLatestTime = " + dbLatestHistory.Date)
			fmt.Println("localLatestTime = " + oneLineHistory.Date)
			sqlManager.InsertHistory("rainbow", oneLineHistory, utils.FetchUUID())
			fmt.Println("command = " + oneLineHistory.Command)
		}
	}
	// fmt.Println(linesHistory)
}
