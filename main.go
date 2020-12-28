package main

import (
	// "bufio"
	// "database/sql"
	"fmt"
	// _ "github.com/go-sql-driver/mysql"
	"historyKeeper/localHistory"
	"io"
	"net/http"
	"os"
)

var DBInfo struct {
	Host     string
	Username string
	Password string
	Port     string
	Database string
}

// var localHistory struct {
//     historyList []string
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
	DBInfo.Host = os.Getenv("db_host")
	DBInfo.Username = os.Getenv("db_username")
	DBInfo.Password = os.Getenv("db_password")
	DBInfo.Port = os.Getenv("db_port")
	DBInfo.Database = os.Getenv("db_database")

	linesHistory := localHistory.FetchLocalHistory()
	fmt.Println(linesHistory)
}
