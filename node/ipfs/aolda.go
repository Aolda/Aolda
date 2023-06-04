package ipfs

import (
	"bytes"
	"fmt"
	"os/exec"
)

func IpfsAdd(filename string) {
	fmt.Println("IpfsAdd: filename = " + filename)
	cmd := exec.Command("/bin/sh", "./add.sh", filename)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to add file to IPFS: %s\n", err)
		fmt.Printf("Error message: %s\n", stderr.String())
	} else {
		fmt.Printf("Output: %s\n", stdout.String())
	}

}

func IpfsGet(filehash string) error {

	fmt.Print("filehash: ")
	fmt.Println(filehash)
	cmd := exec.Command("ipfs", "get", filehash)
	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("Failed to get file from IPFS: %s", err)
	}

	return nil
}

// func main() {
// 	// 서브커맨드 및 파일 이름을 처리하기 위한 플래그 정의
// 	addCommand := flag.NewFlagSet("add", flag.ExitOnError)
// 	fileName := addCommand.String("file", "", "파일 이름")

// 	getCommand := flag.NewFlagSet("get", flag.ExitOnError)
// 	fileHash := getCommand.String("filehash", "", "파일 해시")
// 	// 인자를 분석하여 실행할 서브커맨드 결정
// 	if len(os.Args) < 2 {
// 		fmt.Println("사용법: aolda [command]")
// 		fmt.Println("가능한 서브커맨드: add")
// 		os.Exit(1)
// 	}

// 	switch os.Args[1] {
// 	case "add":
// 		// add 서브커맨드를 파싱하고 실행
// 		addCommand.Parse(os.Args[2:])
// 		if *fileName == "" {
// 			fmt.Println("파일 이름을 지정해야 합니다.")
// 			os.Exit(1)
// 		}
// 		fmt.Printf("파일 추가: %s\n", *fileName)
// 		// 여기서 특정 동작을 수행
// 		IpfsAdd(*fileName)
// 		//해당 파일을 추가했다고 PUB 하기

// 		//confirmTx, err := blockchain.MakeCofirmTx(transaction.Body.FileHash, transaction.Body.FunctionName, res, transaction.Body.Arguments)
// 		//utils.HandleErr(err
// 		// NotifyNewTx(confirmTx)
// 	case "get":
// 		getCommand.Parse(os.Args[2:])
// 		if *fileHash == "" {
// 			fmt.Println("파일 해시를 지정해야 합니다.")
// 			os.Exit(1)
// 		}
// 		fmt.Printf("파일 해시 조회: %s\n", *fileHash)
// 		IpfsGet(*fileHash)
// 	default:
// 		fmt.Println("잘못된 서브커맨드입니다.")
// 		fmt.Println("가능한 서브커맨드: add")
// 		fmt.Println("가능한 서브커맨드: get")
// 		os.Exit(1)
// 	}
// }
