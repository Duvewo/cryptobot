package cryptocurrency

import (
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
)

func Dial(u *url.URL) (*Client, error) {
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), http.Header{})

	if err != nil {
		return nil, err
	}

	return &Client{conn}, nil

}

func (c Client) WriteJSON(data interface{}) error {
	return c.conn.WriteJSON(data)
}

func (c Client) ReadJSON(v interface{}) error {
	return c.conn.ReadJSON(v)
}

func (c Client) Write(messageType int, data []byte) error {
	return c.conn.WriteMessage(messageType, data)
}

func (c Client) Read() (int, []byte, error) {
	return c.conn.ReadMessage()
}
