package model

import "github.com/gorilla/websocket"

type SocketConnection struct {
	WS         *websocket.Conn
	OutPut     chan []byte
	Authorized bool
	ClientId   string
}
