package ip

// return 된 string이 ""이라면 error
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/pion/stun"
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

func GetPublicIPAndPort() (string, int, error) {
	c, err := stun.Dial("udp", "stun.l.google.com:19302")
	if err != nil {
		return "", 0, err
	}

	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)

	var ipStr string
	var port int

	err = c.Do(message, func(res stun.Event) {
		if res.Error != nil {
			err = res.Error
			return
		}

		var xorAddr stun.XORMappedAddress
		if err := xorAddr.GetFrom(res.Message); err != nil {
			err = err
			return
		}

		ipStr = xorAddr.IP.String()
		port = xorAddr.Port
	})

	fmt.Println(ipStr)
	fmt.Println(port)
	return ipStr, port, err
}
