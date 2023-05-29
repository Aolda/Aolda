package blockchain

import (
	db "aolda_node/database"
	utils "aolda_node/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"sync"
)

const (
	defaultDifficulty  int = 2
	difficultyInterval int = 5
	blockInterval      int = 2
	allowedRange       int = 2
)

type blockchain struct {
	NewestHash        string `json:"newestHash"`
	Height            int    `json:"height"`
	CurrentDifficulty int    `json:"currentDifficulty"`
	m                 sync.Mutex
}

type storage interface {
	FindBlock(hash string) []byte
	LoadChain() []byte
	SaveBlock(hash string, data []byte)
	SaveChain(data []byte)
	DeleteAllBlocks()
}

var b *blockchain
var once sync.Once
var dbStorage storage = db.DB{}

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) AddBlock() *Block {
	block := createBlock(b.NewestHash, b.Height+1, getDifficulty(b))
	fmt.Println("AddBlock")
	fmt.Println(block)
	b.NewestHash = block.Header.BlockHash
	b.Height = block.Header.BlockNumber
	b.CurrentDifficulty = block.Header.Difficulty
	// TODO : persistBlockhain(b) pub
	return block
}

func persistBlockhain(b *blockchain) {
	dbStorage.SaveChain(utils.ToBytes(b))
}

func Blocks(b *blockchain) []*Block {
	b.m.Lock()
	defer b.m.Unlock()
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.Header.PreviouseHash != "" {
			hashCursor = block.Header.PreviouseHash
		} else {
			break
		}
	}
	return blocks
}

func Txs(b *blockchain) []*Transaction {
	var txs []*Transaction
	for _, block := range Blocks(b) {
		txs = append(txs, block.Body...)
	}
	return txs
}

func FindTx(b *blockchain, targetHash string) *Transaction {
	if b==nil{
		return nil
	}
	for _, tx := range Txs(b) {
		if tx.Header.Hash == targetHash {
			return tx
		}
	}
	return nil
}

func FindTxByBody(b *blockchain, body TransactionBody) *Transaction {
	for _, tx := range Txs(b) {
		if reflect.DeepEqual(tx.Body, body)  {
			return tx
		}
	}
	return nil
}

func recalculateDifficulty(b *blockchain) int {
	allBlocks := Blocks(b)
	newestBlock := allBlocks[0]
	lastRecalculatedBlock := allBlocks[difficultyInterval-1]
	actualTime := (newestBlock.Header.TimeStamp / 60) - (lastRecalculatedBlock.Header.TimeStamp / 60)
	expectedTime := difficultyInterval * blockInterval
	if actualTime <= (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime >= (expectedTime + allowedRange) {
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty
}

func getDifficulty(b *blockchain) int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		return recalculateDifficulty(b)
	} else {
		return b.CurrentDifficulty
	}
}

func Blockchain() *blockchain {
	once.Do(func() {
		b = &blockchain{
			Height: 0,
		}
		fmt.Println("ehre")
		checkpoint := dbStorage.LoadChain()
		fmt.Println("here")
		if checkpoint == nil {
			b.AddBlock()
		} else {
			b.restore(checkpoint)
		}
	})
	return b
}

func Status(b *blockchain, rw http.ResponseWriter) {
	b.m.Lock()
	defer b.m.Unlock()
	utils.HandleErr(json.NewEncoder(rw).Encode(b))
}

func (b *blockchain) Replace(newBlocks []*Block) {
	b.m.Lock()
	defer b.m.Unlock()
	b.CurrentDifficulty = newBlocks[0].Header.Difficulty
	b.Height = len(newBlocks)
	b.NewestHash = newBlocks[0].Header.BlockHash
	persistBlockhain(b)
	dbStorage.DeleteAllBlocks()
	for _, block := range newBlocks {
		persistBlock(block)
	}
}

func (b *blockchain) AddPeerBlock(newBlock *Block) {
	b.m.Lock()
	m.m.Lock()
	defer b.m.Unlock()
	defer m.m.Unlock()

	b.Height += 1
	b.CurrentDifficulty = newBlock.Header.Difficulty
	b.NewestHash = newBlock.Header.BlockHash

	persistBlockhain(b)
	persistBlock(newBlock)

	for _, tx := range newBlock.Body {
		_, ok := m.Txs[tx.Header.Hash]
		if ok {
			delete(m.Txs, tx.Header.Hash)
		}
	}

}
