package main

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type HubStock struct {
	// Registered clients.
	clients map[*ClientStock]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *ClientStock

	// Unregister requests from clients.
	unregister chan *ClientStock
}

func newHubStock() *HubStock {
	return &HubStock{
		broadcast:  make(chan []byte),
		register:   make(chan *ClientStock),
		unregister: make(chan *ClientStock),
		clients:    make(map[*ClientStock]bool),
	}
}

func (h *HubStock) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
