package sqlManager

import (
	// "bufio"
	// "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"historyKeeper/localHistory"
)

func InsertHistory(username string, linesHistory localHistory.OneLineHistory) {
	query := fmt.Sprintf("INSERT INTO %s(%s, %s, %s) VALUES(?,?,?)",
		dBInfo.tableHistory, "username", "command", "date")

	insertTemplate(query, username, linesHistory.Command, linesHistory.Date)
}

func RegisterUser(username string, password string) {
	if username == "" || password == "" {
		return
	}

	query := fmt.Sprintf("INSERT INTO %s(%s, %s) VALUES(?,?)",
		dBInfo.tableUserInfo, "username", "password")

	register(query, username, password)
}

func register(query string, username string, password string) {
	if username == "" || password == "" {
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
	_, err = stmtIns.Exec(username, password) // Insert tuples (i, i^2)

	if err != nil {
		fmt.Println("err second")
		panic(err.Error()) // proper error handling instead of panic in your app
	}
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
		fmt.Println("insertData " + err.Error())
		// panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	//後にcommand, date 抽象化する
	_, err = stmtIns.Exec(username, command, date) // Insert tuples (i, i^2)

	if err != nil {
		fmt.Println("insertData " + err.Error())
		// panic(err.Error()) // proper error handling instead of panic in your app
	}
}
