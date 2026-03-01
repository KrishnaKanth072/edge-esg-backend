package handlers

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for now
	},
}

type WSMessage struct {
	Type    string                 `json:"type"`
	Data    map[string]interface{} `json:"data"`
	AgentID string                 `json:"agentId,omitempty"`
}

type WSHub struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan WSMessage
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	mu         sync.RWMutex
}

func NewWSHub() *WSHub {
	return &WSHub{
		clients:    make(map[*websocket.Conn]bool),
		broadcast:  make(chan WSMessage, 256),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
	}
}

func (h *WSHub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				client.Close()
			}
			h.mu.Unlock()

		case message := <-h.broadcast:
			h.mu.RLock()
			for client := range h.clients {
				err := client.WriteJSON(message)
				if err != nil {
					client.Close()
					delete(h.clients, client)
				}
			}
			h.mu.RUnlock()
		}
	}
}

func (h *WSHub) HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	h.register <- conn

	defer func() {
		h.unregister <- conn
	}()

	for {
		var msg WSMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			break
		}
		// Echo back or process message
		h.broadcast <- msg
	}
}

func (h *WSHub) BroadcastAgentUpdate(agentID string, status string, progress int) {
	h.broadcast <- WSMessage{
		Type:    "agent_update",
		AgentID: agentID,
		Data: map[string]interface{}{
			"status":   status,
			"progress": progress,
		},
	}
}

func (h *WSHub) BroadcastConsensus(consensus int) {
	h.broadcast <- WSMessage{
		Type: "consensus_update",
		Data: map[string]interface{}{
			"consensus": consensus,
		},
	}
}
