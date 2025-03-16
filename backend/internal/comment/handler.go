package comment

import (
	"database/sql"
	"net/http"

	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
	"socialNetwork/pkg/websocket"
)

type comment struct {
	Hub   *websocket.Hub
	db    *sql.DB
	loger loger.CstmLogger
}

type Comment interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Vote(w http.ResponseWriter, r *http.Request)
}

func NewComment(dep *config.Dependencies) Comment {
	return &comment{db: dep.DB, loger: *dep.Loger, Hub: dep.Hub}
}

func (c *comment) GetAll(w http.ResponseWriter, r *http.Request) {}
func (c *comment) Create(w http.ResponseWriter, r *http.Request) {}
func (c *comment) Vote(w http.ResponseWriter, r *http.Request)   {}
