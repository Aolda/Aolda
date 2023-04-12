package localhost

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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
	address := "localhost"
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return openPort != "" && address != ""
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	HandleErr(err)
	initPeer(conn, address, openPort)
}

func AddPeer(address, port, openPort string) {
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort[1:]), nil)
	HandleErr(err)
	initPeer(conn, address, port)
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
		mparts := strings.Split(string(m), ":")
		if mparts[0] == "file_ready" {
			fmt.Printf("mparts[1]: ")
			fmt.Println(mparts[1])
			receiveFile(p, mparts[1])
		} else {
			peerList := []Peer{}
			json.Unmarshal(m, &peerList)
			for _, v := range peerList {
				peerPort := fmt.Sprintf(":%s", v.Port)
				AddPeer(payload.Address, payload.Port, peerPort)
			}
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

func RestStart(aPort int) {
	port = fmt.Sprintf(":%d", aPort)
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware, loggerMiddleware)
	router.HandleFunc("/ws", Upgrade).Methods("GET")
	router.HandleFunc("/peers", peersAPI).Methods("GET", "POST")
	router.HandleFunc("/dbsync", dbsyncAPI).Methods("GET", "POST")
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// == cli

func usage() {
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:		Set the PORT of the server\n\n")
	os.Exit(0)
}

func CliStart() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	flag.Parse()

	RestStart(*port)
}
