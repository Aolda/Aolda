package contract

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"reflect"
	"unsafe"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	compiler "aolda_node/compiler"
)

/*
*

	@dev true를 반환하면 0이 있는 거임. false면 0이 없는 거임

*
*/
func findZeroFromByte(b []byte) (bool, int) {
	for i := 0; i < len(b); i++ {
		if b[i] == 0 {
			return true, i
		}
	}
	return false, 0
}

func BytesToString(b []byte) string {
	isExist, index := findZeroFromByte(b)
	var bb []byte
	if isExist {
		bb = b[:index]
	} else {
		bb = b
	}
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&bb))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func ByteToInt(b []byte) int {
	intFromByte := new(big.Int)
	intFromByte.SetBytes(b)
	return int(intFromByte.Uint64())
}

func ListenEvent() {
	fmt.Println("::::: Listening Event :::::")
	config := LoadENV()
	fmt.Printf("Listening %s\n", config.CONTRACT_ADDRESS)
	client, err := ethclient.Dial(config.BLOCKCHAIN_WSS)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(config.CONTRACT_ADDRESS)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			// fmt.Println(vLog) // pointer to event log
			if vLog.Topics[0] == common.HexToHash(config.EVENT_SIGNATURE) {
				fmt.Println("Listen 'callAolda'")
			}
			var data []([]byte)
			for i := 0; i < len(vLog.Data); i = i + 32 {
				data = append(data, vLog.Data[i:i+32])
			}

			fileNamePointer := ByteToInt(data[0]) / 32
			functionNamePointer := ByteToInt(data[1]) / 32
			argsPointer := ByteToInt(data[2]) / 32
			argsNum := ByteToInt(data[argsPointer])

			fileName := BytesToString(data[fileNamePointer+1])
			functionName := BytesToString(data[functionNamePointer+1])

			var args []string
			for i := argsPointer + 1; i < argsPointer+1+argsNum; i++ {
				ptr := ByteToInt(data[i]) / 32
				arg := BytesToString(data[argsPointer+ptr+2])
				args = append(args, arg)
			}
			res := compiler.ExecuteJS(fileName, functionName, args)
			fmt.Println(res)
			SetValue(functionName, args, res)
		}
	}
}
