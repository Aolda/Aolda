package p2pTest

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var peers = make(map[string]*websocket.Conn)
var upgrader = websocket.Upgrader{}

func Test_ip() {
	port := "8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")[0]
		if _, ok := peers[ip]; !ok {
			conn, err := upgrader.Upgrade(w, r, nil)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("New peer connected:", ip)
			peers[ip] = conn
		}
	})

	log.Println("Listening on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
