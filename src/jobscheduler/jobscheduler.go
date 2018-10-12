package jobscheduler

import (
	"ipproxyv2"
	"jobqueue"
	"jobtask"
	"net/http"
	"time"

	"github.com/gitstliu/log4go"
	xmlpath "gopkg.in/xmlpath.v2"
)

type Scheduler struct {
	TaskMap       *jobtask.TaskMap
	TaskQueue     *jobqueue.TaskQueue
	MaxRoutineNum int
	IsRunning     bool
}

func CreateScheduler(maxRoutineNum int) *Scheduler {
	return &Scheduler{
		TaskMap:       jobtask.CreateTaskMap(),
		TaskQueue:     jobqueue.CreateTaskQueue(),
		MaxRoutineNum: maxRoutineNum,
		IsRunning:     false}
}

func (this *Scheduler) Run() {
	log4go.Debug("Run")
	if this.IsRunning {
		return
	}
	this.IsRunning = true

	for i := 0; i < this.MaxRoutineNum; i++ {
		log4go.Debug("Create Runnint Tasks")
		go this.runTask()
	}
}

func SendRequest(request *http.Request, proxyAddress string) (*http.Response, error) {

	//	proxy := func(_ *http.Request) (*url.URL, error) {
	//		return url.Parse(proxyAddress)
	//	}

	//	transport := &http.Transport{Proxy: proxy}

	//	client := &http.Client{
	//		Transport: transport,
	//		Timeout:   2 * time.Second,
	//	}

	//	return client.Get(address)

	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	return client.Do(request)
}

func (this *Scheduler) runTask() {

	for true {
		task := this.TaskQueue.PopTask()

		url, getUrlErr := ipproxyv2.GetProxyUrl()

		if getUrlErr != nil {
			this.TaskQueue.AddTask(task)
			continue
		}
		resp, respErr := SendRequest(task.Request, url)
		defer resp.Body.Close()

		if respErr != nil {
			this.TaskQueue.AddTask(task)
			continue
		}

		if resp.StatusCode != 200 {
			log4go.Error("[%v] run [%v] status is %v", url, task.Request.URL.String(), resp.StatusCode)
			continue
		}

		node, nodeErr := xmlpath.ParseHTML(resp.Body)

		if nodeErr != nil {
			log4go.Error(nodeErr)
			continue
		}

		datas, newTasks := task.Decoder(node, resp)

		for _, currTask := range newTasks {
			this.TaskQueue.AddTask(currTask)
		}

		log4go.Info("datas = %v", datas)

		log4go.Info("Success")
	}
}
