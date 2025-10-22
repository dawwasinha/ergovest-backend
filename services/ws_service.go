package services

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Simple WebSocket hub for broadcasting SensorData and Alerts
type wsClient struct {
	conn *websocket.Conn
}

type WSHub struct {
	clients map[*wsClient]struct{}
	mu      sync.RWMutex
}

var hub = &WSHub{
	clients: make(map[*wsClient]struct{}),
}

// Register a new client
func RegisterWSClient(conn *websocket.Conn) *wsClient {
	c := &wsClient{conn: conn}
	hub.mu.Lock()
	hub.clients[c] = struct{}{}
	hub.mu.Unlock()
	return c
}

// Unregister client and close connection
func UnregisterWSClient(c *wsClient) {
	hub.mu.Lock()
	delete(hub.clients, c)
	hub.mu.Unlock()
	_ = c.conn.Close()
}

// Broadcast a raw JSON message to all connected clients
func BroadcastRaw(messageType int, data []byte) {
	hub.mu.RLock()
	clients := make([]*wsClient, 0, len(hub.clients))
	for c := range hub.clients {
		clients = append(clients, c)
	}
	hub.mu.RUnlock()

	for _, c := range clients {
		_ = c.conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
		if err := c.conn.WriteMessage(messageType, data); err != nil {
			// Unregister failing client
			UnregisterWSClient(c)
		}
	}
}
