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
)

type MyData struct {
	fileHash     string   `json:"fileHash"`
	functionName string   `json:"functionName"`
	args         []string `json:"args"`
}

func Listening() {
	http.HandleFunc("/emit", emitHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func emitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		var bodyData MyData
		err = json.Unmarshal(body, &bodyData)
		if err != nil {
			http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
			return
		}

		bodyTx, _ := transaction.MakeAPICallTx(bodyData.fileHash, bodyData.functionName, bodyData.args)
		pubsub.NotifyNewTx(bodyTx)

		file.IpfsGet(bodyData.fileHash)
		res := compiler.ExecuteJS(bodyData.fileHash, bodyData.functionName, bodyData.args)

		resTx, _ := transaction.MakeCofirmTx(bodyData.fileHash, bodyData.functionName, res, bodyData.args)
		pubsub.NotifyNewTx(resTx)

		// type 3 만들기
		fmt.Fprintf(w, res)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
