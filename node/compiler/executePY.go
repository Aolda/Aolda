package compiler

import (
	"fmt"
	"os/exec"
)

func ExecutePY(fileName string, funcName string, args []string) string {
	argsForCommand := []string{"script.py", fileName, funcName}
	argsForCommand = append(argsForCommand, args...)
	pythonPathCmd := exec.Command("which", "python3")
	pythonPath, err := pythonPathCmd.Output()
	if err != nil {
		panic(err)
	}
	cmd := exec.Command(string(pythonPath), argsForCommand...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}
