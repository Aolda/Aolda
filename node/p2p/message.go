package p2p

import (
	"aolda_node/blockchain"
	"aolda_node/utils"
	"encoding/json"
	"fmt"
)

const (
	SEND_NEWEST_BLOCK="SEND_NEWEST_BLOCK"
	MAKE_NEW_BLOCK = "MAKE_NEW_BLOCK"
	REQUEST_ALL_BLOCK = "REQUEST_ALL_BLOCK"
	SEND_ALL_BLOCK = "SEND_ALL_BLOCK"
	MAKE_NEW_TX = "MAKE_NEW_TX"
	MAKE_NEW_PEER="MAKE_NEW_PEER"	
)

type Message struct {
	EventName    string
	Payload string
}

func pub(eventName string, payload interface{}) {
	// m := Message{
	// 	EventName:    eventName,
	// 	Payload: utils.ToJSON(payload),
	// }
	// // pub(m)
}

func sendNewestBlock() {
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)
	pub(SEND_NEWEST_BLOCK, b)
}

func requestAllBlocks() {
	pub(REQUEST_ALL_BLOCK, nil)
}

func sendAllBlocks() {
	pub(SEND_ALL_BLOCK, blockchain.Blocks(blockchain.Blockchain()))
}

func notifyNewBlock(b *blockchain.Block) {
	pub(MAKE_NEW_BLOCK, b)
}

func notifyNewTx(tx *blockchain.Transaction) {
	pub(MAKE_NEW_TX, tx)
}

func notifyNewPeer(address string) {
	pub(MAKE_NEW_PEER, address)
}

func handleMsg(m *Message) {
	switch m.EventName {
	case SEND_NEWEST_BLOCK:
		var payload blockchain.Block
		utils.HandleErr(json.Unmarshal([]byte(m.Payload), &payload))
		b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
		utils.HandleErr(err)
		if payload.Header.BlockNumber >= b.Header.BlockNumber {
			requestAllBlocks()
		} else {
			sendNewestBlock()
		}
	case REQUEST_ALL_BLOCK:
		fmt.Printf("wants all the blocks.\n")
		sendAllBlocks()
	case SEND_ALL_BLOCK:
		// fmt.Printf("Received all the blocks from\n"ey)
		var payload []*blockchain.Block
		utils.HandleErr(json.Unmarshal([]byte(m.Payload), &payload))
		blockchain.Blockchain().Replace(payload)
	case MAKE_NEW_BLOCK:
		var payload *blockchain.Block
		utils.HandleErr(json.Unmarshal([]byte(m.Payload), &payload))
		blockchain.Blockchain().AddPeerBlock(payload)
	case MAKE_NEW_TX:
		var payload *blockchain.Transaction
		utils.HandleErr(json.Unmarshal([]byte(m.Payload), &payload))
		blockchain.Mempool().AddPeerTx(payload)
	case MAKE_NEW_PEER:
		// parts := strings.Split(payload, ":")
		// peer 연결하기
	}
}