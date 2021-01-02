package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func FetchUUID() string {
	// コマンドを実行 .(ピリオド)は使えずにちゃんと sh コマンドを使う
	out, err := exec.Command(os.Getenv("SHELL"), "shell/uuid.zsh").Output()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// fmt.Println(string(out))
	return string(out)
}
