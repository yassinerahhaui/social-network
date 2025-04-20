package websocket

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

type WsManager struct {
	Clients map[uint]*Client
	Mtx     sync.Mutex
}

func NewManager() *WsManager {
	return &WsManager{
		Clients: make(map[uint]*Client),
		Mtx:     sync.Mutex{},
	}
}

func (m *WsManager) AddClient(userID uint, conn *websocket.Conn) {
	m.Mtx.Lock()
	defer m.Mtx.Unlock()
	c := &Client{
		conn: conn,
		send: make(chan []byte, 1024),
	}
	m.Clients[userID] = c
	log.Printf("Client %d added", userID)
	go m.readMessage(userID)
	go m.writeMessage(userID)
}

func (m *WsManager) RemoveClient(userID uint) {
	m.Mtx.Lock()
	defer m.Mtx.Unlock()

	if c, ok := m.Clients[userID]; ok {
		close(c.send)
		delete(m.Clients, userID)
		log.Printf("Client %d removed", userID)
	}
}

func (m *WsManager) readMessage(userID uint) {
	c := m.Clients[userID]
	defer m.RemoveClient(userID)

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message from Client %d: %v", userID, err)
			break
		}
		log.Printf("Received message from Client %d: %s", userID, string(message))

		// Handle the message (e.g., broadcast to other Clients or process it)
		m.Broadcast(userID, message)
	}
}

func (m *WsManager) writeMessage(userID uint) {
	c := m.Clients[userID]
	defer m.RemoveClient(userID)
	fmt.Println("clients:", c)
	for message := range c.send { //
		fmt.Println("message:", message)
		err := c.conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Error writing message to Client %d: %v", userID, err)
			break
		}
	}
	// for id, c := range m.Clients {
	// 	// Skip sending a message to the user that triggered the event
	// 	if id == userID {
	// 		continue
	// 	}

	// 	for message := range c.send {
	// 		if c.conn == nil {
	// 			log.Printf("Client %d's connection is nil", id)
	// 			continue
	// 		}

	// 		err := c.conn.WriteMessage(websocket.TextMessage, message)
	// 		if err != nil {
	// 			log.Printf("Error writing message to Client %d: %v", id, err)
	// 		}
	// 	}
	// }
}

func (m *WsManager) SendMessage(userID uint, message []byte) {
	m.Mtx.Lock()
	defer m.Mtx.Unlock()

	if c, ok := m.Clients[userID]; ok {
		select {
		case c.send <- message:
		default:
			log.Printf("Client %d send buffer is full", userID)
		}
	}
}

func (m *WsManager) Broadcast(userID uint, message []byte) {
	m.Mtx.Lock()
	defer m.Mtx.Unlock()
	for id, c := range m.Clients {
		if id != userID {
			select {
			case c.send <- message:
			default:
				log.Printf("Client %d send buffer is full", id)
			}
		}
	}
}

func (m *WsManager) GetClients() map[uint]*Client {
	return m.Clients
}

func (m *WsManager) GetClient(userID uint) *Client {
	return m.Clients[userID]
}
