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

type LocalHistoryInfo struct {
	historyList []OneLineHistory
}

type OneLineHistory struct {
	timestamp int
	nsec      int
	command   string
}

func fetchLocalHistory() {
	fp, err := os.Open(HistoryFilePath)
	defer fp.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fp)

	// fmt.Println("-----------------------------------")

	buf := make([]byte, initialBufSize)
	scanner.Buffer(buf, maxBufSize)

	for scanner.Scan() {
		fmt.Println("-----------------------------------")
		timeStamp, nSec, command := separateOneLine(scanner.Text())
		fmt.Println(timeStamp)
		fmt.Println(nSec)
		fmt.Println(command)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %q\n", err)
	}
}

func separateOneLine(oneline string) (timeStamp int64, nSec int, command string) {
	fmt.Println(oneline)
	patternTimeStamp := `[0-9]{10}`
	timeStampStr := cutOffOneline(patternTimeStamp, oneline)
	timeStamp, _ = strconv.ParseInt(timeStampStr, 10, 64)

	patternNsec := `:[0-9];+`
	nSecStr := cutOffOneline(patternNsec, oneline)
	nSecStr = strings.TrimLeft(nSecStr, ":")
	nSecStr = strings.TrimRight(nSecStr, ";")
	nSec, _ = strconv.Atoi(nSecStr)

	patternCommand := "[0-9];.*$"
	command = cutOffOneline(patternCommand, oneline)
	command = command[2:]

	return
}

func cutOffOneline(regexPattern string, material string) string {
	// match, _ := regexp.MatchString(regexPattern, material)
	regexp := regexp.MustCompile(regexPattern)
	res := regexp.FindAllStringSubmatch(material, 1)
	return res[0][0]
}

func test() {
	oneline := "hoge:0045-111-2222 boke:0045-222-2222"
	fmt.Println(oneline)

	pattern := `[\d\-]+`
	match, _ := regexp.MatchString(pattern, oneline)
	fmt.Println(match)

	rep := regexp.MustCompile(pattern)

	result := rep.Split(oneline, 1)
	fmt.Println(result[0])

	r := regexp.MustCompile(pattern)
	fmt.Println(r.FindAllStringSubmatch(oneline, -1)) // => "[[0045-111-2222] [0045-222-2222]]"
}

func convertTimeStampToDate(timeStamp int64, nsec int64) string {
	timeStr := time.Unix(timeStamp, nsec).Format("2006.01.02 15:04:05")
	// timeStr := time.Unix(1608972693, 0).Format("2006.01.02 15:04:05")
	return timeStr
}
