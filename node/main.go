package main

import (
	contract "aolda_node/contract"
	database "aolda_node/database"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		contract.ListenEvent()
	}()

	go func() {
		database.Boltdb()
		defer wg.Done()
	}()

	wg.Wait()
}
