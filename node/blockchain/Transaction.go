package blockchain

import (
	utils "aolda_node/utils"
	"aolda_node/wallet"
	"errors"
	"sync"
)

const (
	minerReward int = 50
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
 Transaction 만들기
*/
func makeTx() (*Transaction, error) {
	txHeader := TransactionHeader{
		Type:4,
		Hash:"",
		BlockNumber:0,
		TransactionIndex:0,
		From:wallet.GetPublicKey(),
		Nonce:0,
		Signature: wallet.Signature{},
		TimeStampe :0,
	}
	txBody := TransactionBody{
		FileHash :"",
		FunctionName: "",
		Arguments: nil,
		Result:"",
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
	tx.getId()
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
		Type:4,
		Hash:"",
		BlockNumber:0,
		TransactionIndex:0,
		From:wallet.GetPublicKey(),
		Nonce:0,
		Signature: wallet.Signature{},
		TimeStampe :0,
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
	tx.getId()
	return &tx
}

func (t *Transaction) getId() {
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