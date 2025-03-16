package entity

type Reaction struct {
	CommentID int    `json:"comment"`
	PostID    int    `json:"post"`
	Status    string `json:"status"`
}
