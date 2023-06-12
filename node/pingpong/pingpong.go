package pingpong

import (
	"bufio"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"

	"github.com/ipfs/go-log/v2"
)

var logger = log.Logger("rendezvous")

func handleStream(stream network.Stream) {
	logger.Info("Got a new stream!")

	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	go readData(rw)
	go writeData(rw, peer.ID(stream.ID()))
}

func readData(rw *bufio.ReadWriter) {
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from buffer")
			panic(err)
		}

		if str == "" {
			return
		}
		if str != "\n" {
			fmt.Printf("read > %s", str)
		}
	}
}

func writeData(rw *bufio.ReadWriter, id peer.ID) {
	for {
		sendData := id

		_, err := rw.WriteString(fmt.Sprintf("%s\n", sendData))
		if err != nil {
			fmt.Println("Error writing to buffer")
			panic(err)
		}
		err = rw.Flush()
		if err != nil {
			fmt.Println("Error flushing buffer")
			panic(err)
		}
		fmt.Printf("write > %s\n", sendData)
		time.Sleep(3000 * time.Millisecond)
	}
}

// func main() {
// 	log.SetLogLevel("rendezvous", "info")
// 	config, err := ParseFlags()
// 	if err != nil {
// 		panic(err)
// 	}

// 	host, err := libp2p.New(libp2p.ListenAddrs([]multiaddr.Multiaddr(config.ListenAddresses)...))
// 	if err != nil {
// 		panic(err)
// 	}
// 	logger.Info("Host created. We are:", host.ID())
// 	logger.Info(host.Addrs())

// 	host.SetStreamHandler(protocol.ID(config.ProtocolID), handleStream)

// 	ctx := context.Background()
// 	kademliaDHT, err := dht.New(ctx, host)
// 	if err != nil {
// 		panic(err)
// 	}

// 	logger.Debug("Bootstrapping the DHT")
// 	if err = kademliaDHT.Bootstrap(ctx); err != nil {
// 		panic(err)
// 	}

// 	var wg sync.WaitGroup
// 	for _, peerAddr := range config.BootstrapPeers {
// 		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			if err := host.Connect(ctx, *peerinfo); err != nil {
// 				logger.Warning(err)
// 			} else {
// 				logger.Info("Connection established with bootstrap node:", *peerinfo)
// 			}
// 		}()
// 	}
// 	wg.Wait()

// 	logger.Info("Announcing ourselves...")
// 	routingDiscovery := drouting.NewRoutingDiscovery(kademliaDHT)
// 	dutil.Advertise(ctx, routingDiscovery, config.RendezvousString)
// 	logger.Debug("Successfully announced!")

// 	logger.Debug("Searching for other peers...")
// 	peerChan, err := routingDiscovery.FindPeers(ctx, config.RendezvousString)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for p := range peerChan {
// 		if p.ID == host.ID() {
// 			continue
// 		}
// 		logger.Debug("Found peer:", p)

// 		logger.Debug("Connecting to:", p)
// 		stream, err := host.NewStream(ctx, p.ID, protocol.ID(config.ProtocolID))

// 		if err != nil {
// 			logger.Warning("Connection failed:", err)
// 			continue
// 		} else {
// 			rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

// 			go writeData(rw, peer.ID(stream.ID()))
// 			go readData(rw)
// 		}

// 		logger.Info("Connected to:", p)
// 	}

// 	select {}
// }
