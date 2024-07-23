-- Users Table
CREATE TABLE IF NOT EXISTS users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE
);

-- Games Table
CREATE TABLE IF NOT EXISTS games (
    game_id INTEGER PRIMARY KEY AUTOINCREMENT,
    player1_id INTEGER NOT NULL,
    player2_id INTEGER NOT NULL, -- For now a second player is required
    board_state TEXT NOT NULL,
    turn INTEGER NOT NULL CHECK("turn" IN (1, 2)),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (player1_id) REFERENCES users (user_id),
    FOREIGN KEY (player2_id) REFERENCES users (user_id)
);

-- Moves Table
CREATE TABLE IF NOT EXISTS moves (
    move_id INTEGER PRIMARY KEY AUTOINCREMENT,
    game_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    move_data TEXT NOT NULL,
    move_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (game_id) REFERENCES games (game_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

