package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	args := []string{"12", "23"}
	a := ExecutePY("add.py", "sum_int", args)
	fmt.Println(a)
}

func ExecutePY(fileName string, funcName string, args []string) string {
	argsForCommand := []string{"script.py", fileName, funcName}
	argsForCommand = append(argsForCommand, args...)

	pythonPathCmd := exec.Command("which", "python3")
	output, err := pythonPathCmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return "err"
	}
	path := strings.TrimRight(string(output), "\n")
	cmd := exec.Command(path, argsForCommand...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}
