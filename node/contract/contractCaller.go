package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	aoldaClient "aolda_node/contract/aoldaClient" // for demo
)

type contract struct {
	instance *aoldaClient.AoldaClient
	auth     *bind.TransactOpts
}

var c contract
var once sync.Once

func (co *contract) makeContract(_nodeURL, _privateKey, _contractAddress string) {
	// fmt.Println("contractCaller: MakeContract")
	client, err := ethclient.Dial(_nodeURL)
	if err != nil {
		fmt.Println("Error Fucntion: makeContract _nodeURL error: ")
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(_privateKey)
	if err != nil {
		fmt.Println("Error Fucntion: makeContract privateKey error: ")
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("Error Fucntion: makeContract fromAddress error: ")
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("Error Fucntion: makeContract gasPrice error: ")
		log.Fatal(err)
	}

	co.auth = bind.NewKeyedTransactor(privateKey)
	co.auth.Nonce = big.NewInt(int64(nonce))
	co.auth.Value = big.NewInt(0)     // in wei
	co.auth.GasLimit = uint64(300000) // in units
	co.auth.GasPrice = gasPrice

	address := common.HexToAddress(_contractAddress)
	co.instance, err = aoldaClient.NewAoldaClient(address, client)
	if err != nil {
		fmt.Println("Error Fucntion: makeContract address error: ")
		log.Fatal(err)
	}
}

func Contract(_nodeURL, _privateKey, _contractAddress string) *contract {
	once.Do(func() {
		c.makeContract(_nodeURL, _privateKey, _contractAddress)
	})
	return &c
}

func (c *contract) makeSignature(functionName string, arguments []string) (by [32]byte, err error) {
	// temp := "a"
	// nf := "aa"
	// var a []string
	// a = append(a, "a")
	// return c.instance.MakeSignature(temp, nf, a)
	return by, err
}

func (c *contract) setValue(functionName string, arguments []string, value string) (bool, error) {
	signature, err := c.makeSignature(functionName, arguments)
	if err != nil {
		fmt.Println("Error Function: c.SetValue - makeSignature Problem ")
		return false, err
	}
	_, err = c.instance.SetValue(c.auth, signature, value)
	if err != nil {
		fmt.Println("Error Function: c.SetValue - instance Problem ")
		return false, err
	}

	return true, nil
}

func SetValue(functionName string, arguments []string, result string) {
	config := LoadENV()
	c = *Contract(config.BLOCKCHAIN_URL, config.PRIVATE_KEY, config.CONTRACT_ADDRESS)

	// temp := "a"
	// nf := "aa"
	// var a []string
	// a = append(a, "a")
	// signature, _ := c.instance.MakeSignature(nil, temp, nf, a)
	// fmt.Print("here!")
	// fmt.Print(signature)
	_, err := c.setValue(functionName, arguments, result)
	fmt.Printf("res: %s",result)

	if err != nil {
		fmt.Println("Error Function: SetValue - Signature Problem ")
		log.Fatal(err)
	}
	// fmt.Println(signature) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
}
