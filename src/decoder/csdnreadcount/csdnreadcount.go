package csdnreadcount

import (
	"jobtask"
	"net/http"

	"commonfunctions"

	"github.com/gitstliu/log4go"
	xmlpath "gopkg.in/xmlpath.v2"
)

//var MainTask = &jobtask.Task{Name: "csdn_read_count", Url: "https://blog.csdn.net/love666666shen/article/details/72613143", Decoder: MainDecoder}
var MainTask = &jobtask.Task{Name: "csdn_read_count", Request: commonfunctions.CreateSimpleHttpRequest("GET", "https://blog.csdn.net/love666666shen/article/details/72613143", []byte{}), Decoder: MainDecoder}

func MainDecoder(doc *xmlpath.Node, resp *http.Response) (map[string]interface{}, []*jobtask.Task) {
	//	log4go.Debug(doc.Text())
	//log4go.Debug("Success")
	//	doc.xpath

	log4go.Info("Start")
	path := xmlpath.MustCompile("//*[@id=\"article_content\"]/div[2]/div[1]/div/span[12]/span[2]")
	it := path.Iter(doc)

	for it.Next() {
		log4go.Info("text = %v", it.Node().String())
	}

	return nil, nil
}
