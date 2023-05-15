package localhost

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

const (
	localPort          = ":5001"
	discoverNodeProto  = "DISCOVER_NODE_REQUEST"
	discoverInterval   = 5 * time.Second
	connectionProtocol = "udp"
)

func handleIncomingMessages(conn *net.UDPConn) {
	buffer := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}

		message := string(buffer[:n])
		if strings.TrimSpace(message) == "DISCOVER_NODE_REQUEST" {
			fmt.Printf("Received DISCOVER_NODE_REQUEST from %s\n", addr.String())

			response := "NODE_RESPONSE: " + localPort
			_, err = conn.WriteToUDP([]byte(response), addr)
			if err != nil {
				fmt.Println("Error sending response:", err)
			}
		}
	}
}

func discoverNodeslocal(conn *net.UDPConn) {
	for {
		localIPs := getLocalIPs()
		discoveredNodes := []string{}
		for _, localIP := range localIPs {
			nodes := getNodesByProtocolLocal(discoverNodeProto, localIP)
			discoveredNodes = append(discoveredNodes, nodes...)
		}

		for _, node := range discoveredNodes {
			fmt.Println("Discovered local node:", node)
		}
		time.Sleep(discoverInterval)
	}
}

func getNodesByProtocolLocal(protocol string, ipAddress string) []string {
	cmd := exec.Command("netstat", "-tuln")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error executing netstat command:", err)
		return nil
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting netstat command:", err)
		return nil
	}

	scanner := bufio.NewScanner(stdout)
	listeningPattern := regexp.MustCompile(`^` + connectionProtocol + `.*LISTEN\s*$`)
	nodes := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if listeningPattern.MatchString(line) {
			fields := strings.Fields(line)
			address := fields[3]

			// Check if the address has the desired protocol and IP address
			if !strings.HasPrefix(address, ipAddress) {
				continue
			}

			conn, err := net.DialTimeout(connectionProtocol, address, time.Second)
			if err != nil {
				continue
			}
			defer conn.Close()

			conn.SetReadDeadline(time.Now().Add(2 * time.Second))
			_, err = conn.Write([]byte(protocol))
			if err != nil {
				continue
			}

			buffer := make([]byte, 1024)
			n, err := conn.Read(buffer)
			if err != nil {
				continue
			}

			response := string(buffer[:n])
			if strings.Contains(response, protocol) {
				nodes = append(nodes, address)
			}
		}
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting for netstat command:", err)
	}

	return nodes
}

func getLocalIPs() []string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error getting local IPs:", err)
		return nil
	}

	var localIPs []string
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}
		if ip != nil && !ip.IsLoopback() && ip.To4() != nil {
			localIPs = append(localIPs, ip.String())
		}
	}
	return localIPs
}

func Local_Discover_main() {
	addr, err := net.ResolveUDPAddr("udp4", localPort)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	go handleIncomingMessages(conn)
	go discoverNodeslocal(conn)

	// Keep the program running
	select {}
}
