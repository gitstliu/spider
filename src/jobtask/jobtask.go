package jobtask

import (
	"net/http"
	"sync"

	xmlpath "gopkg.in/xmlpath.v2"
)

type DecodeDOM func(body *xmlpath.Node, resp *http.Response) (map[string]interface{}, []*Task)

type Task struct {
	Name    string
	Request *http.Request
	Decoder DecodeDOM
}

type TaskMap struct {
	Tasks *sync.Map
}

func CreateTaskMap() *TaskMap {
	return &TaskMap{Tasks: &sync.Map{}}
}

func (this *TaskMap) AddTask(name string, decoder DecodeDOM) {
	newTask := &Task{Name: name, Decoder: decoder}
	this.Tasks.Store(name, newTask)
}

func (this *TaskMap) GetTask(name string) *Task {
	task, exist := this.Tasks.Load(name)
	if exist {
		return task.(*Task)
	} else {
		return nil
	}
}
