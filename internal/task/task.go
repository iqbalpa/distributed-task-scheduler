package task

type Task struct {
	Id     int
	Type   string
	Params map[string]string
	Status Status
}

type Status string

const (
	Pending Status = "pending"
	Running Status = "running"
	Failed  Status = "failed"
	Success Status = "success"
)