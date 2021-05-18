package cryptocurrency

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
}

type Request struct {
	Event   string `json:"_event"`
	TzID    int    `json:"tzID,omitempty"`
	Message string `json:"message,omitempty"`
	Data    string `json:"data,omitempty"`
	UID     int    `json:"UID,omitempty"`
}

type Response struct {
	Message string `json:"message"`
}
