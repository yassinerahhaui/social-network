package post

import (
	"encoding/json"
	"net/http"
	"runtime"
	"strconv"

	"socialNetwork/entity"
)

func (p *post) Service_GetAll(w http.ResponseWriter, r *http.Request) {
	id := 1 // r.Context().Value(entity.ContextID).(int)
	posts, err := p.Repo_GetAll(r.Context(), id)
	if err != nil {
		w.Write([]byte(GetError(err)))
		return
	}
	data, err := json.Marshal(posts)
	if err != nil {
		w.Write([]byte(GetError(err)))
		return
	}
	w.Write(data)
}

func (p *post) Service_GetOne(w http.ResponseWriter, r *http.Request) {
	id := 1 // r.Context().Value(entity.ContextID).(int)
	post_str := r.PathValue("id")
	post_id, err := strconv.Atoi(post_str)
	if err != nil {
		w.Write([]byte("id : " + post_str))
		w.Write([]byte(GetError(err)))
		return
	}
	if p.Repo_UserCanPost(r.Context(), id, post_id) {
		post, err := p.Repo_GetOne(r.Context(), post_id)
		if err != nil {
			w.Write([]byte(GetError(err)))
			return
		}
		data, err := json.Marshal(post)
		if err != nil {
			w.Write([]byte(GetError(err)))
			return
		}
		w.Write(data)
	} else {
		w.Write([]byte("{ error: 'post not found' }"))
		return
	}
}

func (p *post) Service_CreateOne(w http.ResponseWriter, r *http.Request) {
	id := 1 // r.Context().Value(entity.ContextID).(int)
	post := entity.Post{}
	json.NewDecoder(r.Body).Decode(&post)
	p.Repo_CreatePost(r.Context(), id, post)
	w.Write([]byte(`{
		result: "done"
	}`))
}

func (p *post) Service_React(w http.ResponseWriter, r *http.Request) {
	id := 1 // r.Context().Value(entity.ContextID).(int)
	react := entity.Reaction{}
	json.NewDecoder(r.Body).Decode(&react)
	if p.Repo_UserCanPost(r.Context(), id, react.PostID) {
		err := p.Repo_React(r.Context(), id, react)
		if err != nil {
			w.Write([]byte(GetError(err)))
			return
		}
		w.Write([]byte(`{
			result: "done"
		}`))
	} else {
		w.Write([]byte(" { error : 'you can't see the post' } "))
	}
}

func GetError(err error) string {
	_, _, line, _ := runtime.Caller(1)
	ln := strconv.Itoa(line)
	return "line > " + ln + " := " + err.Error()
}

// get
// 0 => table(group) contains user id
// 1 => table(follows) contains user id
// 2 => get direct

// react
// 0 => table(group) contains user id
// 1 => table(follows) contains user id
// if not status forbidden
