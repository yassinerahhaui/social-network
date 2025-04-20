CREATE TABLE IF NOT EXISTS notification (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    type INTEGER NOT NULL,
    group_id INTEGER,
    sender_id INTEGER,
    receiver_id INTEGER,
    event_id INTEGER,
    accepted BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
    FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (event_id) REFERENCES events(id) ON DELETE CASCADE, 
    FOREIGN KEY (receiver_id) REFERENCES users(id) ON DELETE CASCADE
);