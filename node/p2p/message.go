package p2p

import (
	"aolda_node/blockchain"
	"aolda_node/utils"
	"context"
	"fmt"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

const (
	SEND_NEWEST_BLOCK = "SEND_NEWEST_BLOCK"
	MAKE_NEW_BLOCK    = "MAKE_NEW_BLOCK"
	REQUEST_ALL_BLOCK = "REQUEST_ALL_BLOCK"
	SEND_ALL_BLOCK    = "SEND_ALL_BLOCK"
	MAKE_NEW_TX       = "MAKE_NEW_TX"
	MAKE_NEW_PEER     = "MAKE_NEW_PEER"
)

type MessageForTx struct {
	EventName string                  `json:"eventName"`
	PeerID    string                  `json:"peerid"`
	Payload   *blockchain.Transaction `json:"payload"`
}

type MessageForBlock struct {
	EventName string            `json:"eventName"`
	PeerID    string            `json:"peerid"`
	Payload   *blockchain.Block `json:"payload"`
}

type MessageForBlocks struct {
	EventName string              `json:"eventName"`
	PeerID    string              `json:"peerid"`
	Payload   []*blockchain.Block `json:"payload"`
}

var psNode *pubsub.PubSub

func PubForTx(eventName string, payload *blockchain.Transaction, ctx context.Context) {
	fmt.Println("-------------------------------")
	fmt.Println("pub the Tx")
	fmt.Println("-------------------------------")

	m := MessageForTx{
		EventName: eventName,
		Payload:   payload,
	}
	//ctx가 고루틴한테 여러가지 정보를 전달하는 주체라고 보면 됨
	//publish내에 자체적으로 고루틴이 돌아가는 구조라 거기로 메시지 전달하면 pub됨
	message := utils.ToJSON(m)
	fmt.Print(message)
	// message를 퍼블리쉬하면됨
	if err := topic.Publish(ctx, []byte(message)); err != nil {
		fmt.Println("### Publish error:", err)
	}

}

func PubForBlock(eventName string, payload *blockchain.Block, ctx context.Context) {

	m := MessageForBlock{
		EventName: eventName,
		Payload:   payload,
	}
	//ctx가 고루틴한테 여러가지 정보를 전달하는 주체라고 보면 됨
	//publish내에 자체적으로 고루틴이 돌아가는 구조라 거기로 메시지 전달하면 pub됨
	message := utils.ToJSON(m)
	fmt.Print(message)
	// message를 퍼블리쉬하면됨
	if err := topic.Publish(ctx, []byte(message)); err != nil {
		fmt.Println("### Publish error:", err)
	}

}

func PubForBlocks(eventName string, payload []*blockchain.Block, ctx context.Context) {

	m := MessageForBlocks{
		EventName: eventName,
		Payload:   payload,
	}
	//ctx가 고루틴한테 여러가지 정보를 전달하는 주체라고 보면 됨
	//publish내에 자체적으로 고루틴이 돌아가는 구조라 거기로 메시지 전달하면 pub됨
	message := utils.ToJSON(m)
	fmt.Print(message)
	// message를 퍼블리쉬하면됨
	if err := topic.Publish(ctx, []byte(message)); err != nil {
		fmt.Println("### Publish error:", err)
	}

}

func SendNewestBlock() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)
	PubForBlock(SEND_NEWEST_BLOCK, b, ctx)
}

func RequestAllBlocks() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	PubForBlock(REQUEST_ALL_BLOCK, nil, ctx)
}

func NotifyNewBlock(b *blockchain.Block) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	PubForBlock(MAKE_NEW_BLOCK, b, ctx)
}

func NotifyNewTx(tx *blockchain.Transaction) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	PubForTx(MAKE_NEW_TX, tx, ctx)
}

// func NotifyNewPeer(address string) {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	Pub(MAKE_NEW_PEER, address, ctx)
// } 1 대 1 연결

func handleMsgForBlock(m *MessageForBlock) {
	switch m.EventName {
	case SEND_NEWEST_BLOCK:
		//var payload *blockchain.Block
		// utils.HandleErr(json.Unmarshal(&m.Payload, &payload))
		b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
		utils.HandleErr(err)
		if m.Payload.Header.BlockNumber >= b.Header.BlockNumber {
			RequestAllBlocks()
		} else {
			SendNewestBlock()
		}
	case REQUEST_ALL_BLOCK:
		fmt.Printf("wants all the blocks.\n")
		//내 peer id를 넣어서 보내기
		SendAllBlocks(peerID)
	case SEND_ALL_BLOCK:
		// fmt.Printf("Received all the blocks from\n"ey)
		var payload []*blockchain.Block
		// utils.HandleErr(json.Unmarshal(&m.Payload, &payload))
		blockchain.Blockchain().Replace(payload)
	case MAKE_NEW_BLOCK:
		var payload *blockchain.Block
		// utils.HandleErr(json.Unmarshal(&m.Payload, &payload))
		blockchain.Blockchain().AddPeerBlock(payload)
	case MAKE_NEW_PEER:
		// parts := strings.Split(payload, ":")
		// peer 연결하기
	}
}

func handleMsgForTx(m *MessageForTx) {
	switch m.EventName {
	case MAKE_NEW_TX:
		var payload *blockchain.Transaction
		// utils.HandleErr(json.Unmarshal(&m.Payload, &payload))
		blockchain.Mempool().AddPeerTx(payload)
	}
}
