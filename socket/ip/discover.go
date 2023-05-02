package ip

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

const (
	portc      = 3000
	bufferSize = 1024
)

type NodeInfo struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

type Message struct {
	Protocol string   `json:"protocol"`
	Payload  NodeInfo `json:"payload"`
}

func startNode(ip string, port int) {
	localAddr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(ip),
	}

	conn, err := net.ListenUDP("udp", &localAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	go func() {
		buffer := make([]byte, bufferSize)
		for {
			n, addr, err := conn.ReadFromUDP(buffer)
			if err != nil {
				fmt.Println("Error reading from UDP:", err)
				continue
			}

			var message Message
			err = json.Unmarshal(buffer[:n], &message)
			if err != nil {
				fmt.Println("Error unmarshaling message:", err)
				continue
			}

			if message.Protocol == "Hello" {
				nodeInfo := message.Payload
				fmt.Printf("Received Hello from %s: %s:%d\n", addr, nodeInfo.IP, nodeInfo.Port)
			}
		}
	}()

	broadcastAddr := net.UDPAddr{
		Port: portc,
		IP:   net.ParseIP("255.255.255.255"),
	}

	nodeInfo := NodeInfo{
		IP:   ip,
		Port: portc,
	}

	message := Message{
		Protocol: "Hello",
		Payload:  nodeInfo,
	}

	for {
		data, _ := json.Marshal(message)
		_, err := conn.WriteToUDP(data, &broadcastAddr)
		if err != nil {
			fmt.Println("Error broadcasting Hello message:", err)
		}

		time.Sleep(5 * time.Second)
	}
}
