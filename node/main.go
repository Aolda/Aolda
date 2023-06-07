package main

import (
	"aolda_node/api"
	contract "aolda_node/contract"
	database "aolda_node/database"
	p2p "aolda_node/p2p"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
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
		defer wg.Done()
		p2p.PubsubPeers()
	}()

	wg.Wait()
}
