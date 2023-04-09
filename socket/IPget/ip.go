package IPget

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

// == private(internal) ip

func GetPrivateIp() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}

// == public(external) ip

type IPInfo struct {
	Origin string `json:"origin"`
}

func GetPublicIp() {
	resp, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var ipInfo IPInfo
	err = json.Unmarshal(body, &ipInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ipInfo.Origin)
}
