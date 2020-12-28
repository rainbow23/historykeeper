package main

import (
	// "bufio"
	"database/sql"
	"fmt"
	// "io"
	// "net/http"
	// "os"
	"time"
)

func sqlConnect() *sql.DB {
	var Db *sql.DB
	var err error

	Db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBInfo.Username, DBInfo.Password, DBInfo.Host, DBInfo.Port, DBInfo.Database))
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	// defer Db.Close()

	return Db
}

func insertZshHistory() {
	Db := sqlConnect()
	defer Db.Close()

	stmtIns, err := Db.Prepare("INSERT INTO shell_history(name) VALUES( ? )") // ? = placeholder

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()        // Close the statement when we leave main() / the program terminates
	_, err = stmtIns.Exec("aaa") // Insert tuples (i, i^2)
	if err != nil {
		// panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func showStoredZshHistory() {
	Db := sqlConnect()
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

		var oneline string
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

		// io.WriteString(w, oneline+"\n")
	}

	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
