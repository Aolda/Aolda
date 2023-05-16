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

	"aolda_node/blockchain"
	compiler "aolda_node/compiler"
	"aolda_node/p2p"
	"aolda_node/utils"
)

/*
*

	@dev true를 반환하면 0이 있는 거임. false면 0이 없는 거임

*
*/
var debug = 0 // debug on/off

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
			if debug == 1 {
				fmt.Println(vLog) // pointer to event log
				fmt.Print("Event_signature: ")
				fmt.Print(config.EVENT_SIGNATURE)
				fmt.Println(common.HexToHash(config.EVENT_SIGNATURE))
				fmt.Print("vlog.topics[0]: ")
				fmt.Println(vLog.Topics[0])
			}
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

			if debug == 1 {
				fmt.Println("data: ")
				fmt.Println(data)
				fmt.Println("fileNamePointer convert")
				fmt.Printf("fileNamePointer : %d\n", fileNamePointer)
				fmt.Printf("functionNamePointer: %d\n", functionNamePointer)
				fmt.Printf("argsNum: %d\n", argsNum)
				fmt.Println("fileName convert")
				fmt.Printf("fileName: ", fileName+"\n")
				fmt.Printf("functionName: ", functionName+"\n")
				fmt.Printf("argsPointer: %d\n", argsPointer)
				fmt.Println("args convert")
			}
			var args []string
			for i := argsPointer + 1; i < argsPointer+1+argsNum; i++ {
				ptr := ByteToInt(data[i]) / 32
				arg := BytesToString(data[argsPointer+ptr+2])
				args = append(args, arg)
			}
			fmt.Println("Execute start")
			evmCallTx,err := blockchain.MakeEvmCallTx(fileName,functionName,args)
			utils.HandleErr(err)
			// fmt.Print(*evmCallTx)
			p2p.NotifyNewTx(evmCallTx)

			res := compiler.ExecuteJS(fileName, functionName, args)

			confirmTx,err := blockchain.MakeCofirmTx(fileName,functionName,res,args)
			utils.HandleErr(err)
			// fmt.Print(*confirmTx)
			p2p.NotifyNewTx(confirmTx)
			// SetValue는 합의 후에
			// SetValue(functionName, args, res)
		}
	}
}
