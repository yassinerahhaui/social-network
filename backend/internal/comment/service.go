package comment

import (
	"encoding/json"
	"net/http"

	"socialNetwork/entity"
)

func (c *comment) Service_GetAll(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(entity.ContextID).(int)
	var comment entity.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		return
	}
	if c.Repo_UserCanPost(r.Context(), id, comment.PostID) != 0 {
		comments, err := c.Repo_GetAll(r.Context(), comment.PostID, id)
		if err != nil {
			return
		}
		data, err := json.Marshal(&comments)
		if err != nil {
			return
		}
		w.Write(data)
	}
}

func (c *comment) Service_Create(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(entity.ContextID).(int)
	var comment entity.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		return
	}
	if c.Repo_UserCanPost(r.Context(), id, comment.PostID) != 0 {
		err := c.Repo_Create(r.Context(), comment, id)
		if err != nil {
			return
		}
		w.Write([]byte(`done`))
	}
}

func (c *comment) Service_Vote(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(entity.ContextID).(int)
	var vote entity.Reaction
	err := json.NewDecoder(r.Body).Decode(&vote)
	if err != nil {
		return
	}
	if c.Repo_UserCanPost(r.Context(), id, vote.PostID) != 0 {
		err := c.Repo_Vote(r.Context(), vote, id)
		if err != nil {
			return
		}
		w.Write([]byte(`done`))
	}
}
