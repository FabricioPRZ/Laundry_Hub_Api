package websocket

import "sync"

type Hub struct {
	clients    map[int]*Client
	register   chan *Client
	unregister chan *Client
	broadcast  chan *Message
	mu         sync.RWMutex
}

type Message struct {
	UserID  int    `json:"user_id"`
	Payload []byte `json:"payload"`
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[int]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client.UserID] = client
			h.mu.Unlock()

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client.UserID]; ok {
				delete(h.clients, client.UserID)
				close(client.send)
			}
			h.mu.Unlock()

		case message := <-h.broadcast:
			h.mu.RLock()
			if client, ok := h.clients[message.UserID]; ok {
				select {
				case client.send <- message.Payload:
				default:
					close(client.send)
					delete(h.clients, client.UserID)
				}
			}
			h.mu.RUnlock()
		}
	}
}
