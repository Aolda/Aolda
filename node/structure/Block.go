package structure

/**
트랜잭션 바디
*/
type TransactionBody struct {
	fileHash       int // 0 = contract 생성, 1 = 컨
	functionName        string
	arguments        string
	result   string
}

/**
트랜잭션 헤더
*/
type TransactionHeader struct {
	_type       int // 0 = contract 생성, 1 = 컨
	hash        string
	blockNumber        int
	transactionIndex   int
	from string
	nonce int
	signature Signature
}

/**
트랜잭션
*/
type Transaction struct {
	header TransactionHeader
	body TransactionBody
}

/**
블럭 헤더
*/
type Blockheader struct {
	nonce int
	previouseHash string
	blockHash string
	merkleroot string
	difficulty int
	minor string
	size int
	number int
	timeStamp int
}


/**
블럭
*/
type Block struct {
	header Blockheader
	body []TransactionBody
}


func Make(_type int, hash string, from string, signature Signature, timeStampe int ) (res MemoryTransactionHeader) {
	res = MemoryTransactionHeader{
		_type,hash,from,signature,timeStampe,
	}
	return res
}

func MakeMemTxBody(fileHash int, functionName, arguments, result string) (res MemoryTransactionBody) {
	res = MemoryTransactionBody{
		fileHash,functionName, arguments,result,
	}
	return res
}

func MakeMemTx(header MemoryTransactionHeader, body MemoryTransactionBody) (res MemoryTransaction){
	res = MemoryTransaction{
		header,body,
	}
	return res;
}
