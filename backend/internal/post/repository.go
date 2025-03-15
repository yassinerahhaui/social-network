package post

import (
	"context"
	"errors"

	"socialNetwork/entity"
)

func (p *post) Repo_UserCanPost(ctx context.Context, id, postid int) int {
	query := `SELECT
		    post.id
		FROM
			posts AS post
		LEFT JOIN follows AS follow ON follow.followed_id = post.user_id
			AND follow.follower_id = $1
		LEFT JOIN group_members AS gm ON gm.group_id = post.group_id
			AND gm.member_id = $1
		WHERE
		    (
		        post.group_id IS NOT NULL
		        AND gm.member_id IS NOT NULL
		        AND post.id = $2
		    )
		    OR (
		        post.group_id IS NULL
		        AND follow.follower_id = $1
		        AND post.id = $2
		    );`
	smtp, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		return 500
	}
	row := smtp.QueryRowContext(ctx)
	var res int
	err = row.Scan(&res)
	if err != nil {
		return 500
	} else {
		return 200
	}
}

func (p *post) Repo_GetAll(ctx context.Context, id int) (posts []entity.Post, err error) {
	prep, err := p.db.PrepareContext(ctx, `SELECT
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
		    	(post.group_id IS NULL AND follow.follower_id = $1)
			OR
				(post.status = 2);`)
	if err != nil {
		return
	}
	res, err := prep.QueryContext(ctx, id)
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
	return
}

func (p *post) Repo_GetOne(ctx context.Context, user_id, post_id int) (post entity.Post, err error) {
	prep, err := p.db.PrepareContext(ctx, `SELECT
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
	res := prep.QueryRowContext(ctx, user_id, post_id)
	err = res.Scan(&post.ID, &post.Title, &post.Content, &post.Image, &post.Status, &post.GroupID, &post.UserName)
	return
}

func (p *post) Repo_CreatePost(ctx context.Context, user_id int, post entity.Post) (err error) {
	prep, err := p.db.PrepareContext(ctx, `INSERT into posts
		SELECT $1, $2, $3, $4, $5, $6
		FROM users AS user
		LEFT JOIN group_members AS gm ON gm.group_id = $5 AND gm.member_id = $1
		WHERE user.id = $1 AND $2 NOT NULL AND $3 NOT NULL AND status NOT NULL`)
	if err != nil {
		return
	}
	res, err := prep.Exec(user_id, post.Title, post.Content, post.Image, post.GroupID, post.Status)
	if err != nil {
		return
	}
	last, err := res.LastInsertId()
	if last == 0 && err == nil {
		err = errors.New("insert error")
	}
	return
}

func (p *post) Repo_React(ctx context.Context, user_id int, react entity.PostReaction) (err error) {
	prep, err := p.db.PrepareContext(ctx, `SELECT $1 as user_id, post.id, $2 as status
		FROM posts AS post
		LEFT JOIN follows AS follow 
		  ON follow.followed_id = post.user_id AND follow.follower_id = $1
		LEFT JOIN group_members AS gm 
		  ON gm.group_id = post.group_id AND gm.member_id = $1
		WHERE 
		  	((post.group_id IS NOT NULL AND gm.member_id IS NOT NULL)
		   		OR 
			(post.group_id IS NULL AND follow.follower_id = $1))
			AND post.id = $3
		OR
			post.status = 2`)
	if err != nil {return}
	_, err = prep.ExecContext(ctx, user_id, react.ID, react.Status)
	return
}
