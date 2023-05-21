package main

import (
	"crypto/rsa"
	"fmt"

	"golang.org/x/crypto/ssh"
)

const (
	FILE_MAKE = 0 + iota
	EVM_CALL
	API_CALL
	CONFIRM_VALUE
	COINBASE
)

// generatePublicKey take a rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func generatePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

	return pubKeyBytes, nil
}

type Test struct{
	Name string `json:"name"`
	Num int `json:"num"`
}

func main() {
	// // Load private key from .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// privateKeyString := os.Getenv("PRIVATE_KEY")
	// // Parse private key
	// privateKey, err := crypto.HexToECDSA(privateKeyString)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // Sign message
	// signature, err := sign.PersonalSign("Hello World.", privateKey)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(signature)
	// fmt.Println(reflect.TypeOf(signature[2:66]))
	// fmt.Println(signature[66:130])
	// fmt.Println(signature[130:132])

	// fmt.Println(" Key:", privateKey)

	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	log.Fatal("Failed to cast public key to ECDSA")
	// }

	// publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// publicKeyHex := fmt.Sprintf("0x%X", publicKeyBytes)

	// // Calculate Ethereum address from public key
	// address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	// fmt.Println("Public Key:", publicKeyHex)
	// fmt.Println("Address:", address)

	// 	test := &Test{
	// 		Name: "happy",
	// 		Num:1,
	// 	}
		
	// 	jsonString, err := json.Marshal(test)
	// 	 if err != nil {
    //     fmt.Printf("Error: %s", err)
    //     return;
    // }

	// 	fmt.Println(string(jsonString))
		fmt.Print(TEST)
	
}