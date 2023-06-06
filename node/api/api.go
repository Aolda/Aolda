package api

import (
	transaction "aolda_node/blockchain"
	"aolda_node/compiler"
	file "aolda_node/ipfs"
	pubsub "aolda_node/p2p"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rs/cors"
)

type bodyObject struct {
	FileHash     string   `json:"fileHash"`
	FunctionName string   `json:"functionName"`
	Args         []string `json:"args"`
}

func Listening() {
	http.HandleFunc("/emit", emitHandler)
	http.HandleFunc("/upload", uploadHandler)
	c := cors.Default().Handler(http.DefaultServeMux)
	log.Fatal(http.ListenAndServe(":8000", c))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		bodyFile, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Failed to retrieve file", http.StatusBadRequest)
			return
		}
		defer bodyFile.Close()

		fileHash := file.IpfsAdd(handler.Filename)

		resTx, _ := transaction.MakeFileTx(fileHash)
		pubsub.NotifyNewTx(resTx)

		jsonData, err := json.Marshal(resTx)
		if err != nil {
			log.Fatal(err)
		}

		// type 3 만들기
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(jsonData))
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func emitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		var bodyData bodyObject
		err = json.Unmarshal(body, &bodyData)
		if err != nil {
			http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
			return
		}

		bodyTx, _ := transaction.MakeAPICallTx(bodyData.FileHash, bodyData.FunctionName, bodyData.Args)
		pubsub.NotifyNewTx(bodyTx)

		file.IpfsGet(bodyData.FileHash)
		res := compiler.ExecuteJS(bodyData.FileHash, bodyData.FunctionName, bodyData.Args)

		resTx, _ := transaction.MakeCofirmTx(bodyData.FileHash, bodyData.FunctionName, res, bodyData.Args)
		pubsub.NotifyNewTx(resTx)

		// type 3 만들기
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, res)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
