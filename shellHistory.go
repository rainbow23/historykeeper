package main

import (
	"fmt"
	"os"
	"os/exec"
	// "strings"
)

func getstatusoutput(args ...string) (status int, output string) {
	exec_command := exec.Command(args[0], args[1:]...)
	std_out, std_err := exec_command.Output()
	status = exec_command.ProcessState.ExitCode()
	if std_err != nil {
		output = std_err.Error()
	} else {
		output = string(std_out)
	}
	return
}

func zshHistory() {
	// command := []string{"echo", "-n", "HelloWorld"}
	// command := []string{"echo", "-c", "source history.sh"}
	// command := []string{"bash", "-c", "source history.sh", " ; echo 'hi'"}
	// shellLog(command)

	// fetchZshHistory()
	// status, output := fetchZshHistory()
	shellLog()
}

func shellLog() {
	// func shellLog(command []string) {
	// status, output := getstatusoutput(command...)
	status, output := fetchZshHistory()
	shell := os.Getenv("SHELL")
	fmt.Printf("--- Result ---------------\n")
	fmt.Printf("Shell        : %s\n", shell)
	// fmt.Printf("Command      : %s\n", command)
	fmt.Printf("StatusCode   : %d\n", status)
	fmt.Printf("ResultMessage: %s\n", output)
	fmt.Printf("--------------------------\n")
}

func fetchZshHistory() (status int, output string) {
	// exec_command := shell.Command("bash", "-c", "exec $SHELL -l | source history.zsh "+" ;")
	exec_command := shell.Command("bash", "-c", 'exec -l $SHELL -c "history.zsh"')


	std_out, std_err := exec_command.Output()
	status = exec_command.ProcessState.ExitCode()

	if std_err != nil {
		output = std_err.Error()
	} else {
		output = string(std_out)
	}
	return
}
