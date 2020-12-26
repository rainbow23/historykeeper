package main

import (
	"bufio"
	"fmt"
	"os"
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
	// HistoryList []struct {
	//     timestamp int
	//     nsec      int
	//     command   string
	// }

	historyList []OneLineHistory
}

type OneLineHistory struct {
	timestamp int
	nsec      int
	command   string
}

func fetchLocalHistory() {
	fp, err := os.Open(HistoryFilePath)
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

		// convertTimeStampToDate()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Scanner error: %q\n", err)
	}
}

func convertTimeStampToDate(timeStamp int64, nsec int64) string {
	timeStr := time.Unix(timeStamp, nsec).Format("2006.01.02 15:04:05")
	// timeStr := time.Unix(1608972693, 0).Format("2006.01.02 15:04:05")
	return timeStr
}
