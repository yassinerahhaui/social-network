package chat

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
	scs "socialNetwork/pkg/sessions"
	ws "socialNetwork/pkg/websocket"

	"github.com/gorilla/websocket"
)

type chat struct {
	Hub            *ws.Hub
	db             *sql.DB
	loger          loger.CstmLogger
	sessionManager *scs.SessionManager
}

type Chat interface {
	WebSocket(w http.ResponseWriter, r *http.Request)
}

func NewChat(dep *config.Dependencies) Chat {
	return &chat{db: dep.DB, loger: *dep.Loger, Hub: dep.Hub}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (use cautiously in production)
	},
}

func (c *chat) WebSocket(w http.ResponseWriter, r *http.Request) {
	// upgrade
	c.loger.Info.Println("0")
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  2048,
		WriteBufferSize: 2048,
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins (use cautiously in production)
		},
	}
	c.loger.Info.Println("1")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		c.loger.Error.Println("Error while upgrading connection:", err)
		json.NewEncoder(w).Encode(entity.ErrorResponse{Error: "Error while upgrading connection"})
		return
	}

	// // Set ping handler
	// conn.SetPingHandler(func(appData string) error {
	// 	log.Printf("Received ping from client!")
	// 	return nil
	// })

	// // Set pong handler
	// conn.SetPongHandler(func(appData string) error {
	// 	log.Printf("Received pong from client!")
	// 	return nil
	// })
	// // Send ping every 10 seconds
	// go func() {
	// 	for {
	// 		conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(time.Second*10))
	// 	}
	// }()
	// track last pong time
	// var lastPongTime time.Time

	// conn.SetPongHandler(func(appData string) error {
	// 	lastPongTime = time.Now()
	// 	return nil
	// })

	// // optional: handle ping if client sends one
	// conn.SetPingHandler(func(appData string) error {
	// 	log.Println("Ping received")
	// 	return conn.WriteControl(websocket.PongMessage, []byte(appData), time.Now().Add(time.Second))
	// })

	// // ping client every 10s
	// go func() {
	// 	ticker := time.NewTicker(10 * time.Second)
	// 	defer ticker.Stop()

	// 	for {
	// 		<-ticker.C

	// 		// check pong timeout
	// 		if time.Since(lastPongTime) > 20*time.Second {
	// 			log.Println("No pong from client, closing connection")
	// 			conn.Close()
	// 			return
	// 		}

	// 		// send ping
	// 		conn.WriteControl(websocket.PingMessage, []byte("ping"), time.Now().Add(time.Second))
	// 	}
	// }()
	/*________________________________________________________________________________________________________*/
	c.loger.Info.Println("2")
	defer conn.Close()
	//get user id from session
	userId := r.Context().Value(entity.ContextID)
	c.loger.Info.Println("3")
	if userId == nil {
		c.loger.Error.Println("User ID not found in session")
		//w.WriteHeader(http.StatusUnauthorized)
		return
	}

	c.loger.Info.Println("4")
	c.loger.Info.Printf("Client %d connected\n", userId)
	WsListing(conn, r.Context())
	c.loger.Info.Printf("Client %d disconnected\n", userId)
}
