package worker

import (
	"main/cmd/coordinator"
	"main/internal/task"
	"time"
)

type WorkerPool struct {
	Size int
	TaskChan chan *task.Task
}

func New(size int) *WorkerPool {
	return &WorkerPool{
		Size: size,
		TaskChan: make(chan *task.Task),
	}
}

func (p *WorkerPool) Start(c *coordinator.Coordinator) {
	for i := range p.Size {
		w := Worker{Id: i}
		go w.Process(p.TaskChan)
	}
	go p.Dispatch(c, p.TaskChan)
}

func (p *WorkerPool) Dispatch(c *coordinator.Coordinator, tc chan *task.Task) {
	for {
		t, err := c.NextTask()
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		tc <- t
	}
}
