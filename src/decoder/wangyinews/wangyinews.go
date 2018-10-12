package wangyinews

import (
	"jobqueue"
	"jobtask"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var MainTask = &jobtask.Task{Name: "wangyi_news", Url: "https://news.163.com/", Decoder: MainDecoder}

func MainDecoder(doc *goquery.Document, resp *http.Response) map[string]interface{} {

}
