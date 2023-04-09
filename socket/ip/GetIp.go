package ip

// return 된 string이 ""이라면 error
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

// == private(internal) ip

func GetPrivateIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// == public(external) ip

type IPInfo struct {
	Origin string `json:"origin"`
}

func GetPublicIp() string {
	resp, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var ipInfo IPInfo
	err = json.Unmarshal(body, &ipInfo)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(ipInfo.Origin)
	return ipInfo.Origin
}
