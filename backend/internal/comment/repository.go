package comment

import (
	"context"

	"socialNetwork/entity"
)

func (c *comment) Repo_UserCanPost(ctx context.Context, id, postid int) int {
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
	smtp, err := c.db.PrepareContext(ctx, query)
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

func (c *comment) Repo_GetAll(ctx context.Context, post_id, user_id int) (comments []entity.Comment, err error) {
	smtp, err := c.db.PrepareContext(ctx, `SELECT 
			id,
			userid,
			content
		FROM
			comments 
		WHERE
			post_id = $1`)
	if err != nil {
		return
	}
	rows, err := smtp.QueryContext(ctx, post_id)
	if err != nil {
		return
	}
	for rows.Next() {
		var comment entity.Comment
		ErrInsideLoop := rows.Scan(&comment.ID, &comment.UserID, &comment.Content)
		if ErrInsideLoop != nil {
			continue
		}
		comments = append(comments, comment)
	}
	return
}

func (c *comment) Repo_Create(ctx context.Context, comment entity.Comment, user_id int) (err error) {
	smtp, err := c.db.PrepareContext(ctx, `INSERT
			INTO
				comments
			(user_id, post_id, content)
				Values
			($1, $2, $3)`)
	if err != nil {
		return
	}
	_, err = smtp.ExecContext(ctx, user_id, comment.PostID, comment.Content)
	return
}

func (c *comment) Repo_Vote(ctx context.Context, vote entity.Reaction, user_id int) (err error) {
	smtp, err := c.db.PrepareContext(ctx, `INSERT
		INTO
			engagments
				(user_id, comment_id, status)
			VALUES
				($1, $2, $3)`)
	if err != nil {
		return
	}
	_, err = smtp.ExecContext(ctx, user_id, vote.CommentID, vote.Status)
	return
}
