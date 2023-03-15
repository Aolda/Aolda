package main

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

func Splitter(s string, sep string, i int) string {
	r := strings.Split(s, sep)
	if len(r)-1 < i {
		return ""
	}
	return r[i]
}

// === websocket

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	// Port :3000 will upgrade the request from :4000
	openPort := r.URL.Query().Get("openPort")
	ip := Splitter(r.RemoteAddr, ":", 0)
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return openPort != "" && ip != ""
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	HandleErr(err)
	initPeer(conn, ip, openPort)
}

func AddPeer(address, port, openPort string) {
	// Port :4000 is requesting an upgrade from the port :3000
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort[1:]), nil)
	HandleErr(err)
	initPeer(conn, address, port)
}

// ==========peer
type peers struct {
	v map[string]*peer
	m sync.Mutex
}

var Peers peers = peers{
	v: make(map[string]*peer),
}

type peer struct {
	key     string
	address string
	port    string
	conn    *websocket.Conn
	inbox   chan []byte
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

func (p *peer) close() {
	Peers.m.Lock()
	defer Peers.m.Unlock()
	p.conn.Close()
	delete(Peers.v, p.key)
}
func initPeer(conn *websocket.Conn, address, port string) *peer {
	key := fmt.Sprintf("%s:%s", address, port)
	p := &peer{
		conn:    conn,
		inbox:   make(chan []byte),
		address: address,
		key:     key,
		port:    port,
	}
	go p.read()
	go p.write()
	Peers.v[key] = p
	return p
}
func (p *peer) write() {
	defer p.close()
	for {
		m, ok := <-p.inbox
		if !ok {
			break
		}
		p.conn.WriteMessage(websocket.TextMessage, m)
	}
}
func (p *peer) read() {
	// delete peer in case of error
	defer p.close()
	for {
		_, m, err := p.conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("%s", m)
	}
}

// ====== rest

var port string

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
		var payload addPeerPayload
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
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// == cli

func usage() {
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:		Set the PORT of the server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest'\n\n")
	os.Exit(0)
}

func CliStart() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		RestStart(*port)
	default:
		usage()
	}
}

func main() {
	CliStart()
}
