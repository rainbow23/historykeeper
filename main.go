package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"os"
	"time"
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

const (
	// 初期バッファサイズ
	initialBufSize = 10000
	// initialBufSize = 1
	// バッファサイズの最大値。Scannerは必要に応じこのサイズまでバッファを大きくして各行をスキャンする。
	// この値がinitialBufSize以下の場合、Scannerはバッファの拡張を一切行わず与えられた初期バッファのみを使う。
	maxBufSize = 1000000
)

func main() {
	prepare()
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func dbProcess() {
	var Db *sql.DB
	var err error

	Db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBInfo.Username, DBInfo.Password, DBInfo.Host, DBInfo.Port, DBInfo.Database))
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer Db.Close()

	// See "Important settings" section.
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)

	rows, err := Db.Query("SELECT * FROM shell_history")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fp, err := os.Open(`/Users/rainbow/.zsh_history`)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	fmt.Println("-----------------------------------")
	buf := make([]byte, initialBufSize)
	scanner.Buffer(buf, maxBufSize)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		stmtIns, err := Db.Prepare("INSERT INTO shell_history(name) VALUES( ? )") // ? = placeholder

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		defer stmtIns.Close()                 // Close the statement when we leave main() / the program terminates
		_, err = stmtIns.Exec(scanner.Text()) // Insert tuples (i, i^2)
		if err != nil {
			// panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %q\n", err)
	}

}

func prepare() {
	DBInfo.Port = os.Getenv("PORT")
	DBInfo.Host = os.Getenv("Db.host")
	DBInfo.Database = os.Getenv("Db.database")
	DBInfo.Password = os.Getenv("Db.password")
	DBInfo.Port = os.Getenv("Db.port")
	DBInfo.Username = os.Getenv("Db.username")
}
