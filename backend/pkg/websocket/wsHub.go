package websocket

type Hub struct {
	Ws         *WsManager
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{}
}