package p2p

import (
	"aolda_node/blockchain"

	peer "github.com/libp2p/go-libp2p-core/peer"
)

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

func MessageForBlockToBlock(messageForBlock *MessageForBlock) *blockchain.Block {
	newBlock := &blockchain.Block{
		Header: blockchain.Blockheader{
			Nonce:         messageForBlock.Payload.Header.Nonce,
			PreviouseHash: messageForBlock.Payload.Header.PreviouseHash,
			BlockHash:     messageForBlock.Payload.Header.BlockHash,
			Merkleroot:    messageForBlock.Payload.Header.Merkleroot,
			Difficulty:    messageForBlock.Payload.Header.Difficulty,
			Minor:         messageForBlock.Payload.Header.Minor,
			Size:          messageForBlock.Payload.Header.Size,
			BlockNumber:   messageForBlock.Payload.Header.BlockNumber,
			TimeStamp:     messageForBlock.Payload.Header.TimeStamp,
			Signature:     messageForBlock.Payload.Header.Signature,
		},
		Body: make([]*blockchain.Transaction, len(messageForBlock.Payload.Body)),
	}

	for i, tx := range messageForBlock.Payload.Body {
		newBlock.Body[i] = &blockchain.Transaction{
			Header: blockchain.TransactionHeader{
				Type:             tx.Header.Type,
				Hash:             tx.Header.Hash,
				BlockNumber:      tx.Header.BlockNumber,
				TransactionIndex: tx.Header.TransactionIndex,
				From:             tx.Header.From,
				Nonce:            tx.Header.Nonce,
				Signature:        tx.Header.Signature,
				TimeStampe:       tx.Header.TimeStampe,
			},
			Body: blockchain.TransactionBody{
				FileHash:     tx.Body.FileHash,
				FunctionName: tx.Body.FunctionName,
				Arguments:    tx.Body.Arguments,
				Result:       tx.Body.Result,
			},
		}
	}

	return newBlock
}

func MessageForBlocksToBlocks(messageForBlocks *MessageForBlocks) []*blockchain.Block {
	blocks := make([]*blockchain.Block, len(messageForBlocks.Payload))

	for i, messageBlock := range messageForBlocks.Payload {
		newBlock := &blockchain.Block{
			Header: blockchain.Blockheader{
				Nonce:         messageBlock.Header.Nonce,
				PreviouseHash: messageBlock.Header.PreviouseHash,
				BlockHash:     messageBlock.Header.BlockHash,
				Merkleroot:    messageBlock.Header.Merkleroot,
				Difficulty:    messageBlock.Header.Difficulty,
				Minor:         messageBlock.Header.Minor,
				Size:          messageBlock.Header.Size,
				BlockNumber:   messageBlock.Header.BlockNumber,
				TimeStamp:     messageBlock.Header.TimeStamp,
				Signature:     messageBlock.Header.Signature,
			},
			Body: make([]*blockchain.Transaction, len(messageBlock.Body)),
		}

		for j, tx := range messageBlock.Body {
			newBlock.Body[j] = &blockchain.Transaction{
				Header: blockchain.TransactionHeader{
					Type:             tx.Header.Type,
					Hash:             tx.Header.Hash,
					BlockNumber:      tx.Header.BlockNumber,
					TransactionIndex: tx.Header.TransactionIndex,
					From:             tx.Header.From,
					Nonce:            tx.Header.Nonce,
					Signature:        tx.Header.Signature,
					TimeStampe:       tx.Header.TimeStampe,
				},
				Body: blockchain.TransactionBody{
					FileHash:     tx.Body.FileHash,
					FunctionName: tx.Body.FunctionName,
					Arguments:    tx.Body.Arguments,
					Result:       tx.Body.Result,
				},
			}
		}

		blocks[i] = newBlock
	}

	return blocks
}

func ConvertToPeerID(peerIDString string) peer.ID {
	peerID, err := peer.Decode(peerIDString)
	if err != nil {
		panic(err)
	}
	return peerID
}
