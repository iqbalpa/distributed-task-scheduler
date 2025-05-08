package coordinator

import (
	"fmt"
	"main/internal/task"
	"sync"
)

type Coordinator struct {
	mu sync.Mutex
	Tasks map[int]*task.Task
	Queue []*task.Task
	NumElems int
}

func New() *Coordinator {
	return &Coordinator{
		mu: sync.Mutex{},
		Queue: []*task.Task{},
		Tasks: make(map[int]*task.Task),
		NumElems: 0,
	}
}

func (c *Coordinator) Add(t *task.Task) {
	c.mu.Lock()
	defer c.mu.Unlock() 
	
	t.Id = c.NumElems
	t.Status = task.Pending
	c.Queue = append(c.Queue, t)
	c.Tasks[c.NumElems] = t
	c.NumElems++
}

func (c *Coordinator) GetStatus(id int) (task.Status, error) {
	c.mu.Lock()
	defer c.mu.Unlock() 

	t, ok := c.Tasks[id]
	if !ok {
		return "", fmt.Errorf("task not found")
	}
	return t.Status, nil
}

func (c *Coordinator) GetAll() ([]*task.Task, error) {
	ret := []*task.Task{}
	for _,m := range c.Tasks {
		ret = append(ret, m)
	}
	return ret, nil
}

func (c *Coordinator) NextTask() (*task.Task, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if len(c.Queue) == 0 {
		return nil, fmt.Errorf("there is no tasks")
	}
	ret := c.Queue[0]
	c.Queue = c.Queue[1:]
	return ret, nil
}
