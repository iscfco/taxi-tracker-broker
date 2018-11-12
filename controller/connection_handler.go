package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"taxi-tracker-broker/constant"
	"taxi-tracker-broker/model"
	"taxi-tracker-broker/model/payload"
	"taxi-tracker-broker/pubsubi"
	"time"
)

type ConnectionHandler struct {
	topicHandler     pubsubi.TopicHandlerI
	socketConnection *model.SocketConnection
}

func NewConnectionHandler(topicHandler pubsubi.TopicHandlerI, conn *model.SocketConnection) ConnectionHandler {
	return ConnectionHandler{
		topicHandler:     topicHandler,
		socketConnection: conn,
	}
}

func (connectionHandler *ConnectionHandler) OnOpen() {
	fmt.Println("New connection created")
}

func (connectionHandler *ConnectionHandler) OnMessage(msg *string) {
	fmt.Println("*** Received:", *msg)
	message := model.Message{}
	err := json.Unmarshal([]byte(*msg), &message)
	if err != nil {
		payload := []byte("error parsing response in server. ")
		connectionHandler.socketConnection.OutPut <- payload
		return
	}

	switch message.TaskType {
	case constant.Subscribe:
		subBody := message.Payload.(map[string]interface{})
		topic := subBody["topic"].(string)
		connectionHandler.topicHandler.Subscribe(&topic, connectionHandler.socketConnection)
	case constant.Publish:
		subBody := message.Payload.(map[string]interface{})
		topic := subBody["topic"].(string)
		msg := subBody["message"].(interface{})
		finalMsg := model.Message{
			TaskType: constant.IncomingMessage,
			Payload: payload.IncomingMessage{
				TopicFrom: topic,
				Message:   msg,
			},
		}
		finalMsgSByte, _ := json.Marshal(finalMsg)
		finalMsgStr := string(finalMsgSByte)
		connectionHandler.topicHandler.Publish(&topic, &finalMsgStr)
	case constant.Unsubscribe:
		subBody := message.Payload.(map[string]interface{})
		topic := subBody["topic"].(string)
		connectionHandler.topicHandler.Unsubscribe(&topic, &connectionHandler.socketConnection.ClientId)
	}
}

func (connectionHandler *ConnectionHandler) OnWrite() {
	defer connectionHandler.socketConnection.WS.Close()
	var msgToSend []byte
	for {
		msgToSend = <-connectionHandler.socketConnection.OutPut
		err := connectionHandler.socketConnection.WS.WriteMessage(websocket.TextMessage, msgToSend)
		if err != nil {
			return
		}
	}
}

func (connectionHandler *ConnectionHandler) OnClose() {
	time.Sleep(20 * time.Millisecond)
	connectionHandler.socketConnection.WS.Close()
	connectionHandler.topicHandler.UnsubscribeInAllTopics(&connectionHandler.socketConnection.ClientId)
}
