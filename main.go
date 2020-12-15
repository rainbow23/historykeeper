package main

import (
	// "bufio"
	// "database/sql"
	// "fmt"
	// _ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"os"
	// "time"
)

var DBInfo struct {
	host     string
	Username string
	Password string
	Port     string
	Database string
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World! \n")
	io.WriteString(w, DBInfo.host)
}

func main() {
	port := os.Getenv("PORT")
	DBInfo.host = os.Getenv("db_host")
	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, nil)
}
