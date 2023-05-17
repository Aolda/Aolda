package p2p

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"sync"

	"aolda_node/blockchain"

	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	drouting "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	dutil "github.com/libp2p/go-libp2p/p2p/discovery/util"
)

var (
	topicNameFlag = flag.String("topicName", "applesauce", "name of topic to join")
)

func PubsubPeers(receivedJsonData string) {

	flag.Parse()
	ctx := context.Background()

	h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		panic(err)
	}
	go discoverPeers(ctx, h)

	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		panic(err)
	}
	topic, err := ps.Join(*topicNameFlag)
	if err != nil {
		panic(err)
	}
	// go PubMessage(ctx, topic) // pub에 대한 goroutine

	sub, err := topic.Subscribe()
	if err != nil {
		panic(err)
	}
	SubMessage(ctx, sub)
}

func initDHT(ctx context.Context, h host.Host) *dht.IpfsDHT {
	// Start a DHT, for use in peer discovery. We can't just make a new DHT
	// client because we want each peer to maintain its own local copy of the
	// DHT, so that the bootstrapping node of the DHT can go down without
	// inhibiting future peer discovery.
	kademliaDHT, err := dht.New(ctx, h)
	if err != nil {
		panic(err)
	}
	if err = kademliaDHT.Bootstrap(ctx); err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	for _, peerAddr := range dht.DefaultBootstrapPeers {
		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := h.Connect(ctx, *peerinfo); err != nil {
				fmt.Println("Bootstrap warning:", err)
			}
		}()
	}
	wg.Wait()

	return kademliaDHT
}

func discoverPeers(ctx context.Context, h host.Host) {
	kademliaDHT := initDHT(ctx, h)
	routingDiscovery := drouting.NewRoutingDiscovery(kademliaDHT)
	dutil.Advertise(ctx, routingDiscovery, *topicNameFlag)

	// Look for others who have announced and attempt to connect to them
	anyConnected := false
	for !anyConnected {
		fmt.Println("Searching for peers...")
		peerChan, err := routingDiscovery.FindPeers(ctx, *topicNameFlag)
		if err != nil {
			panic(err)
		}
		for peer := range peerChan {
			if peer.ID == h.ID() {
				continue // No self connection
			}
			err := h.Connect(ctx, peer)
			if err != nil {
				fmt.Println("Failed connecting to ", peer.ID.Pretty(), ", error:", err)
			} else {
				fmt.Println("Connected to:", peer.ID.Pretty())
				anyConnected = true
			}
		}
	}
	fmt.Println("Peer discovery complete")
}

func SubMessage(ctx context.Context, sub *pubsub.Subscription) {
	//eventname 보고 Tx면 넣음
	for {
		s, err := sub.Next(ctx)
		if err != nil {
			panic(err)
		}
		message := string(s.Message.Data)
		var m Message
		errJ := json.Unmarshal([]byte(message), &m)
		if errJ != nil {
			log.Fatal(err)
		}

		switch m.EventName {
		case SEND_NEWEST_BLOCK:
			//가장 최신의 block에 대해서 보내기
		case REQUEST_ALL_BLOCK:
			//해당 주소로 block 보내기
		case SEND_ALL_BLOCK:
			//현재 local에 있는 block를 받은 block으로 바꾸기
		case MAKE_NEW_BLOCK:
			//local에 number랑 비교해서 추가할건지 말건지
			// 1. 추가안함 -> blockchain 길이가 더 길거나 같으므로 내것이 더 길다고 보내기(fail)

			// 2. 추가 -> mempool 초기화 후 다시 mine
			mempool := blockchain.Mempool()
			mempool.Clear()
		case MAKE_NEW_TX:
			var transaction Transaction
			errT := json.Unmarshal([]byte(m.Payload), &transaction)
			if errT != nil {
				log.Fatal(errT)
			}
			AddTx(transaction)
			// mempool로 직행, 이건 EVM에서 올라와서 pub하는거니깐
		case MAKE_NEW_PEER:
			//peer랑 연결하는 로직
		}
	}
}

// func printMessagesFrom(ctx context.Context, sub *pubsub.Subscription, receivedJsonData string) {
// 	for {
// 		m, err := sub.Next(ctx)
// 		if err != nil {
// 			panic(err)
// 		}
// 		message := string(m.Message.Data)
// 		fmt.Println(m.ReceivedFrom, ": ", message)
// 		parts := strings.Split(message, "/")
// 		if parts[0] == "exec" { // exec/[file name]/[function name]/[argv]
// 			argv := strings.Split(parts[3], " ")
// 			res := compiler.ExecuteJS(parts[1], parts[2], argv)
// 			fmt.Println("Exec Result: " + res)
// 		} else if parts[0] == "upload" { // upload/[filename]
// 			fmt.Println(parts[1])
// 			ipfs.IpfsAdd(parts[1])
// 		} else if parts[0] == "get" { // get/[hash value]/[file name]
// 			ipfs.Ipfsget(parts[1], parts[2])
// 		}

// 	}
// }

// func streamConsoleTo(ctx context.Context, topic *pubsub.Topic) {
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		s, err := reader.ReadString('\n')
// 		if err != nil {
// 			panic(err)
// 		}
// 		if err := topic.Publish(ctx, []byte(s)); err != nil {
// 			fmt.Println("### Publish error:", err)
// 		}
// 	}
// }
