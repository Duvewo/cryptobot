package cryptocurrency

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
}

func (c Client) Write(data []byte) error {
	return c.conn.WriteJSON(data)
}

func (c Client) Read(v *interface{}) error {
	return c.conn.ReadJSON(&v)
}
