package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/rs/cors"
)

var transactionCount prometheus.Counter // 전역 변수로 이동
var subCount prometheus.Counter         // 전역 변수로 이동

type bodyObject struct {
	FileHash     string   `json:"fileHash"`
	FunctionName string   `json:"functionName"`
	Args         []string `json:"args"`
}

func Listening() {
	registry := prometheus.NewRegistry()

	// httpRequestsTotal := prometheus.NewCounter(
	// 	prometheus.CounterOpts{
	// 		Name: "http_requests_total",
	// 		Help: "Total number of HTTP requests",
	// 	},
	// )
	// registry.MustRegister(httpRequestsTotal)

	// transactionCount 카운터 메트릭 생성 및 등록
	// transactionCount = prometheus.NewCounter(
	// 	prometheus.CounterOpts{
	// 		Name: "transaction_count",
	// 		Help: "Total number of transactions",
	// 	},
	// )
	// registry.MustRegister(transactionCount)

	subCount = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "transaction_count",
			Help: "Total number of transactions",
		},
	)
	registry.MustRegister(subCount)

	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	// http.HandleFunc("/emit", emitHandler)
	// http.HandleFunc("/upload", uploadHandler)
	c := cors.Default().Handler(http.DefaultServeMux)
	log.Fatal(http.ListenAndServe(":8000", c))
}

func saveFile(file io.Reader, filename string) error {
	dst, err := os.Create("./src/" + filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}

	return nil
}

func truncateString(str, keyword string) string {
	splits := strings.SplitN(str, keyword, 2)
	if len(splits) > 1 {
		return splits[1]
	}
	return str
}

// func uploadHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		bodyFile, handler, err := r.FormFile("file")
// 		if err != nil {
// 			http.Error(w, "Failed to retrieve file", http.StatusBadRequest)
// 			return
// 		}

// 		defer bodyFile.Close()

// 		err = saveFile(bodyFile, handler.Filename)
// 		if err != nil {
// 			http.Error(w, "Failed to save file", http.StatusInternalServerError)
// 			return
// 		}

// 		fileHash := file.IpfsAdd(handler.Filename)

// 		resTx, _ := transaction.MakeFileTx(fileHash)
// 		resTx.Body.FileHash = truncateString(resTx.Body.FileHash, "about")
// 		pubsub.NotifyNewTx(resTx)
// 		fmt.Println(resTx)

// 		jsonData, err := json.Marshal(resTx)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		// type 3 만들기
// 		w.Header().Set("Content-Type", "application/json")
// 		fmt.Fprintf(w, string(jsonData))
// 	} else {
// 		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 	}
// }

// func emitHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		body, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			http.Error(w, "Error reading request body", http.StatusInternalServerError)
// 			return
// 		}

// 		var bodyData bodyObject
// 		err = json.Unmarshal(body, &bodyData)
// 		if err != nil {
// 			http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
// 			return
// 		}

// 		bodyTx, _ := transaction.MakeAPICallTx(bodyData.FileHash, bodyData.FunctionName, bodyData.Args)
// 		pubsub.NotifyNewTx(bodyTx)

// 		file.IpfsGet(bodyData.FileHash)
// 		res := compiler.ExecuteJS(bodyData.FileHash, bodyData.FunctionName, bodyData.Args)

// 		resTx, _ := transaction.MakeCofirmTx(bodyData.FileHash, bodyData.FunctionName, res, bodyData.Args)
// 		pubsub.NotifyNewTx(resTx)

// 		jsonData, err := json.Marshal(resTx)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		fmt.Fprintf(w, string(jsonData))
// 	} else {
// 		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 	}
// }

func TransactionCounter() {
	fmt.Println("transactionCount")
	transactionCount.Inc()
}

func SubCounter() {
	subCount.Inc()
}
