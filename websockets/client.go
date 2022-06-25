package websockets

import "github.com/gorilla/websocket"

type Client struct {
	hub      *Hub
	id       string
	socket   *websocket.Conn
	outbound chan []byte
}

func (c *Client) Write() {
	for {
		select {
		case message, ok := <-c.outbound:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
