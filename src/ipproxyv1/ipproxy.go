package ipproxyv1

type IPProxy struct {
	Code    int64
	Message []*IPMessage
}

type IPMessage struct {
	Id                    int64
	Content               string
	AssessTimes           int
	SuccessTimes          int
	AvgResponseTime       float32
	ContinuousFailedTimes int
	Score                 float32
	InsertTime            int64
	UpdateTime            int64
}

func ProxyRequest(req *http.Request) (*url.URL, error) {
	response, getIPError := http.Get("http://10.0.192.59:9999/sql?query=SELECT%20*%20FROM%20valid_proxy%20ORDER%20BY%20RANDOM()%20limit%201")
	if getIPError != nil {
		log4go.Error(getIPError)
		return nil, getIPError
	}

	body, readBodyErr := ioutil.ReadAll(response.Body)

	if readBodyErr != nil {
		log4go.Error(readBodyErr)
		return nil, readBodyErr
	}

	ipproxy := &IPProxy{}
	//	log4go.Debug(string(body))

	decodeIPProxyError := json.Unmarshal(body, ipproxy)

	if decodeIPProxyError != nil {
		log4go.Error(decodeIPProxyError)
		return nil, decodeIPProxyError
	}
	//	log4go.Debug(ipproxy.Message[0].Content)
	return url.Parse(ipproxy.Message[0].Content)
}
