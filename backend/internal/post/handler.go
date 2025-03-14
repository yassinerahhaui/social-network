package post

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
	"socialNetwork/pkg/websocket"
)

type Post interface {
	Post(w http.ResponseWriter, r *http.Request)
	// GetPosts(w http.ResponseWriter, r *http.Request)
	// React(w http.ResponseWriter, r *http.Request)
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

func (u *post) GetPosts(w http.ResponseWriter, r *http.Request) {
	id, valid := r.Context().Value(entity.ContextID).(int)
	if !valid {
		return
	}
	prep, err := u.db.PrepareContext(r.Context(), `SELECT
			post.id
			post.user_id
		    title,
		    content,
			post.image,
			post.group_id,
		    (SELECT nickname FROM users AS u WHERE post.user_id=u.id) AS creator
		FROM
		    posts AS post
		LEFT JOIN groups AS "group" ON "group".id = post.group_id
		LEFT JOIN group_members AS gm ON gm.member_id = $1 AND gm.group_id = "group".id
		LEFT JOIN follows AS follow 
		    ON (follow.followed_id = post.user_id AND 
		        follow.follower_id = $1 AND 
		        post.status = 0 AND 
		        post.group_id IS NULL)
		WHERE
		    (post.group_id IS NOT NULL AND gm.member_id IS NOT NULL)
		    OR 
		    (post.group_id IS NULL AND follow.follower_id = $1);`)
	if err != nil {
		return
	}
	posts := []entity.Post{}
	res, err := prep.QueryContext(r.Context(), id)
	if err != nil {
		return
	}
	for res.Next() {
		post := entity.Post{}
		err := res.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Image, &post.GroupID, post.UserName)
		if err != nil {
			continue
		}
		posts = append(posts, post)
	}
	data, err := json.Marshal(posts)
	if err != nil {
		return
	}
	w.Write(data)
}

/*             NextJs                */
func (u *post) React(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(entity.ContextID).(int)
	prep, err := u.db.PrepareContext(r.Context(), `SELECT $1 as user_id, post.id, $2 as status
		FROM posts AS post
		LEFT JOIN follows AS follow 
		  ON follow.followed_id = post.user_id AND follow.follower_id = $1
		LEFT JOIN group_members AS gm 
		  ON gm.group_id = post.group_id AND gm.member_id = $1
		WHERE 
		  ((post.group_id IS NOT NULL AND gm.member_id IS NOT NULL)
		    OR 
		  (post.group_id IS NULL AND follow.follower_id = $1))
		    AND post.id = $3`)
	if err != nil {
		return
	}
	react := entity.PostReaction{}
	err = json.NewDecoder(r.Body).Decode(&react)
	if err != nil {
		return
	}
	_, err = prep.Exec(id, react.ID, react.Status)
	if err != nil {
		return
	}
	w.Write([]byte(`{
		result: "done"
	}`))
}

func (u *post) Post(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		u.CreatePost(w, r)
	case "GET":
		u.GetPost(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (u *post) CreatePost(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(entity.ContextID).(int)
	prep, err := u.db.PrepareContext(r.Context(), `INSERT into posts
		SELECT $1, $2, $3, $4, $5, $6
		FROM users AS user
		LEFT JOIN group_members AS gm ON gm.group_id = $5 AND gm.member_id = $1
		WHERE user.id = $1 AND $2 NOT NULL AND $3 NOT NULL AND status NOT NULL`)
	if err != nil {
		return
	}
	post := entity.Post{}
	json.NewDecoder(r.Body).Decode(&post)
	_, err = prep.Exec(id, post.Title, post.Content, post.Image, post.GroupID, post.Status)
	if err != nil {
		return
	}
	w.Write([]byte(`{
		result: "done"
	}`))
}

func (u *post) GetPost(w http.ResponseWriter, r *http.Request) {
	prep, err := u.db.PrepareContext(r.Context(), `SELECT
		    post.id,
		    post.title,
		    post.content,
			post.image,
			post.status,
			post.group_id,
		    user.nickname AS creator
		FROM posts AS post 
		INNER JOIN users AS user ON user.id = post.user_id
		LEFT JOIN group_members AS gm ON gm.group_id=post.group_id AND gm.member_id = $1
		LEFT JOIN follows AS follow ON follow.followed_id = post.user_id AND follow.follower_id = 1
		WHERE
		    (post.group_id IS NOT NULL AND gm.member_id = $1 AND post.id=$2)
		    OR 
		    (post.group_id IS NULL AND follow.follower_id = $1 AND post.id=$2);`)
	if err != nil {
		return
	}
	post := entity.Post{}
	res := prep.QueryRowContext(r.Context(), post.ID)
	err = res.Scan(&post.ID, &post.Title, &post.Content, &post.Image, &post.Status, &post.GroupID, post.UserName)
	if err != nil {
		return
	}
	data, err := json.Marshal(post)
	if err != nil {
		return
	}
	w.Write(data)
}
