package ip

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// == utils
func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// === websocket

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	openPort := r.URL.Query().Get("openPort")
	publicAddress := r.Header.Get("X-Real-IP")
	if publicAddress == "" {
		publicAddress = r.Header.Get("X-Forwarded-For")
		if publicAddress == "" {
			publicAddress = r.RemoteAddr
		}
	}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return r.Header.Get("Origin") == "http://"+publicAddress || r.Header.Get("Origin") == "https://"+publicAddress
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	HandleErr(err)
	initPeer(conn, publicAddress, openPort)
}

func AddPeer(publicAddress, port, openPort string) {
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", publicAddress, port, openPort[1:]), nil)
	HandleErr(err)
	initPeer(conn, publicAddress, port)
}

// ==========peer
type peers struct {
	v map[string]*Peer
	m sync.Mutex
}

var Peers peers = peers{
	v: make(map[string]*Peer),
}

type Peer struct {
	Key     string
	Address string
	Port    string
	Conn    *websocket.Conn
	Inbox   chan []byte
}

func AllPeers(p *peers) []string {
	p.m.Lock()
	defer p.m.Unlock()
	var keys []string
	for key := range p.v {
		keys = append(keys, key)
	}
	return keys
}

func (p *Peer) close() {
	Peers.m.Lock()
	defer Peers.m.Unlock()
	p.Conn.Close()
	delete(Peers.v, p.Key)
}

func initPeer(conn *websocket.Conn, address, port string) *Peer {
	key := fmt.Sprintf("%s:%s", address, port)
	p := &Peer{
		Key:     key,
		Address: address,
		Port:    port,
		Conn:    conn,
		Inbox:   make(chan []byte),
	}
	fmt.Println(p.Address)
	go p.readListener()
	go p.writeListener()
	Peers.v[key] = p
	p.write()
	return p
}

func (p *Peer) write() {
	var res []byte
	for i := range Peers.v {
		address := strings.Split(i, ":")
		marJson, _ := json.Marshal(struct {
			Address string
			Port    string
		}{
			Address: address[0],
			Port:    address[1],
		})
		res = append(res, marJson...)
	}
	p.Inbox <- res
}

func (p *Peer) writeListener() {
	defer p.close()
	for {
		m, ok := <-p.Inbox
		if !ok {
			break
		}
		p.Conn.WriteMessage(websocket.TextMessage, m)
	}
}

func (p *Peer) readListener() {
	defer p.close()
	for {
		_, m, err := p.Conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("msg : ", string(m))
		peerList := []Peer{}
		json.Unmarshal(m, &peerList)
		for _, v := range peerList {
			peerPort := fmt.Sprintf(":%s", v.Port)
			AddPeer(payload.Address, payload.Port, peerPort)
		}
	}
}

var port string
var payload addPeerPayload

type addPeerPayload struct {
	Address, Port string
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		next.ServeHTTP(rw, r)
	})
}

func peersAPI(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		json.NewDecoder(r.Body).Decode(&payload)
		AddPeer(payload.Address, payload.Port, port)
		rw.WriteHeader(http.StatusOK)

	case "GET":
		json.NewEncoder(rw).Encode(AllPeers(&Peers))
	}
}

func RestStart(aPort int, publicAddress string) {
	port = fmt.Sprintf(":%d", aPort)
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware, loggerMiddleware)
	router.HandleFunc("/ws", Upgrade).Methods("GET")
	router.HandleFunc("/peers", peersAPI).Methods("GET", "POST")
	fmt.Printf("Listening on http://%s%s\n", publicAddress, port)
	log.Fatal(http.ListenAndServe(publicAddress+port, router))
}

// == cli

func usage() {
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:		Set the PORT of the server\n\n")
	os.Exit(0)
}

func CliStart() {
	// if len(os.Args) == 1 {
	// 	usage()
	// }

	// ip := GetPrivateIp()
	// port := flag.Int("port", 4000, "Set port of the server")
	// flag.Parse()

	// //ip, port, err := GetPublicIPAndPort()
	// fmt.Printf("ip is %x\n", ip)
	// // if err != nil {
	// // 	fmt.Println("GetPublicIPAndPort ERROR")
	// // 	os.Exit(0)
	// // }

	// RestStart(port, ip)

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <local-ip-address> <port>")
		return
	}

	ip := GetPrivateIp()
	port = os.Args[2]
	portInt, err := strconv.Atoi(port) // port를 정수로 변환
	if err != nil {
		fmt.Println("Invalid port number")
		return
	}
	startNode(ip, portInt)

	// 이 부분을 추가합니다.
	aPort, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("Invalid port number. Please provide a valid integer for the port number.")
		return
	}
	RestStart(aPort, ip)

}