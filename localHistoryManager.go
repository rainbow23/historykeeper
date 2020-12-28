package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	HistoryFilePath = "/Users/rainbow/.zsh_history"

	// 初期バッファサイズ
	initialBufSize = 10000
	// initialBufSize = 1
	// バッファサイズの最大値。Scannerは必要に応じこのサイズまでバッファを大きくして各行をスキャンする。
	// この値がinitialBufSize以下の場合、Scannerはバッファの拡張を一切行わず与えられた初期バッファのみを使う。
	maxBufSize = 1000000
)

// type LocalHistoryInfo struct {
//     historyList []OneLineHistory
// }

type LinesHistory []OneLineHistory

type OneLineHistory struct {
	date    string
	command string
}

func fetchLocalHistory() (linesHistory LinesHistory) {
	fp, err := os.Open(HistoryFilePath)
	defer fp.Close()

	if err != nil {
		panic(err)
	}

	// var linesHistory LinesHistory
	scanner := bufio.NewScanner(fp)

	// fmt.Println("-----------------------------------")

	buf := make([]byte, initialBufSize)
	scanner.Buffer(buf, maxBufSize)

	for scanner.Scan() {
		fmt.Println("-----------------------------------")
		timeStamp, nSec, command := separateOneLine(scanner.Text())
		// fmt.Println(timeStamp)
		// fmt.Println(nSec)
		// fmt.Println(command)

		date := convertTimeStampToDate(timeStamp, nSec)

		oneLineHistory := OneLineHistory{date, command}
		// fmt.Printf("command=%s\n", oneLineHistory.command)
		// fmt.Printf("date=%s\n", oneLineHistory.date)
		linesHistory = append(linesHistory, oneLineHistory)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %q\n", err)
	}
	return
}

func separateOneLine(oneline string) (timeStamp int64, nSec int64, command string) {
	// fmt.Println(oneline)
	patternTimeStamp := `[0-9]{10}`
	if isFoundPattern(patternTimeStamp, oneline) {
		timeStampStr := cutOffOneline(patternTimeStamp, oneline)
		timeStamp, _ = strconv.ParseInt(timeStampStr, 10, 64)
	}

	patternNsec := `:[0-9];+`
	if isFoundPattern(patternNsec, oneline) {
		nSecStr := cutOffOneline(patternNsec, oneline)
		nSecStr = strings.TrimLeft(nSecStr, ":")
		nSecStr = strings.TrimRight(nSecStr, ";")
		nSec, _ = strconv.ParseInt(nSecStr, 10, 64)
	}

	patternCommand := "[0-9];.*$"
	if isFoundPattern(patternCommand, oneline) {
		command = cutOffOneline(patternCommand, oneline)
		command = command[2:]
	}

	return
}

func isFoundPattern(regexPattern string, material string) bool {
	match, _ := regexp.MatchString(regexPattern, material)
	return match
}

func cutOffOneline(regexPattern string, oneLine string) string {
	regexp := regexp.MustCompile(regexPattern)
	res := regexp.FindAllStringSubmatch(oneLine, 1)
	// fmt.Println(res)
	if res != nil {
		return res[0][0]
	}
	return "noFound"
}

func convertTimeStampToDate(timeStamp int64, nsec int64) string {
	if timeStamp != 0 {
		return time.Unix(timeStamp, nsec).Format("2006.01.02 15:04:05")
	}
	return ""
}
