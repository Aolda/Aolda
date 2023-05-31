package p2p

import "aolda_node/blockchain"

func MessageForTxToTransaction(messageForTx *MessageForTx) *blockchain.Transaction {
	newTransaction := &blockchain.Transaction{
		Header: blockchain.TransactionHeader{
			Type:             messageForTx.Payload.Header.Type,
			Hash:             messageForTx.Payload.Header.Hash,
			BlockNumber:      messageForTx.Payload.Header.BlockNumber,
			TransactionIndex: messageForTx.Payload.Header.TransactionIndex,
			From:             messageForTx.Payload.Header.From,
			Nonce:            messageForTx.Payload.Header.Nonce,
			Signature:        messageForTx.Payload.Header.Signature,
			TimeStampe:       messageForTx.Payload.Header.TimeStampe,
		},
		Body: blockchain.TransactionBody{
			FileHash:     messageForTx.Payload.Body.FileHash,
			FunctionName: messageForTx.Payload.Body.FunctionName,
			Arguments:    messageForTx.Payload.Body.Arguments,
			Result:       messageForTx.Payload.Body.Result,
		},
	}
	return newTransaction
}
