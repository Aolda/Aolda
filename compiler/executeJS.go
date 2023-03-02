package compiler

import (
	"fmt"
	"os/exec"
	"strings"
)

func ExecuteJS(fileName string, funcName string, args []string) string {
	mergedArgs := strings.Join(args, " ")
	fmt.Println(mergedArgs)
	script := "./compiler/script.js"
	cmd := exec.Command("node", script, fileName, funcName, ...args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}
