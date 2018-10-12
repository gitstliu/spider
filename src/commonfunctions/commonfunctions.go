package commonfunctions

import (
	"bytes"
	"net/http"
)

func CreateSimpleHttpRequest(Method string, Url string, Body []byte) *http.Request {
	request, err := http.NewRequest(Method, Url, bytes.NewBuffer(Body))

	if err != nil {
		panic(err)
	}
	return request
}
