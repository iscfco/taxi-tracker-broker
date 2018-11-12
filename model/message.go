package model

type Message struct {
	TaskType string      `json:"task_type"`
	Payload  interface{} `json:"payload"`
}
