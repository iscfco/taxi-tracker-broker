package pubsubimp

import (
	"fmt"
	"taxi-tracker-broker/model"
	"taxi-tracker-broker/pubsubi"
)

type TopicHandler struct {
	topicBag     map[string]map[string]*model.SocketConnection
	clientTopics map[string][]string
}

func NewTopicHandler() pubsubi.TopicHandlerI {
	newTopicContainer := TopicHandler{}
	newTopicContainer.topicBag = make(map[string]map[string]*model.SocketConnection)
	newTopicContainer.clientTopics = make(map[string][]string)
	return &newTopicContainer
}

func (t *TopicHandler) Subscribe(topic *string, subscriber *model.SocketConnection) {
	_, ok := t.topicBag[*topic]
	if !ok {
		t.topicBag[*topic] = make(map[string]*model.SocketConnection)
	}
	t.topicBag[*topic][subscriber.ClientId] = subscriber
	fmt.Println(t.topicBag)
	for _, clientTopic := range t.clientTopics[subscriber.ClientId] {
		if clientTopic == *topic {
			goto L
		}
	}
	t.clientTopics[subscriber.ClientId] = append(t.clientTopics[subscriber.ClientId], *topic)
L:
	fmt.Println(t.clientTopics)
}

func (t *TopicHandler) Publish(topic, msg *string) {
	_, ok := t.topicBag[*topic]
	if ok {
		for _, subscriber := range t.topicBag[*topic] {
			subscriber.OutPut <- []byte(*msg)
		}
	}
	fmt.Println("*** $$ Publisher:", *msg, " to topic:", *topic)
}

func (t *TopicHandler) Unsubscribe(topic, clientId *string) {
	_, ok := t.topicBag[*topic]
	if ok {
		delete(t.topicBag[*topic], *clientId)
	}
	subscribersLen := len(t.topicBag[*topic])
	if subscribersLen == 0 {
		delete(t.topicBag, *topic)
	}
}

func (t *TopicHandler) UnsubscribeInAllTopics(clientId *string) {
	_, ok := t.clientTopics[*clientId]
	if ok {
		for _, topic := range t.clientTopics[*clientId] {
			t.Unsubscribe(&topic, clientId)
		}
	}
}
