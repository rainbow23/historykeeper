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
	io.WriteString(w, "HELLO WORLD!\n")
	// output := fmt.Sprintf("Host: %s\nUsername: %s\nPassword: %s\nPort:%s\nDatabase: %s",
	//     DBInfo.Host,
	//     DBInfo.Username,
	//     DBInfo.Password,
	//     DBInfo.Port,
	//     DBInfo.Database)

	// io.WriteString(w, output)
	showStoredZshHistory(w)
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
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}

func showStoredZshHistory(w http.ResponseWriter) {
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

//test

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

		var oneline string
		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for _, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			oneline += "  " + value

			// output := fmt.Sprintf("%s: %s\n", columns[i], value)
			// io.WriteString(w, output)
		}

		io.WriteString(w, oneline+"\n")
	}

	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func insertZshHistory(Db *sql.DB) {
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
	DBInfo.Host = os.Getenv("db_host")
	DBInfo.Username = os.Getenv("db_username")
	DBInfo.Password = os.Getenv("db_password")
	DBInfo.Port = os.Getenv("db_port")
	DBInfo.Database = os.Getenv("db_database")
}
