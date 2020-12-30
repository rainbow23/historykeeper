package sqlManager

import (
	// "bufio"
	"database/sql"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	// "historyKeeper/localHistory"
	// "historyKeeper/utils"
	// "io"
	// "net/http"
	// "os"
	"time"
)

func fetchUserInfo() {

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