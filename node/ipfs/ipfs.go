package ipfs

import (
	"fmt"
	"log"
	"os/exec"
)

func IpfsAdd(filename string) {
	fmt.Println("filename: " + filename)
	cmd := exec.Command("/bin/sh", "./ipfs/add.sh", filename)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("add.sh 실행 실패: %s\n", err)
	}

	log.Printf("결합된 출력:\n%s\n", string(output))
}

func Ipfsget(filehash, filename string) {
	cmd := exec.Command("/bin/sh", "./ipfs/get.sh", filehash, filename)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("get.sh 실행 실패: %s\n", err)
	}

	log.Printf("결합된 출력:\n%s\n", string(output))

}
