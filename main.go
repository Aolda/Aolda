package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/boltdb/bolt"
)

func ConvertJS(db *bolt.DB, err error, path string) {
	fmt.Println("::::: CREATE Database :::::")

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("functions"))
		if b == nil {
			log.Println("Bucket not found.")
			return nil
		}
		index := 0
		b.ForEach(func(k, v []byte) error {
			//log.Printf(" %d | %s %s\n", index, k, v)
			file, err := os.Create(path + "/" + string(k))
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
		log.Printf("CREATE %d FILE IN /src\n", index)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func InitDB(db *bolt.DB, err error, dir string) {
	fmt.Println("::::: INIT Database :::::")

	err = db.Update(func(tx *bolt.Tx) error {
		// Create a bucket to store data in.
		b, err := tx.CreateBucketIfNotExists([]byte("functions"))
		if err != nil {
			return err
		}

		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}

		for _, fileInDir := range files {
			file, err := os.Open(dir + "/" + fileInDir.Name())
			if err != nil {
				return err
			}
			defer file.Close()

			data, err := ioutil.ReadAll(file)
			if err != nil {
				return err
			}

			err = b.Put([]byte(fileInDir.Name()), data)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func ViewAllDB(db *bolt.DB, err error) {

	fmt.Println("::::: VIEW Database :::::")

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("functions"))
		if b == nil {
			log.Println("Bucket not found.")
			return nil
		}
		index := 0
		log.Println("__________________________")
		b.ForEach(func(k, v []byte) error {
			log.Printf(" %d | %s\n", index, k)
			index++
			return nil
		})
		log.Println("__________________________")

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateDB(db *bolt.DB, err error, path string) {
	fmt.Println("::::: UPDATE Database :::::")

	err = db.Update(func(tx *bolt.Tx) error {
		// Create a bucket to store data in.
		b, err := tx.CreateBucketIfNotExists([]byte("functions"))
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		err = b.Put([]byte(strings.TrimLeft(file.Name(), "src/")), data)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("::::: boltdb Database :::::")

	db, err := bolt.Open("nodeDB.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		var cmd string
		fmt.Scanln(&cmd)

		switch cmd {
		case "init":
			var path string
			fmt.Print("enter Path : ")
			fmt.Scanln(&path)
			InitDB(db, err, path)
		case "update":
			var path string
			fmt.Print("enter Path : ")
			fmt.Scanln(&path)
			UpdateDB(db, err, path)
		case "create":
			var path string
			fmt.Print("enter Path : ")
			fmt.Scanln(&path)
			ConvertJS(db, err, path)
		case "view":
			ViewAllDB(db, err)
		}
	}
}
