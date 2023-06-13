package p2p

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	myMetric = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "my_metric",
		Help: "내 커스텀 메트릭",
	},
		[]string{"TransactionName", "Time"},
	)

	collectedValues []Value
	mutex           sync.Mutex
)

type Value struct {
	TransactionName string
	Time            string
	Value           float64
}

func init() {
	prometheus.MustRegister(myMetric)
}

func collectValues(Transaction string) {
	println("collectValue")
	// 현재 시간 가져오기
	currentTime := time.Now().Format(time.RFC3339Nano)

	// 값과 함께 TransactionName과 Time을 저장
	storeValue(Value{
		TransactionName: Transaction,
		Time:            currentTime,
	})

}

func storeValue(value Value) {
	mutex.Lock()
	defer mutex.Unlock()

	// 값을 저장
	collectedValues = append(collectedValues, value)

	// myMetric에 값을 업데이트
	myMetric.With(prometheus.Labels{
		"TransactionName": value.TransactionName,
		"Time":            value.Time,
	}).Set(value.Value)
	print("collectValue: ")
	println(collectedValues)
}

func getCollectedValues() []Value {
	mutex.Lock()
	defer mutex.Unlock()

	// 수집된 값을 가져옴
	values := collectedValues

	// 수집된 값을 파일에 저장
	err := saveValuesToFile(values)
	if err != nil {
		fmt.Println("Failed to save values to file:", err)
	}

	// 수집된 값들을 초기화
	collectedValues = []Value{}

	return values
}

func saveValuesToFile(values []Value) error {
	// 값을 쉼표로 구분하여 문자열로 변환
	var lines []string
	for _, value := range values {
		line := fmt.Sprintf("%s,%s,%.2f", value.TransactionName, value.Time, value.Value)
		lines = append(lines, line)
	}
	valuesStr := strings.Join(lines, "\n")

	// 파일에 값들을 저장
	err := ioutil.WriteFile("./value/values.txt", []byte(valuesStr), 0644)
	if err != nil {
		return err
	}

	return nil
}
