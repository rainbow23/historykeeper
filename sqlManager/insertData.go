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
	//command string, date string) {
	// INSERT INTO shell_history2(username, command, date)  values('rainbow', 'git add ./', '2020-12-20 1');
	// query := "INSERT INTO t1(name, dt) VALUES(?,?)"

	query := fmt.Sprintf("INSERT INTO %s(%s, %s, %s) VALUES(?,?,?)",
		database, "username", "command", "date")
	// output := fmt.Sprintf("%s: %s\n", columns[i], value)

	// stmtIns, err := Db.Prepare("INSERT INTO t1(name, dt) VALUES(?,?)") // ? = placeholder

	insertTemplate(query, username, linesHistory.Command, linesHistory.Date)
}

/*
 * heroku local webで実行する必要がある
 */
func insertTemplate(query string, username string, command string, date string) {
	if command == "" || date == "" {
		return
	}

	Db := sqlConnect()
	defer Db.Close()

	//dbの最新保存日付より新しい日付のローカル履歴があれば保存する

	stmtIns, err := Db.Prepare(query) // ? = placeholder
	// stmtIns, err := Db.Prepare("INSERT INTO t1(name, dt) VALUES(?,?)") // ? = placeholder

	if err != nil {
		fmt.Println("err first")
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	_, err = stmtIns.Exec(username, command, date) // Insert tuples (i, i^2)

	if err != nil {
		fmt.Println("err second")
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
