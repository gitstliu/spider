package ipproxyv2

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/gitstliu/log4go"
)

var ippool = &IPProxy{}

type IPProxy struct {
	Result []string
	Length int
}

func GetProxyUrl() (string, error) {

	ipproxy := ippool
	if ipproxy.Length > 0 {
		pos := rand.Intn(len(ipproxy.Result))
		return ipproxy.Result[pos], nil
	}

	return "", errors.New("IP pool length is 0")
}

func GetIPPool() *IPProxy {
	return ippool
}

func FlushIPPool() {
	for true {
		response, getIPError := http.Get("http://10.67.51.93:10000/proxy_ips")
		if getIPError != nil {
			log4go.Error(getIPError)
			return
		}

		body, readBodyErr := ioutil.ReadAll(response.Body)

		if readBodyErr != nil {
			log4go.Error(readBodyErr)
			return
		}

		ipproxy := &IPProxy{}
		//		fmt.Println(string(body))

		decodeIPProxyError := json.Unmarshal(body, ipproxy)

		if decodeIPProxyError != nil {
			log4go.Error(decodeIPProxyError)
			return
		}
		//	log4go.Debug(ipproxy.Message[0].Content)
		//	return url.Parse(ipproxy.Message[0].Content)

		ipproxy.Length = len(ipproxy.Result)
		log4go.Info("ipproxy.Length = %v", ipproxy.Length)
		ippool = ipproxy

		//		fmt.Println(*ippool)

		time.Sleep(10 * time.Second)
	}
}
