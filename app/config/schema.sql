CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    username TEXT NOT NULL UNIQUE CHECK(length(username) >= 3 AND length(username) <= 50),
    email TEXT NOT NULL UNIQUE CHECK(length(email) <= 255),
    password TEXT NOT NULL,
    token TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    item_id TEXT UNIQUE, -- Formatted ID, e.g. "posts-1"
    user_id TEXT ,
    title TEXT NOT NULL,
    content TEXT NOT NULL CHECK(length(content) <= 10000),
    categories TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    item_id TEXT UNIQUE, -- Formatted ID, e.g. "comments-1"
    post_id INTEGER NOT NULL,
    user_id TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS reactions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT NOT NULL,
    item_id TEXT NOT NULL, -- Refer to either post or comment
    reaction_type TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    UNIQUE (user_id, item_id)
);

CREATE TABLE IF NOT EXISTS reactionCount (
    item_id TEXT NOT NULL, -- Unified item ID for posts and comments
    reaction_type TEXT NOT NULL, -- e.g. 'like', 'dislike'
    count INTEGER DEFAULT 0,
    UNIQUE (item_id, reaction_type)
);

CREATE TRIGGER IF NOT EXISTS set_posts_id
AFTER INSERT ON posts
BEGIN
    UPDATE posts
    SET item_id = 'posts-' || NEW.id
    WHERE id = NEW.id;
END;

CREATE TRIGGER IF NOT EXISTS set_comments_id
AFTER INSERT ON comments
BEGIN
    UPDATE comments
    SET item_id = 'comments-' || NEW.id
    WHERE id = NEW.id;
END;

CREATE TRIGGER IF NOT EXISTS increment_reaction_count
AFTER INSERT ON reactions
BEGIN
    INSERT INTO reactionCount (item_id, reaction_type, count)
    VALUES (NEW.item_id, NEW.reaction_type, 1)
    ON CONFLICT (item_id, reaction_type) 
    DO UPDATE SET count = count + 1;
END;

CREATE TRIGGER IF NOT EXISTS decrement_reaction_count
AFTER DELETE ON reactions
BEGIN
    UPDATE reactionCount
    SET count = count - 1
    WHERE item_id = OLD.item_id AND reaction_type = OLD.reaction_type;
    DELETE FROM reactionCount
    WHERE count = 0;
END;

CREATE TRIGGER IF NOT EXISTS update_reaction_count
AFTER UPDATE ON reactions
BEGIN
    UPDATE reactionCount
    SET count = count - 1
    WHERE item_id = OLD.item_id AND reaction_type = OLD.reaction_type;
    
    DELETE FROM reactionCount
    WHERE count = 0;

    INSERT INTO reactionCount (item_id, reaction_type, count)
    VALUES (NEW.item_id, NEW.reaction_type, 1)
    ON CONFLICT (item_id, reaction_type) 
    DO UPDATE SET count = count + 1;
END;