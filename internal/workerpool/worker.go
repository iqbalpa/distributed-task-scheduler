package worker

import (
	"fmt"
	"main/internal/task"
	"time"
)

type Worker struct{
	Id int
}

func (w *Worker) Process(tc chan *task.Task) {
	for t := range tc {
		fmt.Printf("Worker %d is processing task ID: %d\n", w.Id, t.Id)
		t.Status = task.Running
		time.Sleep(2 * time.Second)
		t.Status = task.Success
	}
}