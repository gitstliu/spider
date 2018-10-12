package decoder

//import (
//	"fmt"
//	"jobqueue"
//	"jobtask"
//	"net/http"

//	"github.com/PuerkitoBio/goquery"
//)

//var Task

//type DecodeFunc func(doc *goquery.Document, resp *http.Response) map[string]interface{}

//type Decoder struct {
//	CurrTask *jobtask.Task
//}

//mainTask := &jobtask.Task{Name:"百度风云榜",Url:"http://top.baidu.com/buzz?b=341&c=513&fr=topcategory_c513", Decoder:MainDecoder}

//func MainDecoder(doc *goquery.Document, resp *http.Response) map[string]interface{} {
//	fmt.Println(doc.Text())
//}
