package chat

import (
	"bytes"
	"context"
	"errors"
	"log"
	"socialNetwork/entity"
	ws "socialNetwork/pkg/websocket"

	"github.com/gorilla/websocket"
)

func WsListing(conn *websocket.Conn, ctx context.Context) error {
	// defer conn.Close() i don't see why closing the connection here ?
	// and also closing it in the parent function
	// so for now i removed it from here ?

	m := ws.NewManager()
	val := ctx.Value(entity.ContextID)
	idInt, ok := val.(int)
	id := uint(idInt) // convert after
	if !ok {
		// handle the error: not found or wrong type
		return errors.New("user ID missing or invalid in context")
	}
	m.AddClient(id, conn)
	defer m.RemoveClient(id)
	// Listen for messages from the client
	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			return err
		}
		message = bytes.TrimSpace(bytes.Replace(message, []byte("\n"), []byte(" "), -1))
		m.SendMessage(id, message)
		log.Printf("received message from Client %d: %s", id, message)
	}
}
