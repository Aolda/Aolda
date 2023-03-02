package database

import (
	"log"
	"os"

	"github.com/boltdb/bolt"
)

func ConvertJS() {
	db, err := bolt.Open("../node/nodeDB.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("functions"))
		if b == nil {
			log.Println("Bucket not found.")
			return nil
		}
		index := 0
		b.ForEach(func(k, v []byte) error {
			//log.Printf(" %d | %s %s\n", index, k, v)
			file, err := os.Create("./src/" + string(k))
			if err != nil {
				log.Println(err)
			}
			defer file.Close()

			_, err = file.Write(v)
			if err != nil {
				log.Println(err)
			}

			index++
			log.Printf("CREATE %s IN /src\n", string(k))
			return nil
		})
		log.Printf("CREATE %d JS FILE IN /src\n", index)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
