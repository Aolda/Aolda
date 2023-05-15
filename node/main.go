package main

import (
	//contract "aolda_node/contract"
	pubsub "aolda_node/pubsub"

	//database "aolda_node/database"
	//socket "aolda_node/socket"
	//ipfs "aolda_node/ipfs"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	contract.ListenEvent()
	// }()

	go func() {
		defer wg.Done()
		pubsub.PubsubPeers()
	}()

	// go func() {
	// 	database.Boltdb()
	// 	defer wg.Done()
	// }()

	// go func() {
	// 	socket.CliStart()
	// 	defer wg.Done()
	// }()

	wg.Wait()
}
