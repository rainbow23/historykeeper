package main

import (
	// "bufio"
	// "database/sql"
	// "fmt"
	// _ "github.com/go-sql-driver/mysql"
	"historyKeeper/localHistory"
	"historyKeeper/sqlManager"
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

	linesHistory := localHistory.FetchLocalHistory()
	for _, oneLineHistory := range linesHistory {
		sqlManager.InsertZshHistory(oneLineHistory)
		// fmt.Println(line.Date)
		// fmt.Println(line.Command)
	}
	// fmt.Println(linesHistory)
}
