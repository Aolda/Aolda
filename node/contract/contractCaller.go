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
	client, err := ethclient.Dial(_nodeURL)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(_privateKey)
	if err != nil {
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
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
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
		log.Fatal(err)
	}
}

func Contract(_nodeURL, _privateKey, _contractAddress string) *contract {
	once.Do(func() {
		c.makeContract(_nodeURL, _privateKey, _contractAddress)
	})
	return &c
}

func (c *contract) makeSignature(functionName string, arguments []string) ([32]byte, error) {
	return c.instance.MakeSignature(nil, functionName, arguments)
}

func (c *contract) setValue(functionName string, arguments []string, value string) (bool, error) {
	signature, err := c.makeSignature(functionName, arguments)
	if err != nil {
		return false, err
	}

	_, err = c.instance.SetValue(c.auth, signature, value)
	if err != nil {
		return false, err
	}

	return true, nil
}

func SetValue(functionName string, arguments []string, result string) {
	config := LoadENV()
	c = *Contract(config.BLOCKCHAIN_URL, config.PRIVATE_KEY, config.CONTRACT_ADDRESS)

	signature, err := c.setValue(functionName, arguments, result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(signature) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
}
