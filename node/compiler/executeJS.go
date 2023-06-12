package compiler

import (
	"fmt"
	"os/exec"
)

func ExecuteJS(fileName string, funcName string, args []string) string {
	argsForCommand := []string{"./compiler/script.js", fileName, funcName}
	argsForCommand = append(argsForCommand, args...)
	cmd := exec.Command("node", argsForCommand...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}
