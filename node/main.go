package main

import (
	contract "aolda_node/contract"
	database "aolda_node/database"
	p2p "aolda_node/p2p"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	defer database.Close()

	go func() {
		defer wg.Done()
		contract.ListenEvent()
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
