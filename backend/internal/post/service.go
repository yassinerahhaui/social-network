package post

import (
	"encoding/json"
	"net/http"
	"strconv"

	"socialNetwork/entity"
)

func (p *post) Service_GetAll(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(entity.ContextID).(int)
	posts, err := p.Repo_GetAll(r.Context(), id)
	if err != nil {
		return
	}
	data, err := json.Marshal(posts)
	if err != nil {
		return
	}
	w.Write(data)
}

func (p *post) Service_GetOne(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(entity.ContextID).(int)
	post_str := r.FormValue("id")
	post_id, err := strconv.Atoi(post_str)
	if err != nil {
		return
	}
	post, err := p.Repo_GetOne(r.Context(), id, post_id)
	if err != nil {
		return
	}
	data, err := json.Marshal(post)
	if err != nil {
		return
	}
	w.Write(data)
}

func (p *post) Service_CreateOne(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(entity.ContextID).(int)
	post := entity.Post{}
	p.Repo_CreatePost(r.Context(), id, post)
	json.NewDecoder(r.Body).Decode(&post)
	w.Write([]byte(`{
		result: "done"
	}`))
}

func (p *post) Service_React(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(entity.ContextID).(int)
	react := entity.PostReaction{}
	err := json.NewDecoder(r.Body).Decode(&react)
	if err != nil {
		return
	}
	err = p.Repo_React(r.Context(), id, react)
	if err != nil {
		return
	}
	w.Write([]byte(`{
		result: "done"
	}`))
}

// get
// 0 => table(group) contains user id
// 1 => table(follows) contains user id
// 2 => get direct

// react
// 0 => table(group) contains user id
// 1 => table(follows) contains user id
// if not status forbidden
