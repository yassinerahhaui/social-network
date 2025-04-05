package server

import (
	"time"

	"socialNetwork/internal/chat"
	"socialNetwork/internal/comment"
	"socialNetwork/internal/group"
	"socialNetwork/internal/post"
	"socialNetwork/internal/user"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/database"
	"socialNetwork/pkg/loger"
	scs "socialNetwork/pkg/sessions"
	"socialNetwork/pkg/websocket"
)

type App struct {
	SessionManager *scs.SessionManager
	comment.Comment
	chat.Chat
	group.Group
	user.User
	post.Post
	Loger *loger.CstmLogger
}

func NewApp(dep *config.Dependencies) *App {
	return &App{
		SessionManager: dep.SessionManager,
		Loger:          dep.Loger,
		Comment:        comment.NewComment(dep),
		Chat:           chat.NewChat(dep),
		Group:          group.NewGroup(dep /* we need to add the hub to group*/),
		User:           user.NewUser(dep /* we need to add the hub*/),
		Post:           post.Newpost(dep),
	}
}

func NewTestApplication() (*App, *config.Conf) {
	loger := loger.NewLogger()
	// And a form decoder.
	// formDecoder := form.NewDecoder()
	// And a session manager instance. Note that we use the same settings as
	// production, except that we *don't* set a Store for the session manager.
	// If no store is set, the SCS package will default to using a transient
	// in-memory store, which is ideal for testing purposes.
	sessionManager := scs.New()
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Cookie.Secure = true

	cfg, err := config.NewConfig()
	if err != nil {
		loger.Error.Panicln("error while reading config file\n", err)
	}
	db, err := database.InitDB(cfg.TestDatabase)
	if err != nil {
		loger.Error.Panicf("error occured while connecting database: %s", err.Error())
	}
	// defer db.Close()

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
	}, cfg
}
