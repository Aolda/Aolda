package p2p

import (
	"aolda_node/blockchain"
	"aolda_node/utils"
	"context"
	"encoding/json"
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

type Message struct {
	EventName string
	Payload   string
}

func UnmarshalMessagePayload(message Message) (*Transaction, error) {
	var transaction Transaction
	err := json.Unmarshal([]byte(message.Payload), &transaction)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func Pub(eventName string, payload interface{}, ctx context.Context, topic *pubsub.Topic) {
	m := Message{
		EventName: eventName,
		Payload:   utils.ToJSON(payload),
	}
	//ctx가 고루틴한테 여러가지 정보를 전달하는 주체라고 보면 됨
	//publish내에 자체적으로 고루틴이 돌아가는 구조라 거기로 메시지 전달하면 pub됨
	message := utils.ToJSON(m)
	fmt.Print(message)
	// message를 퍼블리쉬하면됨
	if err := topic.Publish(ctx, []byte(message)); err != nil {
		fmt.Println("### Publish error:", err)
	}
	// // pub(m)
}

func SendNewestBlock(ctx context.Context, topic *pubsub.Topic) {
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)
	Pub(SEND_NEWEST_BLOCK, b, ctx, topic)
}

func RequestAllBlocks(ctx context.Context, topic *pubsub.Topic) {
	Pub(REQUEST_ALL_BLOCK, nil, ctx, topic)
}

func SendAllBlocks(ctx context.Context, topic *pubsub.Topic) {
	Pub(SEND_ALL_BLOCK, blockchain.Blocks(blockchain.Blockchain()), ctx, topic)
}

func NotifyNewBlock(b *blockchain.Block, ctx context.Context, topic *pubsub.Topic) {
	Pub(MAKE_NEW_BLOCK, b, ctx, topic)
}

func NotifyNewTx(tx *blockchain.Transaction, ctx context.Context, topic *pubsub.Topic) {
	Pub(MAKE_NEW_TX, tx, ctx, topic)
}

func NotifyNewPeer(address string, ctx context.Context, topic *pubsub.Topic) {
	Pub(MAKE_NEW_PEER, address, ctx, topic)
}

func handleMsg(m *Message) {
	switch m.EventName {
	case SEND_NEWEST_BLOCK:
		var payload blockchain.Block
		utils.HandleErr(json.Unmarshal([]byte(m.Payload), &payload))
		b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
		utils.HandleErr(err)
		if payload.Header.BlockNumber >= b.Header.BlockNumber {
			RequestAllBlocks()
		} else {
			SendNewestBlock()
		}
	case REQUEST_ALL_BLOCK:
		fmt.Printf("wants all the blocks.\n")
		SendAllBlocks()
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
