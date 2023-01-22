package websocket

import (
	"strings"
	"sync"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	// We need a mutex to protect the clients map
	mu sync.Mutex
}

func (h *Hub) GetClientByUid(uid string) *Client {
	h.mu.Lock()
	defer h.mu.Unlock()
	for client := range h.clients {
		if client.Uid == uid {
			return client
		}
	}
	return nil
}

// Get count of unregistered clients (guests)
func (h *Hub) GetGuestCount() int {
	h.mu.Lock()
	defer h.mu.Unlock()
	count := 0
	for client := range h.clients {
		if strings.HasPrefix("guest", client.Uid) {
			count++
		}
	}
	return count
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			func() {
				h.mu.Lock()
				defer h.mu.Unlock()
				h.clients[client] = true
			}()
		case client := <-h.Unregister:
			func() {
				h.mu.Lock()
				defer h.mu.Unlock()
				if _, ok := h.clients[client]; ok {
					delete(h.clients, client)
					close(client.Send)
				}
			}()
		case message := <-h.Broadcast:
			func() {
				h.mu.Lock()
				defer h.mu.Unlock()
				for client := range h.clients {
					select {
					case client.Send <- message:
					default:
						close(client.Send)
						delete(h.clients, client)
					}
				}
			}()
		}
	}
}