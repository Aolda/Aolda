package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		// IP 주소와 포트 번호
		ip := "192.168.0.100"
		port := "8080"

		// 웹소켓 연결을 초기화합니다.
		initPeer(conn, ip, port)
	})

	fmt.Println("Server started on port 8000")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func initPeer(conn *websocket.Conn, ip string, port string) {
	fmt.Println("Connecting to", ip, "on port", port)

	// IP 주소와 포트 번호를 이용하여 웹소켓 연결을 초기화합니다.
	peerConn, _, err := websocket.DefaultDialer.Dial("ws://"+ip+":"+port+"/ws", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 웹소켓 연결이 성공적으로 이루어졌을 때 수행할 동작을 여기에 작성합니다.
	fmt.Println("Connected to", ip, "on port", port)

	// 연결된 웹소켓을 계속 유지합니다.
	go func() {
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}()

	// 연결된 웹소켓을 계속 유지합니다.
	go func() {
		for {
			_, _, err := peerConn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}()
}
