package localhost

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

var dbpayload dbsyncPayload

type dbsyncPayload struct {
	Address, Port, FileName string
	Type                    int
}

func findPeerByAddressAndPort(address, port string) (*Peer, error) {
	for _, peer := range Peers.v {
		if peer.Address == address && peer.Port == port {
			return peer, nil
		}
	}
	return nil, errors.New("peer not found")
}

func dbsync(address, filename, port string) {

	// 원하는 파일 전송
	fileparts := strings.Split(filename, ",")
	for i := range fileparts {
		fileparts[i] = strings.TrimSpace(fileparts[i])
		peer, err := findPeerByAddressAndPort(address, port)
		if err != nil {
			HandleErr(err)
		}
		fmt.Printf("fileparts: ")
		fmt.Println(fileparts[i])
		sendFile(peer, fileparts[i])
	}
	fmt.Println(fileparts)
}

func sendFile(p *Peer, filename string) {
	// 파일을 읽기용으로 연 파일 객체를 반환
	path := "../node/src/" // main.go 파일 기준 경로
	filepath := path + filename
	file, err := os.Open(filepath)
	HandleErr(err)
	defer file.Close()

	readyMsg := "file_ready" + ":" + filename
	err = p.Conn.WriteMessage(websocket.TextMessage, []byte(readyMsg))
	HandleErr(err)

	// 파일 내용을 WebSocket 연결을 통해 보냄
	buffer := make([]byte, 1024)
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				// 파일의 끝에 도달한 경우
				err = p.Conn.WriteMessage(websocket.TextMessage, []byte("file_eof:"+filename))
				HandleErr(err)
				fmt.Println("End of file reached.")
			} else {
				fmt.Println("Error reading file:", err)
			}
			break
		}
		err = p.Conn.WriteMessage(websocket.BinaryMessage, buffer[:bytesRead])
		HandleErr(err)
	}

	fmt.Println("File", filepath, "sent.")
}

func receiveFile(p *Peer, filename string) {
	// 파일을 받아서 저장할 경로와 파일 이름을 결정
	path := "../node/src/"
	filepath := path + filename

	// 파일을 열고 쓰기용으로 연 파일 객체를 반환
	file, err := os.Create(filepath)
	HandleErr(err)
	defer file.Close()

	// 받은 데이터를 파일에 씀
	for {
		messageType, data, err := p.Conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				fmt.Println("Peer connection closed normally.")
			} else {
				fmt.Println("Peer connection closed with error:", err)
			}
			break
		}
		if messageType == websocket.TextMessage && string(data) == "file_eof:"+filename {
			fmt.Println("End of file reached.")
			break
		}
		_, err = file.Write(data)
		HandleErr(err)
	}
	fmt.Println("File", filename, "received and saved to", filepath)
}

func readSrcAndSendList(rw http.ResponseWriter) {
	files, err := ioutil.ReadDir("../node/src")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	var fileInfos []map[string]interface{}
	for _, f := range files {
		fileInfo := make(map[string]interface{})
		fileInfo["name"] = f.Name()
		fileInfo["size"] = f.Size()
		fileInfo["modified"] = f.ModTime().String()
		fileInfos = append(fileInfos, fileInfo)
	}

	fileInfoJson, err := json.Marshal(fileInfos)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	// JSON 형식의 파일 리스트를 반환합니다.
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(fileInfoJson)
}

func dbsyncAPI(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST": // DB에서 원하는 파일을 전송
		json.NewDecoder(r.Body).Decode(&dbpayload)
		dbsync(dbpayload.Address, dbpayload.FileName, dbpayload.Port)
		rw.WriteHeader(http.StatusOK)

	case "GET": // src 파일 list 전송
		readSrcAndSendList(rw)
	default:
		http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
	}
}