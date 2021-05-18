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
