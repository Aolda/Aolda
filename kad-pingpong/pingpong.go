package main

import (
	"context"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/multiformats/go-multiaddr"
)

func main() {
	ctx := context.Background()

	host, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/8001"))
	if err != nil {
		panic(err)
	}

	pingTimeout := 5 * time.Second

	dht, err := dht.New(ctx, host)
	if err != nil {
		panic(err)
	}

	peerAddress := "/ip4/127.0.0.1/tcp/8000"
	peerAddr, err := multiaddr.NewMultiaddr(peerAddress)
	if err != nil {
		panic(err)
	}

	peerInfo, err := peer.AddrInfoFromP2pAddr(peerAddr)
	if err != nil {
		panic(err)
	}

	err = host.Connect(ctx, *peerInfo)
	if err != nil {
		panic(err)
	}

	pongCh := make(chan time.Duration)

	go func() {
		start := time.Now()
		err := dht.Ping(ctx, peerInfo.ID)
		if err != nil {
			panic(err)
		}
		duration := time.Since(start)
		pongCh <- duration
	}()

	select {
	case duration := <-pongCh:
		fmt.Println("Pong received from", peerAddress, "in", duration)
	case <-time.After(pingTimeout):
		fmt.Println("Ping timeout for", peerAddress)
	}

	host.Close()
}
