package chat

import "socialNetwork/entity"


func (r *chat) GetChats(id int) ([]entity.Chat, error) {
	query := `SELECT * FROM chats WHERE user_id = $1 OR friend_id = $1`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var chats []entity.Chat
	for rows.Next() {
		var chat entity.Chat
		if err := rows.Scan(&chat.ID, &chat.UserID, &chat.FriendID, &chat.Message, &chat.CreatedAt); err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return chats, nil
}

func (r *chat) GetChatById(id int) (entity.Chat, error) {
	query := `SELECT * FROM chats WHERE id = $1`
	row := r.db.QueryRow(query, id)
	var chat entity.Chat
	if err := row.Scan(&chat.ID, &chat.UserID, &chat.FriendID, &chat.Message, &chat.CreatedAt); err != nil {
		return chat, err
	}
	return chat, nil
}

func (r *chat) CreateChat(chat entity.Chat) (int, error) {
	query := `INSERT INTO chats (user_id, friend_id, message) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(query, chat.UserID, chat.FriendID, chat.Message).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *chat) GetChatMessages(chatID int) ([]entity.Message, error) {
	query := `SELECT * FROM chat_messages WHERE chat_id = $1`
	rows, err := r.db.Query(query, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var messages []entity.Message
	for rows.Next() {
		var message entity.Message
		if err := rows.Scan(&message.ID, &message.ChatID, &message.SenderID, &message.Content, &message.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *chat) CreateMessage(message entity.Message) (int, error) {
	query := `INSERT INTO chat_messages (chat_id, sender_id, content) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(query, message.ChatID, message.SenderID, message.Content).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *chat) SaveMessage(message entity.Message) error {
	query := `INSERT INTO chat_messages (chat_id, sender_id, content) VALUES ($1, $2, $3)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(message.ChatID, message.SenderID, message.Content)
	if err != nil {
		return err
	}
	return nil
}