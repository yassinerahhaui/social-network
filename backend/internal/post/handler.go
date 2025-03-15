package post

import (
	"database/sql"
	"net/http"

	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
	"socialNetwork/pkg/websocket"
)

type Post interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	CreatePost(w http.ResponseWriter, r *http.Request)
	GetPost(w http.ResponseWriter, r *http.Request)
	React(w http.ResponseWriter, r *http.Request)
}

type post struct {
	Hub   *websocket.Hub
	db    *sql.DB
	loger loger.CstmLogger
}

var _ string = `{
					"id": 1,
					"title": "My First Post",
					"image": "iVBORw0KGgoAAAANSUhEUgAA...kJggg==",
					"content": "This is the content of the post.",
					"user_id": 42,
					"status": 1,
					"group": 5,
					"created_at": "2025-03-11T12:34:56Z",
					"updated_at": "2025-03-11T12:34:56Z"
				}`

func Newpost(dep *config.Dependencies) Post {
	return &post{db: dep.DB, loger: *dep.Loger, Hub: dep.Hub}
}

func (p *post) GetPosts(w http.ResponseWriter, r *http.Request) {}

func (p *post) React(w http.ResponseWriter, r *http.Request) {}

func (p *post) CreatePost(w http.ResponseWriter, r *http.Request) {}

func (p *post) GetPost(w http.ResponseWriter, r *http.Request) {}
