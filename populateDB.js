const sqlite3 = require('sqlite3').verbose();
const db = new sqlite3.Database('forum.db');

// Function to generate random strings for users and posts
function randomString(length) {
    const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    let result = '';
    for (let i = 0; i < length; i++) {
        result += chars.charAt(Math.floor(Math.random() * chars.length));
    }
    return result;
}

// Insert fake users
function insertUsers(numUsers = 10) {
    for (let i = 0; i < numUsers; i++) {
        const userId = randomString(12);
        const username = randomString(8);
        const email = username + '@example.com';
        const password = randomString(12);
        const token = randomString(50);
        
        db.run(`
            INSERT INTO users (id, username, email, password, token) 
            VALUES (?, ?, ?, ?, ?)`, [userId, username, email, password, token]);
    }
}

// Insert fake posts
function insertPosts(numPosts = 20) {
    db.all("SELECT id FROM users", (err, rows) => {
        if (err) {
            console.error(err);
            return;
        }
        
        const userIds = rows.map(row => row.id);
        
        for (let i = 0; i < numPosts; i++) {
            const userId = userIds[Math.floor(Math.random() * userIds.length)];
            const title = randomString(15);
            const content = randomString(100);
            const categories = Math.floor(Math.random() * 15) + 1; // Random bitmask for categories (1 to 15)
            
            db.run(`
                INSERT INTO posts (user_id, title, content, categories) 
                VALUES (?, ?, ?, ?)`, [userId, title, content, categories]);
        }
    });
}

// Insert fake comments
function insertComments(numComments = 30) {
    db.all("SELECT id FROM posts", (err, rows) => {
        if (err) {
            console.error(err);
            return;
        }

        const postIds = rows.map(row => row.id);

        db.all("SELECT id FROM users", (err, rows) => {
            if (err) {
                console.error(err);
                return;
            }

            const userIds = rows.map(row => row.id);

            for (let i = 0; i < numComments; i++) {
                const postId = postIds[Math.floor(Math.random() * postIds.length)];
                const userId = userIds[Math.floor(Math.random() * userIds.length)];
                const content = randomString(50);

                db.run(`
                    INSERT INTO comments (post_id, user_id, content) 
                    VALUES (?, ?, ?)`, [postId, userId, content]);
            }
        });
    });
}

// Insert fake reactions
function insertReactions(numReactions = 50) {
    db.all("SELECT item_id FROM posts", (err, rows) => {
        if (err) {
            console.error(err);
            return;
        }

        const postIds = rows.map(row => row.item_id);

        db.all("SELECT item_id FROM comments", (err, rows) => {
            if (err) {
                console.error(err);
                return;
            }

            const commentIds = rows.map(row => row.item_id);
            const itemIds = postIds.concat(commentIds);
            const reactionTypes = ['like', 'dislike'];

            for (let i = 0; i < numReactions; i++) {
                const itemId = itemIds[Math.floor(Math.random() * itemIds.length)];
                const userId = randomString(12); // Assuming random users for reactions
                const reactionType = reactionTypes[Math.floor(Math.random() * reactionTypes.length)];

                db.run(`
                    INSERT INTO reactions (user_id, item_id, reaction_type) 
                    VALUES (?, ?, ?)`, [userId, itemId, reactionType]);
            }
        });
    });
}

// Function to populate the database with fake data
function populateDB() {
    insertUsers(10);
    setTimeout(() => {
        insertPosts(20);
        setTimeout(() => {
            insertComments(30);
            setTimeout(() => {
                insertReactions(50);
                console.log("Database populated with fake data.");
            }, 500);
        }, 500);
    }, 500);
}

// Run the populate function
populateDB();

// Close the DB connection after some time (to ensure all queries are done)
setTimeout(() => {
    db.close();
}, 3000);
