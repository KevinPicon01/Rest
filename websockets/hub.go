package websockets

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/segmentio/ksuid"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Hub struct {
	clients    []*Client
	register   chan *Client
	unregister chan *Client
	mutex      *sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make([]*Client, 0),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		mutex:      &sync.Mutex{},
	}
}
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.onConnect(client)
		case client := <-h.unregister:
			h.onDisconnect(client)
		}
	}
}

func (h *Hub) onConnect(client *Client) {
	fmt.Println("Client connected:", client.id)
	client.outbound <- []byte("Welcome to the chat!")
	h.mutex.Lock()
	defer h.mutex.Unlock()
	client.id = client.socket.RemoteAddr().String()
	h.clients = append(h.clients, client)
}
func (h *Hub) onDisconnect(client *Client) {
	fmt.Println("Client disconnected:", client.id)
	client.outbound <- []byte("Goodbye!")
	client.socket.Close()
	h.mutex.Lock()
	defer h.mutex.Unlock()
	for i, c := range h.clients {
		if c.id == client.id {
			h.clients = append(h.clients[:i], h.clients[i+1:]...)
			break
		}
	}
}

func (hub *Hub) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	client := &Client{
		hub:      hub,
		id:       ksuid.New().String(),
		socket:   conn,
		outbound: make(chan []byte),
	}
	hub.register <- client
	defer func() {
		hub.unregister <- client
		conn.Close()
	}()
	go client.Write()
}

func (hub *Hub) Broadcast(message interface{}, ignore *Client) {
	data, _ := json.Marshal(message)
	for _, client := range hub.clients {
		if client != ignore {
			client.outbound <- data
		}
	}
}
