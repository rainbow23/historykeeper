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
	"os"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

func main() {
	prepare()
	http.HandleFunc("/", hello)
	http.HandleFunc("/update", updateHistory)
	http.HandleFunc("/showAll", showAllHistory)
	http.HandleFunc("/registerUser", registerUser)
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "HELLO WORLD!\n")
}

func prepare() {
	sqlManager.Prepare()
}

func showAllHistory(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "showAllHistory !\n")
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "registerUser !\n")
	sqlManager.RegisterUser("test", "test")
}

func updateHistory(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "PROCESS !\n")

	dbLatestHistory := sqlManager.FetchLatestUserInfo("rainbow")
	fmt.Println("dbLatestHistory date =  " + dbLatestHistory.Date)
	io.WriteString(w, "dbLatestHistory date =  "+dbLatestHistory.Date+"\n")

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

			io.WriteString(w, "dbLatestTime = "+dbLatestHistory.Date+"\n")
			io.WriteString(w, "localLatestTime = "+oneLineHistory.Date+"\n")
			io.WriteString(w, "command = "+oneLineHistory.Command+"\n")
		}
	}
	// fmt.Println(linesHistory)
}
