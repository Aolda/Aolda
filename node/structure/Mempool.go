package structure

/**
mempool에 들어갈 트랜잭션 데이터
*/
type MemoryTransaction struct {
	header MemoryTransactionHeader
	body MemoryTransactionBody
}

/**
 mempool에 들어갈 트랜잭션 데이터의 헤더
*/
type MemoryTransactionHeader struct {
	_type       int // 0 = contract 생성, 1 = 컨
	hash        string
	from        string
	signature   Signature
	timeStampe int
}

/**
 mempool에 들어갈 트랜잭션 데이터의 헤더
*/
type MemoryTransactionBody struct {
	fileHash       int // 0 = contract 생성, 1 = 컨
	functionName        string
	arguments        string
	result   string
}


func MakeMemTxHeader(_type int, hash string, from string, signature Signature, timeStampe int ) (res MemoryTransactionHeader) {
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
