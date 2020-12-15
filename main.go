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
	Host     string
	Username string
	Password string
	Port     string
	Database string
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World! \n")
	io.WriteString(w, DBInfo.Host+" <br>>")
	io.WriteString(w, DBInfo.Username+" <br>>")
	io.WriteString(w, DBInfo.Password+" <br>>")
	io.WriteString(w, DBInfo.Port+" <br>>")
	io.WriteString(w, DBInfo.Database+" <br>>")
}

func main() {
	port := os.Getenv("PORT")
	DBInfo.Host = os.Getenv("db_host")
	DBInfo.Database = os.Getenv("db_database")
	DBInfo.Password = os.Getenv("db_password")
	DBInfo.Port = os.Getenv("db_port")
	DBInfo.Username = os.Getenv("db_username")

	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, nil)
}
