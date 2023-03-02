package database

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/fsnotify/fsnotify"
)

func initDB(db *bolt.DB, err error, dir string) {
	log.Println("INIT DB")

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

func viewAllDB(db *bolt.DB, err error) {
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

func updateDB(db *bolt.DB, err error, path string) {
	log.Println("UODATE DB")
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

func Boltdb() {
	fmt.Println("::::: boltdb Database :::::")

	db, err := bolt.Open("nodeDB.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dir := "./src"
	initDB(db, err, dir)
	viewAllDB(db, err)

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			// Watch the directory for changes.
			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				log.Fatal(err)
			}
			defer watcher.Close()

			done := make(chan bool)
			go func() {
				for {
					select {
					case event := <-watcher.Events:
						updateDB(db, err, event.Name)
						viewAllDB(db, err)
					case err := <-watcher.Errors:
						log.Println("error:", err)
					}
				}
			}()

			err = watcher.Add(path)
			if err != nil {
				log.Fatal(err)
			}
			<-done
		}

		return nil
	})
}
