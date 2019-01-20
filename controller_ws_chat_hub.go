// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// HubChat maintains the set of active clientChats and broadcasts messages to the
// clientChats.
type HubChat struct {
	// Registered clientChats.
	clientChats map[*ClientChat]bool

	// Inbound messages from the clientChats.
	broadcast chan []byte

	// Register requests from the clientChats.
	register chan *ClientChat

	// Unregister requests from clientChats.
	unregister chan *ClientChat
}

func newHubChat() *HubChat {
	return &HubChat{
		broadcast:   make(chan []byte),
		register:    make(chan *ClientChat),
		unregister:  make(chan *ClientChat),
		clientChats: make(map[*ClientChat]bool),
	}
}

func (h *HubChat) run() {
	for {
		select {
		case clientChat := <-h.register:
			h.clientChats[clientChat] = true
		case clientChat := <-h.unregister:
			if _, ok := h.clientChats[clientChat]; ok {
				delete(h.clientChats, clientChat)
				close(clientChat.send)
			}
		case message := <-h.broadcast:
			for clientChat := range h.clientChats {
				select {
				case clientChat.send <- message:
				default:
					close(clientChat.send)
					delete(h.clientChats, clientChat)
				}
			}
		}
	}
}
