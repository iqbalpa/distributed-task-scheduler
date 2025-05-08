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

func (c *Coordinator) Init() {
	c.Queue = []*task.Task{}
	c.NumElems = 0
	c.Tasks = make(map[int]*task.Task)
	c.mu = sync.Mutex{}
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
