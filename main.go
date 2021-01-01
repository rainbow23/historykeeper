package main

import (
	// "bufio"
	// "database/sql"
	"fmt"
	// _ "github.com/go-sql-driver/mysql"
	"historyKeeper/localHistory"
	"historyKeeper/sqlManager"
	"historyKeeper/utils"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

func main() {
	prepare()
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/update", UpdateHandler)
	http.HandleFunc("/showAll", ShowAllHandler)
	http.HandleFunc("/registerUser", registerUser)
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, nil)
}

var templates = template.Must(template.ParseFiles("templates/index.html", "templates/show.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"Title": "index"}
	renderTemplate(w, "index", data)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	if err := templates.ExecuteTemplate(w, tmpl+".html", data); err != nil {
		log.Fatalln("Unable to execute template.")
	}
}

func prepare() {
	sqlManager.Prepare()
}

func ShowAllHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{"Title": "index"}
	if err := templates.ExecuteTemplate(w, "show.html", data); err != nil {
		log.Fatalln("Unable to execute template.")
	}
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "registerUser !\n")
	sqlManager.RegisterUser("test", "test")
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateHandler1")
	if r.Method != "POST" {
		http.Error(w, "Allowed POST method only", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(32 << 20) // maxMemory
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("UpdateHandler2")
	file, _, err := r.FormFile("update")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	fmt.Println("UpdateHandler3")

	path := "/tmp/.zsh_hisotry"

	f, err := os.Create(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	fmt.Println("UpdateHandler4")
	io.Copy(f, file)
	updateHistory(path)
	http.Redirect(w, r, "/showAll", http.StatusFound)
}

func updateHistory(historyPath string) {
	// func updateHistory(w http.ResponseWriter, r *http.Request) {

	dbLatestHistory := sqlManager.FetchLatestUserInfo("rainbow")
	fmt.Println("dbLatestHistory date =  " + dbLatestHistory.Date)
	// io.WriteString(w, "dbLatestHistory date =  "+dbLatestHistory.Date+"\n")

	dbLatestTime, _ := time.Parse(TimeFormat, dbLatestHistory.Date)

	linesHistory := localHistory.FetchLocalHistory(historyPath)
	for _, oneLineHistory := range linesHistory {
		//dbのuuidに紐づく最新日付より新しいシェル履歴がローカルにあれば追加する
		localLatestTime, _ := time.Parse(TimeFormat, oneLineHistory.Date)

		if localLatestTime.After(dbLatestTime) {
			fmt.Println("dbのuuidに紐づく最新日付より新しいシェル履歴がローカルにあれば追加する")
			fmt.Println("dbLatestTime = " + dbLatestHistory.Date)
			fmt.Println("localLatestTime = " + oneLineHistory.Date)
			sqlManager.InsertHistory("rainbow", oneLineHistory, utils.FetchUUID())
			fmt.Println("command = " + oneLineHistory.Command)

			/*
			 * io.WriteString(w, "dbLatestTime = "+dbLatestHistory.Date+"\n")
			 * io.WriteString(w, "localLatestTime = "+oneLineHistory.Date+"\n")
			 * io.WriteString(w, "command = "+oneLineHistory.Command+"\n")
			 */
		}
	}
	// fmt.Println(linesHistory)
}
