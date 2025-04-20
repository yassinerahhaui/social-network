package server

import (
	"net/http"
	"socialNetwork/internal/chat"
	"socialNetwork/internal/comment"
	"socialNetwork/internal/group"
	"socialNetwork/internal/post"
	"socialNetwork/internal/user"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/database"
	"socialNetwork/pkg/loger"
	scs "socialNetwork/pkg/sessions"
	"socialNetwork/pkg/store"
	"socialNetwork/pkg/websocket"
	"time"
)

// handler used  for status-checking or uptime monitoring of your server.
func ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("OK"))
}

func NewTestApplication() (*App, *config.Conf) {
	loger := loger.NewTestLogger()
	//loger := loger.NewLogger()

	cfg, err := config.NewConfig()
	if err != nil {
		loger.Error.Panicln("error while reading config file\n", err)
	}
	db, err := database.InitDB(cfg.TestDatabase)
	if err != nil {
		loger.Error.Panicf("error occured while connecting database: %s", err.Error())
	}
	// defer db.Close()

	// And a form decoder.
	// formDecoder := form.NewDecoder()
	// And a session manager instance. Note that we use the same settings as
	// production, except that we *don't* set a Store for the session manager.
	// If no store is set, the SCS package will default to using a transient
	// in-memory store, which is ideal for testing purposes.
	sessionManager := scs.New()
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Store = store.New(db)

	dep := &config.Dependencies{
		SessionManager: sessionManager,
		DB:             db,
		Loger:          loger,
		/*Legislation is the process or result of enrolling, enacting, or promulgating laws by a legislature, parliament, or analogous governing body.*/
		//hub := websocket.NewHub() // need console legislation
		Hub: websocket.NewHub(),
	}

	return &App{
		SessionManager: sessionManager,
		Loger:          loger,
		User:           user.NewUser(dep /* we need to add the hub*/),
		Comment:        comment.NewComment(dep),
		Chat:           chat.NewChat(dep),
		Group:          group.NewGroup(dep /* we need to add the hub to group*/),
		Post:           post.Newpost(dep),
	}, cfg
}
