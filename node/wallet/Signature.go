package wallet

import (
	"aolda_node/utils"
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"

	etherSign "github.com/etaaa/Golang-Ethereum-Personal-Sign"
)

type Signature struct {
	R string
	S string
	V int
}

/**
 전자 서명 후 시그니쳐 리턴
*/
func Sign(message string) (Signature) {
	// Load private key from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyString := os.Getenv("PRIVATE_KEY")
	// Parse private key
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatal(err)
	}
	// Sign message
	signature, err := etherSign.PersonalSign(message, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	v,_ := strconv.ParseInt(signature[130:132],16,64)
	res := Signature{
		R: signature[2:66],
		S: signature[66:130],
		V: int(v),
	}
	return res
}

/**
 env에 있는 PRIVATE_KEY로 부터 address를 가져옴
*/
func GetPublicKey() string {
	privateKeyString := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	utils.HandleErr(err)

	fmt.Println(" Key:", privateKey)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to cast public key to ECDSA")
	}

	// publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// publicKeyHex := fmt.Sprintf("0x%X", publicKeyBytes)

	// Calculate Ethereum address from public key
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return address
}
