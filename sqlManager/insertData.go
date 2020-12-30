package sqlManager

import (
	// "bufio"
	// "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"historyKeeper/localHistory"
)

const (
	database = "shell_history2"
)

func InsertHistory(username string, linesHistory localHistory.OneLineHistory) {
	query := fmt.Sprintf("INSERT INTO %s(%s, %s, %s) VALUES(?,?,?)",
		database, "username", "command", "date")

	insertTemplate(query, username, linesHistory.Command, linesHistory.Date)
}

func insertTemplate(query string, username string, command string, date string) {
	if command == "" || date == "" {
		return
	}

	Db := sqlConnect()
	defer Db.Close()

	//dbの最新保存日付より新しい日付のローカル履歴があれば保存する

	stmtIns, err := Db.Prepare(query) // ? = placeholder

	if err != nil {
		fmt.Println("err first")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	//後にcommand, date 抽象化する
	_, err = stmtIns.Exec(username, command, date) // Insert tuples (i, i^2)

	if err != nil {
		fmt.Println("err second")
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
