package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"socialNetwork/entity"
	"socialNetwork/pkg/config"

	"golang.org/x/crypto/bcrypt"
)

/*___________ THOS FUNCTIONS FOR AUTHENTICATION ___________*/

func (r *user) GetUserProfileById(ctx context.Context, targetId int) (entity.User, error) {
	var user entity.User
	requesterId := ctx.Value(entity.ContextID)
	query := `
		SELECT 
			u.id, u.nickname, u.email, u.password, u.avatar, 
			u.first_name, u.last_name, u.birthday, u.about_me, 
			u.status,
			(SELECT COUNT(*) FROM follows WHERE followed_id = u.id) AS followers_count,
			(SELECT COUNT(*) FROM follows WHERE follower_id = u.id) AS following_count,
			CASE 
				WHEN EXISTS (SELECT 1 FROM follows WHERE follower_id = $2 AND followed_id = $1) THEN 2
				WHEN EXISTS (SELECT 1 FROM follows WHERE follower_id = $1 AND followed_id = $2) THEN 1
				ELSE 0 
			END AS following_state
		FROM users u
		WHERE u.id = $1
	`

	err := r.db.QueryRowContext(ctx, query, requesterId, targetId).Scan(
		&user.ID, &user.Nickname, &user.Email, &user.Password, &user.Avatar,
		&user.First, &user.Last, &user.DateOfBirth, &user.AboutMe, &user.Status,
		&user.FollowersCount, &user.FollowingCount, &user.FollowingState,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user not found")
		}
		return user, err
	}

	return user, nil
}


// this function is used to get user by username
func (r *user) GetUserByUsername(username string) (entity.User, error) {
	// SQL query that includes the counts and following state
	query := `SELECT * FROM users WHERE nickname = $1 OR email = $1`

	var user entity.User
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return user, err
	}
	err = stmt.QueryRow(username).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.First,
		&user.Last,
		&user.DateOfBirth,
		&user.Avatar,
		&user.Nickname,
		&user.AboutMe,
		&user.Status,
		&user.FollowersCount,
		&user.FollowingCount,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			return user, err
		}
	}
	return user, nil
}

// this function is used to create new user
func (u *user) CreateUser(user entity.User) error {
	// ok, err := u.IsExistsService(int(user.ID))
	ok, err := u.CheckUserByUsername(user.Nickname)
	if err != nil {
		return err
	}
	if ok {
		return config.ErrUserAlreadyExists
	}
	query := `INSERT INTO users(
						Email,
						Password,
						first_name,
						last_name,
						birthday,
						Nickname,
						about_me,
						Avatar)
						VALUES($1, $2, $3, $4, $5, $6, $7, $8)`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		user.Email,
		user.Password,
		user.First,
		user.Last,
		user.DateOfBirth,
		user.Nickname,
		user.AboutMe,
		user.Avatar)
	if err != nil {
		return err
	}
	return nil
}

// this function is used to update user
func (r *user) UpdateUser(user entity.User) error {
	query := `UPDATE users SET
						Email = $1,
						Password = $2,
						first_name = $3,
						last_name = $4,
						birthday = $5,
						nick_name = $6,
						about_me = $7
						WHERE id = $8`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		user.Email,
		user.Password,
		user.First,
		user.Last,
		user.DateOfBirth,
		user.Nickname,
		user.AboutMe)
	if err != nil {
		return err
	}
	return nil
}

// this function is used to delete user by id
func (r *user) DeleteUser(id uint) error {
	query := `DELETE FROM users WHERE id = $1`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// this function is used to delete user by nickname
func (u *user) DeleteUserByNickName(nickName string) error {
	ok, err := u.CheckUserByUsername(nickName)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	query := `DELETE FROM users WHERE Nickname = $1`
	stmt, err := u.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(nickName)
	if err != nil {
		return err
	}
	return nil
}

/* ___________ THOSE FUNCS USED FOR CHECK USER CREDENTIALS ___________ */

// We'll use the Exists method to check if a user exists with a specific ID.
func (u *user) IsUserExist(id uint) (bool, error) {
	var exists bool
	stmt := "SELECT EXISTS(SELECT true FROM users WHERE id = ?)"
	err := u.db.QueryRow(stmt, id).Scan(&exists)
	return exists, err
}

// We'll use the Authenticate method to verify whether a user exists with
// the provided email address and password. This will return the relevant
// user ID if they do.
func (u *user) authenticateRepo(email, password string) (int, error) {
	// u.loger.Info.Println("email:", email)
	// u.loger.Info.Println("password:", password)
	// Retrieve the id and hashed password associated with the given email. If
	// no matching email exists we return the ErrInvalidCredentials error.
	var id int
	var hashedPassword []byte
	stmt := "SELECT id, password FROM users WHERE email = ? OR nickname  = ?"
	err := u.db.QueryRow(stmt, email, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, config.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	// Check whether the hashed password and plain-text password provided match.
	// If they don't, we return the ErrInvalidCredentials error.
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, config.ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	// Otherwise, the password is correct. Return the user ID.
	return id, nil
}

// this function checks by email if the user exists
func (r *user) CheckUserByEmail(email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`
	var exists bool
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return exists, err
	}
	err = stmt.QueryRow(email).Scan(&exists)
	if err != nil {
		return exists, err
	}
	return exists, nil
}

// this function checks by username if the user exists
func (r *user) CheckUserByUsername(username string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE Nickname = $1)`
	var exists bool
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return exists, err
	}
	err = stmt.QueryRow(username).Scan(&exists)
	if err != nil {
		return exists, fmt.Errorf("error executing query: %w", err)
	}
	return exists, nil
}

/*________________ THOSE FUNCS USED TO CHECK FOLOWING ____________ */
func (u *user) isFollowedBy(follower, followed int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM follows WHERE follower_id = $1 AND followed_id = $2)`
	var exists bool
	stmt, err := u.db.Prepare(query)
	if err != nil {
		return exists, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(follower, followed).Scan(&exists)
	if err != nil {
		return exists, fmt.Errorf("error executing query: %w", err)
	}

	return exists, nil
}

func (u *user) isFollowingEither(follower, followed int) (bool, error) {
	query := `SELECT EXISTS(
                SELECT 1 FROM follows WHERE (follower_id = $1 AND followed_id = $2) 
                OR (follower_id = $2 AND followed_id = $1)
              )`

	stmt, err := u.db.Prepare(query)
	if err != nil {
		return false, fmt.Errorf("error preparing query: %w", err)
	}
	defer stmt.Close()

	var exists bool
	err = stmt.QueryRow(follower, followed).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error executing query: %w", err)
	}

	return exists, nil
}

func (u *user) FollowRepository(followerId, followedId int) error {
	// Check if the follow relationship already exists
	queryCheck := `SELECT 1 FROM follows WHERE follower_id = $1 AND followed_id = $2`
	var exists int
	err := u.db.QueryRow(queryCheck, followerId, followedId).Scan(&exists)

	if err == nil { // Row exists → Unfollow (delete)
		queryDelete := `DELETE FROM follows WHERE follower_id = $1 AND followed_id = $2`
		_, err = u.db.Exec(queryDelete, followerId, followedId)
		if err != nil {
			return fmt.Errorf("failed to unfollow: %w", err)
		}
		return nil
	} else if err != sql.ErrNoRows { // Any other error (DB issue)
		return fmt.Errorf("database error: %w", err)
	}

	// Row does not exist → Follow (insert)
	queryInsert := `INSERT INTO follows (follower_id, followed_id) VALUES ($1, $2)`
	_, err = u.db.Exec(queryInsert, followerId, followedId)
	if err != nil {
		return fmt.Errorf("failed to follow: %w", err)
	}
	return nil
}

// need some changes to follow up with the macro image
func (u *user) GroupContainsMember(groupId, userId int) (bool, error) {
	query := `SELECT 1 FROM members WHERE group_id = ? AND user_id = ? LIMIT 1`
	var exists int

	err := u.db.QueryRow(query, groupId, userId).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("error checking group membership: %v", err)
	}

	return true, nil
}

/*___________ THOS FUNC USED FOR FOLLOWERS ___________*/
// this function is used to get followers by user id
func (u *user) GetFollowers(ctx context.Context, id int) ([]entity.User, error) {
	query := `
		SELECT users.nickname, users.avatar 
		FROM users
		INNER JOIN follows ON users.id = follows.follower_id
		WHERE follows.followed_id = $1`
		
	rows, err := u.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []entity.User
	for rows.Next() {
		var follower entity.User
		if err := rows.Scan(&follower.Nickname, &follower.Avatar); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	return followers, nil
}


// this function is used to get following by user id
func (u *user) GetFollowing(ctx context.Context, id int) ([]entity.User, error) {
	query := `
		SELECT users.nickname, users.avatar 
		FROM users
		INNER JOIN follows ON users.id = follows.followed_id 
		WHERE follows.follower_id = $1`
		
	rows, err := u.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []entity.User
	for rows.Next() {
		var follower entity.User
		if err := rows.Scan(&follower.Nickname, &follower.Avatar); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	return followers, nil
}

// this function is used to follow user
// func (r *user) Follow(follower, following uint) error {}
// this function is used to unfollow user
// func (r *user) Unfollow(follower, following uint) error {}
// this function is used to get followers count
// func (r *user) GetFollowersCount(id uint) (uint, error) {}
// this function is used to get following count
// func (r *user) GetFollowingCount(id uint) (uint, error) {}
