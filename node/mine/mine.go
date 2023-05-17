package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"
)

var (
	maxNonce = math.MaxInt64
)

var targetBits = 16

type ProofOfWork struct {
	target *big.Int
}

func NewProofOfWork(difficulty int) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))
	pow := &ProofOfWork{target}
	return pow
}

func (pow *ProofOfWork) Run(ctx context.Context, data []byte) (int, []byte) {
	var hashInt big.Int
	var hash [32]byte

	nonce := 0

	for nonce < maxNonce {
		select {
		case <-ctx.Done():
			fmt.Println("\nMining interrupted, restarting...")
			return nonce, hash[:]
		default:
			dataToHash := bytes.Join([][]byte{data, IntToHex(int64(nonce))}, []byte{})
			hash = sha256.Sum256(dataToHash)

			fmt.Printf("\r%x", hash)
			hashInt.SetBytes(hash[:])

			if hashInt.Cmp(pow.target) == -1 {
				fmt.Printf("\n\nFound nonce: %d\nHash: %x\n", nonce, hash)
				//nonce 값을 블럭에 넣고
				//SendNewestBlock를 ㄱㄱ
				return nonce, hash[:]
			} else {
				nonce++
			}
		}
	}
	fmt.Print("\n\n")
	return nonce, hash[:]
}

func IntToHex(n int64) []byte {
	return []byte(fmt.Sprintf("%x", n))
}

var (
	ctx    context.Context
	cancel context.CancelFunc
	m      sync.Mutex
	wg     sync.WaitGroup
)

func startMining(pow *ProofOfWork, data string) {
	m.Lock()
	ctx, cancel = context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		pow.Run(ctx, []byte(data))
	}()
	m.Unlock()
}

func stopMining() {
	m.Lock()
	if cancel != nil {
		cancel()
		wg.Wait()
	}
	m.Unlock()
}

func main() {
	difficulty := 20 // set your difficulty
	pow := NewProofOfWork(difficulty)

	fmt.Printf("Start mining the data \"Hello, World!\"\n")
	startMining(pow, "Hello, World!")

	// Stop mining after 30 seconds.
	time.Sleep(30 * time.Second)
	fmt.Printf("\n\nStop mining\n")
	stopMining()
}
