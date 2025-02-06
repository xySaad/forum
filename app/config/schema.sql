CREATE TABLE IF NOT EXISTS users (
  id BIGINT NOT NULL UNIQUE,                     -- Public Snowflake id (exposed)
  internal_id INTEGER PRIMARY KEY AUTOINCREMENT, -- Internal (for joins)
  token TEXT NOT NULL UNIQUE,
  username TEXT NOT NULL UNIQUE,
  email TEXT NOT NULL UNIQUE,
  profile_picture TEXT,
  password TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS items (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE
);

INSERT OR IGNORE INTO items (name) VALUES ("post");
INSERT OR IGNORE INTO items (name) VALUES ("comment");


CREATE TABLE IF NOT EXISTS posts (
  id BIGINT NOT NULL UNIQUE,                      -- Public Snowflake ID (exposed)
  internal_id INTEGER PRIMARY KEY AUTOINCREMENT,  -- Internal
  user_internal_id INTEGER NOT NULL,              -- 🔑 Refers to users.internal_id
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_internal_id) REFERENCES users(internal_id)
);

CREATE TABLE IF NOT EXISTS categories (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE
);

INSERT OR IGNORE INTO categories (name) VALUES ("sport");
INSERT OR IGNORE INTO categories (name) VALUES ("finance");
INSERT OR IGNORE INTO categories (name) VALUES ("technology");
INSERT OR IGNORE INTO categories (name) VALUES ("science");

CREATE TABLE IF NOT EXISTS post_categories (
  post_internal_id INTEGER NOT NULL,
  category_id INTEGER NOT NULL,
  FOREIGN KEY (post_internal_id) REFERENCES posts(internal_id) ON DELETE CASCADE,
  FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE IF NOT EXISTS comments (
  id BIGINT UNIQUE NOT NULL,                      -- Public Snowflake ID (exposed)
  internal_id INTEGER PRIMARY KEY AUTOINCREMENT,  -- Internal
  post_internal_id INTEGER NOT NULL,              -- 🔑 Refers to posts.internal_id
  user_internal_id INTEGER NOT NULL,              -- 🔑 Refers to users.internal_id
  content TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (post_internal_id) REFERENCES posts(internal_id),
  FOREIGN KEY (user_internal_id) REFERENCES users(internal_id)
);

CREATE TABLE IF NOT EXISTS reactions (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE
);

INSERT OR IGNORE INTO reactions (name) VALUES ("like");
INSERT OR IGNORE INTO reactions (name) VALUES ("dislike");


CREATE TABLE IF NOT EXISTS item_reactions (
  item_type INTEGER NOT NULL,
  item_internal_id INTEGER NOT NULL,
  user_internal_id INTEGER NOT NULL,
  reaction_type INTEGER NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (item_type) REFERENCES items(id),
  FOREIGN KEY (reaction_type) REFERENCES reactions(id),
  FOREIGN KEY (user_internal_id) REFERENCES users(internal_id) ON DELETE CASCADE,
  UNIQUE (user_internal_id, item_internal_id, item_type)
);

