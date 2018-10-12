package jobqueue

import (
	"jobtask"
)

type TaskQueue struct {
	Queue chan *jobtask.Task
}

func CreateTaskQueue() *TaskQueue {
	return &TaskQueue{Queue: make(chan *jobtask.Task, 1000000)}
}

func (this *TaskQueue) PopTask() *jobtask.Task {
	return <-this.Queue
}

func (this *TaskQueue) AddTask(task *jobtask.Task) {
	this.Queue <- task
	return
}
