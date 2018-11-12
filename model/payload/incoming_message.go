package payload

type IncomingMessage struct {
	TopicFrom string      `json:"topic_from"`
	Message   interface{} `json:"message"`
}
