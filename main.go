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
)

// type LinesHistory []OneLineHistory

// type OneLineHistory struct {
//     Date    string
//     Command string
// }

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

	history := sqlManager.FetchLatestUserInfo("rainbow")
	fmt.Println(history.Date)

	linesHistory := localHistory.FetchLocalHistory()
	for _, oneLineHistory := range linesHistory {
		//dbの日付、ローカルシェルの日付を
		// oneLineHistory.Date

		sqlManager.InsertHistory("rainbow", oneLineHistory, utils.FetchUUID())
		fmt.Println(oneLineHistory.Date)
		fmt.Println(oneLineHistory.Command)
	}
	// fmt.Println(linesHistory)
}
