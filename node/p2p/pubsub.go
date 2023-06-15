package p2p

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"strings"
	"sync"

	blockchain "aolda_node/blockchain"
	"aolda_node/compiler"
	database "aolda_node/database"
	"aolda_node/utils"

	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	drouting "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	dutil "github.com/libp2p/go-libp2p/p2p/discovery/util"
)

const (
	txN = iota
	blockN
	blocksN
)

var (
	topicNameFlag = flag.String("topicName", "AoldaNetwork", "name of topic to join") // flag 이름, 값, 설명
	topic         *pubsub.Topic
	topicNode     *pubsub.Topic
	peerh         host.Host
	peerID        peer.ID
)

func PubsubPeers() {
	fmt.Println(":::::  Pubsub Peers  :::::")
	flag.Parse()
	ctx := context.Background()
	Nodectx := context.Background()
	h, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"))
	if err != nil {
		panic(err)
	}
	peerh = h

	go discoverPeers(ctx, h)
	addr := fmt.Sprintf("host : %s\n", h.Addrs()[0])
	parts := strings.Split(addr, "/")
	port := parts[len(parts)-1]

	database.InitDB(port)

	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		panic(err)
	}

	Nodeps, err := pubsub.NewGossipSub(Nodectx, h)
	if err != nil {
		panic(err)
	} // 개별 node에 대한 NewGossipSub

	topic, err = ps.Join(*topicNameFlag)
	if err != nil {
		panic(err)
	}

	peerID := h.ID()
	peerIDString := peerID.Pretty()

	topicNode, err = Nodeps.Join(peerIDString)
	if err != nil {
		panic(err)
	} // 개별 node에 대한 NewGossipSub, peerID를 topic으로 만듬

	sub, err := topic.Subscribe()
	if err != nil {
		panic(err)
	}
	SubMessage(ctx, sub)

	Nodesub, err := topicNode.Subscribe()
	if err != nil {
		panic(err)
	}
	SubMessage(Nodectx, Nodesub) // peerID를 topic으로 만든 것에 대해 sub listen
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
				// fmt.Println("Failed connecting to ", peer.ID.Pretty(), ", error:", err)
			} else {
				fmt.Println("Connected to:", peer.ID.Pretty())
				anyConnected = true
			}
		}
	}
	fmt.Println("Peer discovery complete")
	blockchain.Blockchain() // TODO : 블럭을 요ㅇㅏ고 해됨
	// blockchain.StopMine()
}

func SubMessage(ctx context.Context, sub *pubsub.Subscription) {
	//eventname 보고 Tx면 넣음
	for {
		s, err := sub.Next(ctx)
		if err != nil {
			panic(err)
		}
		var isConvert int
		message := string(s.Message.Data)
		var messageForTx MessageForTx
		var messageForBlock MessageForBlock
		var messageForBlocks MessageForBlocks

		errT := json.Unmarshal([]byte(message), &messageForTx)
		if errT == nil {
			fmt.Println("MessageForTx:", messageForTx)
			isConvert = txN
		}

		if isConvert != txN {
			errB := json.Unmarshal([]byte(message), &messageForBlock)
			if errB == nil {
				fmt.Println("MessageForBlock:", messageForBlock)
				isConvert = blockN
			}
		}

		if isConvert != blockN {
			errBs := json.Unmarshal([]byte(message), &messageForBlocks)
			if errBs == nil {
				fmt.Println("MessageForBlocks:", messageForBlocks)
				isConvert = blocksN
			}
		}

		if isConvert == txN {
			switch messageForTx.EventName {
			case MAKE_NEW_TX:
				mempool := blockchain.Mempool()

				// var tx *blockchain.Transaction
				// errtx := json.Unmarshal([]byte(*messageForTx.Payload), tx)
				// if errtx != nil {
				// 	log.Fatal("Fail to convert tx")
				// }

				fmt.Println("-------------------------------")
				fmt.Println("Convert message to TX")
				fmt.Println("-------------------------------")

				transaction := MessageForTxToTransaction(&messageForTx)

				//type에 따라 실행 로직이 다름
				switch transaction.Header.Type {
				case 0: // 파일 생성 (solidity 외의 언어로 작성된 스마트 컨트랙트 생성)

					// 특정 위치(./src)에 있는 fileHash(이름)를 올리고 해당 hash 값을 fileHash에 넣어주기
					// confirm 후 해당 tx를 mempool에 넣고 tx pub

				case 1: // EVM에서 AOLDA를 호출한 기록
					mempool.AddTx(transaction)
					// mempool로 직행, 이건 EVM에서 올라와서 pub하는거니깐
					fmt.Print("Blockchain: ")
					fmt.Println(blockchain.Blockchain())
					if blockchain.FindTxByBody(blockchain.Blockchain(), transaction.Body) == nil {

						res := compiler.ExecuteJS(transaction.Body.FileHash, transaction.Body.FunctionName, transaction.Body.Arguments)
						fmt.Print("res: ")
						fmt.Println(res)
						transaction.Header.Type = 3 //type1에 대한 결과값이므로 맞게 type 변경
						confirmTx, err := blockchain.MakeCofirmTx(transaction.Body.FileHash, transaction.Body.FunctionName, res, transaction.Body.Arguments)
						utils.HandleErr(err)
						fmt.Print("res: ")
						fmt.Println(confirmTx)
						// fmt.Print(*confirmTx)
						NotifyNewTx(confirmTx)
						// SetValue는 합의 후에
						// SetValue(functionName, args, res)
					}
				case 2: // USER가 API를 사용해 직접 Aolda Node를 호출한 기록
					mempool.AddTx(transaction)
				case 3: // type1과 type2에 대한 결과값
					// peer 입장에서는 실행 결과이므로, 걍 mempool에 저장
					mempool.AddTx(transaction)
				case 4: // 블록 채굴에 대한 트랜잭션
				}

			}
		} else if isConvert == blockN {
			switch messageForBlock.EventName {
			case SEND_NEWEST_BLOCK:
				//가장 최신의 block에 대해서 특정 노드에 보내는거?
			case MAKE_NEW_BLOCK:
				//local에 number랑 비교해서 추가할건지 말건지
				// 1. 추가안함 -> blockchain 길이가 더 길거나 같으므로 내것이 더 길다고 보내기(fail)

				// 2. 추가 -> mempool 초기화 후 다시 mine
				blockchain.StopMine()
				mempool := blockchain.Mempool()
				mempool.Clear()
				// case MAKE_NEW_PEER:
				// 	//peer랑 연결하는 로직
			}

		} else {
			switch messageForBlocks.EventName {
			case REQUEST_ALL_BLOCK:
				//해당 주소로 block 보내기
				peerID := ConvertToPeerID(messageForBlock.PeerID)
				SendAllBlocks(peerID)
			case SEND_ALL_BLOCK:
				//현재 local에 있는 block를 받은 block으로 바꾸기

			}
		}
	}
}
