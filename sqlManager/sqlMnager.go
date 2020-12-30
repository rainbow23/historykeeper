package sqlManager

import (
	// "bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	// "historyKeeper/localHistory"
	// "historyKeeper/utils"
	// "io"
	// "net/http"
	"os"
	"time"
)

var dBInfo struct {
	host     string
	username string
	password string
	port     string
	database string
}

func Prepare() {
	dBInfo.host = os.Getenv("db_host")
	dBInfo.username = os.Getenv("db_username")
	dBInfo.password = os.Getenv("db_password")
	dBInfo.port = os.Getenv("db_port")
	dBInfo.database = os.Getenv("db_database")

}

func sqlConnect() *sql.DB {
	var Db *sql.DB
	var err error

	Db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dBInfo.username, dBInfo.password, dBInfo.host, dBInfo.port, dBInfo.database))
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	// sql.Time_zone
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)

	return Db
}

/*
 * heroku local webで実行する必要がある
 */
// func InsertZshHistory(query string, linesHistory localHistory.OneLineHistory) {
//     Db := sqlConnect()
//     defer Db.Close()

//     fmt.Println(linesHistory.Date)
//     fmt.Println(linesHistory.Command)

//     // var latestHistory = showLatestZshHistory()

//     //dbの最新保存日付より新しい日付のローカル履歴があれば保存する

//     stmtIns, err := Db.Prepare(query) // ? = placeholder
//     // stmtIns, err := Db.Prepare("INSERT INTO t1(name, dt) VALUES(?,?)") // ? = placeholder

//     if err != nil {
//         panic(err.Error()) // proper error handling instead of panic in your app
//     }

//     defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

//     _, err = stmtIns.Exec(linesHistory.Command, linesHistory.Date) // Insert tuples (i, i^2)

//     if err != nil {
//         // panic(err.Error()) // proper error handling instead of panic in your app
//     }
// }

func showLatestZshHistory() string {
	Db := sqlConnect()
	defer Db.Close()

	// See "Important settings" section.
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)

	rows, err := Db.Query("SELECT MAX(dt),id FROM t1")
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

			fmt.Println(value)
			return value
			// output := fmt.Sprintf("%s: %s\n", columns[i], value)
			// io.WriteString(w, output)
		}
	}

	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return ""
}
