package useractivities

import (
	"database/sql"
	"net/http"
	"strings"

	"forum/app/handlers"
	"forum/app/modules"
)

func GetUSer(conn *modules.Connection, forumDB *sql.DB) {
	SpUrl := strings.Split(conn.Req.URL.String(), "/")
	if len(SpUrl) != 3 {
		conn.NewError(http.StatusNotFound, 404, "not found", "")
		return
	}
	token, err := conn.Req.Cookie("token")
	if err != nil || token.Value == "" {
		conn.NewError(http.StatusUnauthorized, http.StatusUnauthorized, "unauthorized", "")
		return
	}
	_, httpErr := handlers.GetUserIDByToken(token.Value, forumDB)
	if httpErr != nil {
		conn.Error(httpErr)
		return
	}
	switch SpUrl[2] {
	case "posts":
		// GetUSerPosts(conn, userId, forumDB)
	case "like", "dislike":
		// GetUserReactions(conn, userId, SpUrl[2], forumDB)
	default:
		conn.NewError(http.StatusNotFound, 404, "not found", "")
	}
}

// func GetUserReactions(conn *modules.Connection, uId int, reaction string, db *sql.DB) {
// 	Posts := []modules.Post{}
// 	query := `SELECT SUBSTRING_INDEX(item_id,'_',-1) FROM reactions WHERE user_id=? AND SUBSTRING_INDEX(item_id,'_',1='posts') AND reaction_type=?`
// 	rows, err := db.Query(query, uId, reaction)
// 	if err != nil {
// 		conn.NewError(http.StatusInternalServerError, 500, "internal pointer variable", "")
// 		return
// 	}
// 	for rows.Next() {
// 		postID := ""
// 		if err := rows.Scan(&postID); err != nil {

// 			conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
// 			return
// 		}
// 		var post modules.Post
// 		query = `SELECT id, user_id, content, title, created_at FROM posts WHERE id= ? ORDER BY updated_at DESC`
// 		if err = db.QueryRow(query, postID).Scan(&post.ID, &post.Publisher.Id, &post.Text, &post.Title, &post.Categories, &post.CreationTime); err != nil {
// 			conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
// 			return
// 		}
// 		err := posts.GetPublicUser(&post.Publisher, db)
// 		if err != nil {
// 			log.Warn(err)
// 		}
// 		herr := posts.GetPostCategories(post.Categories, post.ID, db)
// 		if herr != nil {
// 			log.Warn(herr)
// 		}
// 		Posts = append(Posts, post)
// 	}
// 	err = json.NewEncoder(conn.Resp).Encode(Posts)
// 	if err != nil {
// 		conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
// 	}
// }

// func GetUSerPosts(conn *modules.Connection, userId int, db *sql.DB) {
// 	userPosts := []modules.Post{}
// 	query := `SELECT id, user_id, content, title, created_at FROM posts WHERE user_id = ? ORDER BY updated_at DESC LIMIT 10 OFFSET ?`
// 	rows, err := db.Query(query, userId)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			conn.NewError(http.StatusUnauthorized, 401, "unauthorized", "")
// 			return
// 		}
// 		conn.Error(errors.HttpInternalServerError)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var post modules.Post
// 		if err := rows.Scan(&post.ID, &post.Publisher.Id, &post.Text, &post.Title, &post.CreationTime); err != nil {
// 			conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
// 		}
// 		err := posts.GetPublicUser(&post.Publisher, db)
// 		if err != nil {
// 			log.Warn(err)
// 		}
// 		herr := posts.GetPostCategories(post.Categories, post.ID, db)
// 		if herr != nil {
// 			log.Warn(herr)
// 		}
// 		userPosts = append(userPosts, post)
// 	}
// 	err = json.NewEncoder(conn.Resp).Encode(userPosts)
// 	if err != nil {
// 		conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
// 	}
// }
