package main

import (
	contract "aolda_node/contract"
	database "aolda_node/database"
	socket "aolda_node/socket"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		contract.ListenEvent()
	}()

	go func() {
		database.Boltdb()
		defer wg.Done()
	}()

	go func() {
		socket.CliStart()
		defer wg.Done()
	}()

	wg.Wait()
}
