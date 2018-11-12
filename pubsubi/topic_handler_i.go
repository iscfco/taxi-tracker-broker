package pubsubi

import "taxi-tracker-broker/model"

type TopicHandlerI interface {
	Subscribe(topic *string, subscriber *model.SocketConnection)
	Publish(topic, msg *string)
	Unsubscribe(topic, clientId *string)
	UnsubscribeInAllTopics(clientId *string)
}
