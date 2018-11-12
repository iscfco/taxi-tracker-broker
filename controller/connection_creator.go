package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
	"taxi-tracker-broker/model"
	"taxi-tracker-broker/pubsubi"
)

type connectionCreator struct {
	topicHandler pubsubi.TopicHandlerI
}

func NewConnectionCreator(topicHandler pubsubi.TopicHandlerI) connectionCreator {
	return connectionCreator{
		topicHandler: topicHandler,
	}
}

func (c *connectionCreator) CreateConnection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("connection received, begin")
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  5120,
		WriteBufferSize: 5120,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	clientId := vars["clientId"]

	socketConn := &model.SocketConnection{
		WS:         ws,
		OutPut:     make(chan []byte),
		Authorized: false,
		ClientId:   clientId,
	}

	connectionHandler := NewConnectionHandler(c.topicHandler, socketConn)
	go connectionHandler.OnOpen()
	go connectionHandler.OnWrite()
	setOnMessage(&connectionHandler)
}

func setOnMessage(conn *ConnectionHandler) {
	defer conn.socketConnection.WS.Close()
	for {
		msgType, msg, err := conn.socketConnection.WS.ReadMessage()
		if err != nil {
			conn.socketConnection.WS.Close()
			return
		}

		if msgType == websocket.TextMessage {
			msgStr := string(msg)
			go conn.OnMessage(&msgStr)
			continue
		}
	}
}
