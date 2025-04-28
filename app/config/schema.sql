PRAGMA foreign_keys = ON;
CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY NOT NULL,
  username TEXT NOT NULL UNIQUE,
  age INTEGER,
  gender TEXT NOT NULL,
  firstname TEXT NOT NULL,
  lastname TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  profile_picture TEXT,
  password TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS message (
  id INTEGER REFERENCES users(id),
  receiver INTEGER REFERENCES users(id),
  msg TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS sessions (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER NOT NULL,
  token TEXT UNIQUE NOT NULL,
  expires_at DATETIME NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE TABLE IF NOT EXISTS items (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE
);
INSERT
  OR IGNORE INTO items (name)
VALUES ("post");
INSERT
  OR IGNORE INTO items (name)
VALUES ("comment");
CREATE TABLE IF NOT EXISTS posts (
  id INTEGER PRIMARY KEY NOT NULL,
  user_id INTEGER NOT NULL,
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE TABLE IF NOT EXISTS categories (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE
);
INSERT
  OR IGNORE INTO categories (name)
VALUES ("sport");
INSERT
  OR IGNORE INTO categories (name)
VALUES ("finance");
INSERT
  OR IGNORE INTO categories (name)
VALUES ("technology");
INSERT
  OR IGNORE INTO categories (name)
VALUES ("science");
CREATE TABLE IF NOT EXISTS post_categories (
  post_id INTEGER NOT NULL,
  -- References posts.id
  category_id INTEGER NOT NULL,
  FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
  FOREIGN KEY (category_id) REFERENCES categories(id) UNIQUE (post_id, category_id)
);
CREATE TABLE IF NOT EXISTS comments (
  id INTEGER PRIMARY KEY NOT NULL,
  post_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  content TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (post_id) REFERENCES posts(id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE TABLE IF NOT EXISTS reactions (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL UNIQUE
);
INSERT
  OR IGNORE INTO reactions (name)
VALUES ("like");
INSERT
  OR IGNORE INTO reactions (name)
VALUES ("dislike");
CREATE TABLE IF NOT EXISTS item_reactions (
  item_type INTEGER NOT NULL,
  item_id INTEGER NOT NULL,
  user_id INTEGER NOT NULL,
  reaction_id INTEGER NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (item_type) REFERENCES items(id),
  FOREIGN KEY (reaction_id) REFERENCES reactions(id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  UNIQUE (user_id, item_id, item_type)
);