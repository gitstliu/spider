package main

import (
	"decoder/csdnreadcount"
	"ipproxyv2"
	"jobscheduler"
	"time"

	"github.com/gitstliu/go-commonfunctions"
	"github.com/gitstliu/log4go"
)

func main() {
	count := 1
	defer commonfunctions.PanicHandler()

	log4go.LoadConfiguration("spiderconfig/log.xml")

	go ipproxyv2.FlushIPPool()

	time.Sleep(3 * time.Second)

	scheduler := jobscheduler.CreateScheduler(count)
	//	log4go.Debug("Add TaskQueue")
	for i := 0; i < count; i++ {
		scheduler.TaskQueue.AddTask(csdnreadcount.MainTask)
	}

	//	log4go.Debug("Add TaskMap")
	//	scheduler.TaskMap.AddTask(topbaidu.MainTask.Name, topbaidu.MainTask.Decoder)
	scheduler.Run()

	for true {
		time.Sleep(1 * time.Second)
	}
}
