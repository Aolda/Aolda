package blockchain

import (
	utils "aolda_node/utils"
	"aolda_node/wallet"
	"errors"
	"sync"
	"time"
)

const (
	minerReward int = 50
)

const (
	FILE_MAKE = 0 + iota
	EVM_CALL
	API_CALL
	CONFIRM_VALUE
	COINBASE
)
/**
트랜잭션
*/
type Transaction struct {
	Header TransactionHeader `json:"header"`
	Body TransactionBody`json:"body"`
}

/**
트랜잭션 바디
*/
type TransactionBody struct {
	FileHash       string `json:"fileHash"`
	FunctionName        string `json:"functionName"`
	Arguments        []string `json:"arguments"`
	Result   string `json:"result"`// type이 4면 채굴량을 hex값으로 기록하자 
}

/**
트랜잭션 헤더
*/
type TransactionHeader struct {
	Type       int  `json:"type"`// 0 = contract 생성, 1 = 컨, 4= coinbase
	Hash        string `json:"hash"`
	BlockNumber        int `json:"blockNumber"`
	TransactionIndex   int `json:"transactionIndex"`
	From string `json:"from"`
	Nonce int `json:"nonce"`
	Signature wallet.Signature `json:"signature"`
	TimeStampe int `json:"timeStampe"`
}

type mempool struct {
	Txs map[string]*Transaction
	m   sync.Mutex
}

var m *mempool
var memOnce sync.Once

/**
 밈풀 가져오기 (singleton pattern)
*/
func Mempool() *mempool {
	memOnce.Do(func() {
		m = &mempool{
			Txs: make(map[string]*Transaction),
		}
	})
	return m
}

/**
 밈풀 초기화
*/
func (mp *mempool) Clear() {
	mp.m.Lock()
	defer mp.m.Unlock()

	mp.Txs = make(map[string]*Transaction)
}

/**
 EVM에서 Call했을 때 발생하는 Transaction 만들기
*/
func MakeEvmCallTx(fileHash, functionName string, args []string) (*Transaction, error) {
	return makeTx(EVM_CALL, fileHash,functionName,"",args)
}

/**
 API로 노드 직접 Call했을 때 발생하는 Transaction 만들기
*/
func MakeAPICallTx(fileHash, functionName string, args []string) (*Transaction, error) {
	return makeTx(API_CALL, fileHash,functionName,"",args)
}

/**
 add.js 등 file 추가했을 때 생성되는 Transaction 만들기
*/
func MakeFileTx(fileHash string) (*Transaction, error) {
	return makeTx(FILE_MAKE, fileHash,"","",nil)
}

/**
 값 확정하는 Transaction 만들기
*/
func MakeCofirmTx(fileHash, functionName, result string, args []string) (*Transaction, error) {
	return makeTx(CONFIRM_VALUE, fileHash, functionName, result, args)
}

/**
 Transaction 만들기
*/
func makeTx(_type int, fileHash, functionName, result string, args []string) (*Transaction, error) {
	txBody := TransactionBody{
		FileHash :fileHash,
		FunctionName: functionName,
		Arguments: args,
		Result:result,
	}

	txHeader := TransactionHeader{
		Type:_type,
		Hash:"",
		BlockNumber:0,
		TransactionIndex:0,
		From:wallet.GetPublicKey(),
		Nonce:0,
		Signature: wallet.Signature{},
		TimeStampe : int(time.Now().Unix()),
	}

	tx := &Transaction{
		Header: txHeader,
		Body: txBody,
	}
	return confirmTx(tx)
}

/**
 Transaction 서명
*/
func confirmTx(tx *Transaction) (*Transaction, error) {
	tx.getHash()
	tx.sign()
	valid := validate(tx)
	if !valid {
		return nil, errors.New("Tx Invalid")
	}
	return tx, nil
}

/**
	
*/
func validate(tx *Transaction) bool {
	return true
}

/**
mempool에서 트랜잭션 가져오기
*/
func (m *mempool) TxToConfirm() []*Transaction {
	coinbase := makeCoinbaseTx()
	var txs []*Transaction
	for _, tx := range m.Txs {
		txs = append(txs, tx)
	}
	txs = append(txs, coinbase)
	m.Txs = make(map[string]*Transaction)
	return txs
}

/**
채굴자한테 보상을 주는 Transaction 생성
*/
func makeCoinbaseTx() *Transaction {
	txHeader := TransactionHeader{
		Type:COINBASE,
		Hash:"",
		BlockNumber:0,
		TransactionIndex:0,
		From:wallet.GetPublicKey(),
		Nonce:0,
		Signature: wallet.Signature{},
		TimeStampe : int(time.Now().Unix()),
	}
	txBody := TransactionBody{
		FileHash :"",
		FunctionName: "",
		Arguments: nil,
		Result:"",
	}

	tx := Transaction{
		Header: txHeader,
		Body: txBody,
	}
	tx.getHash()
	return &tx
}

func (t *Transaction) getHash() {
	t.Header.Hash = utils.Hash(t)
}

func (t *Transaction) sign() {
	stringBody := utils.ToJSON(t.Body)
	
	t.Header.Signature = wallet.Sign(stringBody)
}

func (m *mempool) AddPeerTx(tx *Transaction) {
	m.m.Lock()
	defer m.m.Unlock()

	m.Txs[tx.Header.Hash] = tx

}

func (m *mempool) AddTx(tx *Transaction) {
	m.Txs[tx.Header.Hash] = tx
}