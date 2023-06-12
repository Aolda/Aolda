package main

import (
	api "aolda_node/api"
	contract "aolda_node/contract"
	database "aolda_node/database"
	socket "aolda_node/socket"
	p2p "aolda_node/p2p"

	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	defer database.Close()
	go func() {
		defer wg.Done()
		contract.ListenEvent()
	}()

	go func() {
		defer wg.Done()
		api.Listening()
	}()

	go func() {
		database.Boltdb()
		defer wg.Done()
	}()

	go func() {
		socket.CliStart()
		defer wg.Done()

		defer wg.Done()
		p2p.PubsubPeers()
	}()

	wg.Wait()
}
