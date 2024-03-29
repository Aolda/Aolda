package blockchain

import (
	utils "aolda_node/utils"
	"aolda_node/wallet"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"
)

/*
*
블럭 헤더
*/
type Blockheader struct {
	Nonce         int              `json:"nonce`
	PreviouseHash string           `json:"previouseHash"`
	BlockHash     string           `json:"blockHash"`
	Merkleroot    string           `json:"merkleroot"`
	Difficulty    int              `json:"difficulty"`
	Minor         string           `json:"minor"`
	Size          int              `json:"size"`
	BlockNumber   int              `json:"blockNumber"`
	TimeStamp     int              `json:"timeStamp"`
	Signature     wallet.Signature `json:"signature"`
}

/*
*
블럭
*/
type Block struct {
	quit   chan bool
	Header Blockheader
	Body   []*Transaction
}

var miningBlock *Block = &Block{}
var miningBlockOnce sync.Once

/**
밈풀 가져오기 (singleton pattern)
*/

func MiningBlock() *Block {
	return miningBlock
}

/*
*
block을 생성해줌
*/
func createBlock(prevHash string, height int, diff int) *Block {
	fmt.Print("create")
	blockHeader := &Blockheader{
		Nonce:         0,
		PreviouseHash: prevHash,
		BlockHash:     "",
		Merkleroot:    "",
		Difficulty:    diff,
		Minor:         wallet.GetPublicKey(),
		Size:          0,
		BlockNumber:   height,
		TimeStamp:     0,
	}
	m := MiningBlock()
	m = &Block{
		Header: *blockHeader,
	}

	m.Body = Mempool().TxToConfirm()
	mine()
	// TODO : block.persist() => 이 타이밍에 pub
	return m
}

func persistBlock(b *Block) {
	dbStorage.SaveBlock(b.Header.BlockHash, utils.ToBytes(b))
}

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) {
	blockBytes := dbStorage.FindBlock(hash)
	if blockBytes == nil {
		return nil, errors.New("block not found")
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

/*
*
블럭 채굴
*/
// func (b *Block) mine() {
// 	target := strings.Repeat("0", b.Header.Difficulty)
// 	for {
// 		b.Header.TimeStamp = int(time.Now().Unix())
// 		hash := utils.Hash(b)
// 		fmt.Printf("Target:%s\nHash:%s\nNonce:%d\n\n", target, hash, b.Header.Nonce)
// 		if strings.HasPrefix(hash, target) {
// 			b.Header.BlockHash = hash
// 			break
// 		} else {
// 			b.Header.Nonce++
// 		}
// 	}
// }
func mine() {
	b := MiningBlock()
	target := strings.Repeat("0", b.Header.Difficulty)
	b.quit = make(chan bool)
	for {
		select {
		case <-b.quit:
			fmt.Println("Mining was stopped")
			return
		default:
			b.Header.TimeStamp = int(time.Now().Unix())
			hash := utils.Hash(b)
			fmt.Printf("Target:%s\nHash:%s\nNonce:%d\n\n", target, hash, b.Header.Nonce)
			if strings.HasPrefix(hash, target) {
				b.Header.BlockHash = hash
				return
			} else {
				b.Header.Nonce++
			}
		}
	}
}

func StopMine() {
	MiningBlock().quit <- true
	close(MiningBlock().quit)
}
