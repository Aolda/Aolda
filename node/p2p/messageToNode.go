package p2p

import (
	"aolda_node/blockchain"
	"aolda_node/utils"
	"context"
	"fmt"

	peer "github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func TestTxtoNode(tx *blockchain.Transaction, NodeId peer.ID) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	PubForTxkToNode(MAKE_NEW_PEER, tx, NodeId, ctx)
}

func SendAllBlocks(nodeID peer.ID) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	PubForBlocksToNode(SEND_ALL_BLOCK, blockchain.Blocks(blockchain.Blockchain()), nodeID, ctx)
}

func PeerIDPubSubMaker(nodeID peer.ID, ctxNode context.Context) *pubsub.Topic {
	nodeIDString := nodeID.Pretty()
	if psNode == nil {
		var err error
		psNode, err = pubsub.NewGossipSub(ctxNode, peerh)
		if err != nil {
			panic(err)
		}
	}

	TotopicNode, err := psNode.Join(nodeIDString)
	if err != nil {
		panic(err)
	}
	return TotopicNode
}

func PubForTxkToNode(eventName string, payload *blockchain.Transaction, nodeID peer.ID, ctxNode context.Context) {
	m := MessageForTx{
		EventName: eventName,
		Payload:   payload,
	}

	message := utils.ToJSON(m)
	TotopicNode := PeerIDPubSubMaker(nodeID, ctxNode)
	if psNode == nil {
		var err error
		psNode, err = pubsub.NewGossipSub(ctxNode, peerh)
		if err != nil {
			panic(err)
		}
	}

	if err := TotopicNode.Publish(context.Background(), []byte(message)); err != nil {
		fmt.Println("### Publish error:", err)
	}

	err := TotopicNode.Close() // 한번만 pub하고 끝내면 되므로, 닫아버림
	if err != nil {
		fmt.Println("### Error closing topic:", err)
	}
}

func PubForBlockToNode(eventName string, payload *blockchain.Block, nodeID peer.ID, ctxNode context.Context) {
	m := MessageForBlock{
		EventName: eventName,
		Payload:   payload,
	}

	message := utils.ToJSON(m)
	TotopicNode := PeerIDPubSubMaker(nodeID, ctxNode)
	if psNode == nil {
		var err error
		psNode, err = pubsub.NewGossipSub(ctxNode, peerh)
		if err != nil {
			panic(err)
		}
	}

	if err := TotopicNode.Publish(context.Background(), []byte(message)); err != nil {
		fmt.Println("### Publish error:", err)
	}

	err := TotopicNode.Close() // 한번만 pub하고 끝내면 되므로, 닫아버림
	if err != nil {
		fmt.Println("### Error closing topic:", err)
	}
}

func PubForBlocksToNode(eventName string, payload []*blockchain.Block, nodeID peer.ID, ctxNode context.Context) {
	m := MessageForBlocks{
		EventName: eventName,
		Payload:   payload,
	}
	//ctx가 고루틴한테 여러가지 정보를 전달하는 주체라고 보면 됨
	//publish내에 자체적으로 고루틴이 돌아가는 구조라 거기로 메시지 전달하면 pub됨
	message := utils.ToJSON(m)
	TotopicNode := PeerIDPubSubMaker(nodeID, ctxNode)
	if psNode == nil {
		var err error
		psNode, err = pubsub.NewGossipSub(ctxNode, peerh)
		if err != nil {
			panic(err)
		}
	}

	if err := TotopicNode.Publish(context.Background(), []byte(message)); err != nil {
		fmt.Println("### Publish error:", err)
	}

	err := TotopicNode.Close() // 한번만 pub하고 끝내면 되므로, 닫아버림
	if err != nil {
		fmt.Println("### Error closing topic:", err)
	}
}
