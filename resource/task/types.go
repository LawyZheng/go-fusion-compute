package task

type Status string

const (
	WAITING    Status = "waiting"
	RUNNING    Status = "running"
	SUCCESS    Status = "success"
	FAILED     Status = "failed"
	CANCELLING Status = "cancelling"
)

type Task struct {
	Urn        string            `json:"urn"`
	Uri        string            `json:"uri"`
	Type       string            `json:"type"`
	EntityUrn  string            `json:"entityUrn"`
	EntityName string            `json:"entityName"`
	StartTime  string            `json:"startTime"`
	FinishTime string            `json:"finishTime"`
	User       string            `json:"user"`
	Status     Status            `json:"status"`
	Progress   int               `json:"progress"`
	Reason     string            `json:"reason"`
	ReasonDes  string            `json:"reasonDes"`
	Params     map[string]string `json:"params"` // reserved field
}

func (t *Task) IsDone() bool {
	return t.Status == SUCCESS || t.Status == FAILED
}

type TaskResult struct {
	Task
	Err error
}
