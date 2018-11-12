package payload

type TaskResult struct {
	ResultCode string `json:"result_code"`
	Message    string `json:"message"`
	TaskId     string `json:"task_id"`
}
