package useractivities

import (
	"database/sql"
	"net/http"

	"forum/app/api/posts"
	reactions "forum/app/api/reaction"
	"forum/app/handlers"
	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"
)

func GetUSerLiked(conn *modules.Connection, forumDB *sql.DB) {
	token, err := conn.Req.Cookie("token")
	if err != nil || token.Value == "" {
		conn.NewError(http.StatusUnauthorized, http.StatusUnauthorized, "unauthorized", "")
		return
	}
	userId, httpErr := handlers.GetUserIDByToken(token.Value, forumDB)
	if httpErr != nil {
		conn.Error(httpErr)
		return
	}
	GetUserReactions(conn, userId, "like", forumDB)
}

func GetUserReactions(conn *modules.Connection, uId, reaction string, db *sql.DB) {
	Posts := []modules.Post{}
	query := `SELECT SUBSTRING(item_id,'_',-1) FROM reactions WHERE user_id=? AND SUBSTRING(item_id,'_',1='posts') AND reaction_type=?`
	rows, err := db.Query(query, uId, reaction)
	if err != nil {
		conn.NewError(http.StatusInternalServerError, 500, "internal pointer variable", "")
		return
	}
	for rows.Next() {
		postID := ""
		if err := rows.Scan(&postID); err != nil {

			conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
			return
		}
		var post modules.Post
		query = `SELECT id, user_id, content, title, created_at FROM posts WHERE id= ? ORDER BY updated_at DESC`
		if err = db.QueryRow(query, postID).Scan(&post.ID, &post.Publisher.Id, &post.Text, &post.Title, &post.Categories, &post.CreationTime); err != nil {
			conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
			return
		}
		post.Dislikes, post.Likes, post.Reaction = reactions.GetReactions("post-"+post.ID, uId, db)

		err := posts.GetPublicUser(&post.Publisher, db)
		if err != nil {
			log.Warn(err)
		}
		herr := posts.GetPostCategories(&post.Categories, post.ID, db)
		if herr != nil {
			log.Warn(herr)
		}
		Posts = append(Posts, post)
	}
	conn.Respond(Posts)
}

func GetUSerPosts(conn *modules.Connection, db *sql.DB) {
	token, err := conn.Req.Cookie("token")
	if err != nil || token.Value == "" {
		conn.NewError(http.StatusUnauthorized, http.StatusUnauthorized, "unauthorized", "")
		return
	}
	userId, httpErr := handlers.GetUserIDByToken(token.Value, db)
	if httpErr != nil {
		conn.Error(httpErr)
		return
	}
	userPosts := []modules.Post{}
	query := `SELECT id, user_id, content, title, created_at FROM posts WHERE user_id = ? ORDER BY updated_at DESC LIMIT 10 OFFSET ?`
	rows, err := db.Query(query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			conn.NewError(http.StatusUnauthorized, 401, "unauthorized", "")
			return
		}
		conn.Error(errors.HttpInternalServerError)
	}
	defer rows.Close()

	for rows.Next() {
		var post modules.Post
		if err := rows.Scan(&post.ID, &post.Publisher.Id, &post.Text, &post.Title, &post.CreationTime); err != nil {
			conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
		}
		post.Dislikes, post.Likes, post.Reaction = reactions.GetReactions("post-"+post.ID, userId, db)
		err := posts.GetPublicUser(&post.Publisher, db)
		if err != nil {
			log.Warn(err)
		}
		herr := posts.GetPostCategories(&post.Categories, post.ID, db)
		if herr != nil {
			log.Warn(herr)
		}
		userPosts = append(userPosts, post)
	}
	conn.Respond(userPosts)
}

func Profile(conn *modules.Connection, db *sql.DB) {
	cookie, err := conn.Req.Cookie("token")
	if err != nil || cookie.Value == "" {
		conn.Error(errors.HttpUnauthorized)
		return
	}

	uuid, herr := handlers.GetUserIDByToken(cookie.Value, db)
	if herr != nil {
		conn.Error(herr)
	}
	query := `SELECT username , email,profile FROM  users WHERE uuid = ?`
	user := modules.User{}
	err = db.QueryRow(query, uuid).Scan(&user.Username, &user.Email, &user.ProfilePicture)
	if err != nil {
		if err == sql.ErrNoRows {
			conn.Error(errors.HttpUnauthorized)
			return
		}
		conn.Error(errors.HttpInternalServerError)
		return
	}
	conn.Respond(user)
}
