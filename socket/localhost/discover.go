package localhost

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

const (
	udpPort    = 4000
	bufferSize = 1024
)

type NodeInfo struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

func startNode(ip string) {
	localAddr := net.UDPAddr{
		Port: udpPort,
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
			fmt.Println(n, addr)
			if err != nil {
				fmt.Println("Error reading from UDP:", err)
				continue
			}

			var nodeInfo NodeInfo
			err = json.Unmarshal(buffer[:n], &nodeInfo)
			if err != nil {
				fmt.Println("Error unmarshaling node info:", err)
				continue
			}

			fmt.Printf("Received node info from %s: %s:%d\n", addr, nodeInfo.IP, nodeInfo.Port)
		}
	}()

	broadcastAddr := net.UDPAddr{
		Port: udpPort,
		IP:   net.ParseIP("255.255.255.255"),
	}

	nodeInfo := NodeInfo{
		IP:   ip,
		Port: udpPort,
	}

	for {
		data, _ := json.Marshal(nodeInfo)
		_, err := conn.WriteToUDP(data, &broadcastAddr)
		if err != nil {
			fmt.Println("Error broadcasting node info:", err)
		}

		time.Sleep(5 * time.Second)
	}
}
