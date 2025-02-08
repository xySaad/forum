package comments

import (
	"database/sql"
	"net/http"
	"strconv"

	"forum/app/api/posts"
	"forum/app/handlers"
	"forum/app/modules"
	"forum/app/modules/log"
)

func GetComments(conn *modules.Connection, forumDB *sql.DB) {
	URL := conn.Req.URL
	post_id := URL.Query().Get("p_id")
	Offset := URL.Query().Get("offset")
	from := URL.Query().Get("from")
	user_id := ""
	cookie, err := conn.Req.Cookie("token")
	if err == nil && cookie.Value != "" {
		forumDB.QueryRow(`SELECT user_internal_id FROM users WHERE token =?`, cookie.Value).Scan(&user_id)
	}
	if from == "" {
		conn.NewError(http.StatusBadRequest, 400, "invalid data", "")
		return
	}
	fromN, err := strconv.Atoi(from)
	if err != nil {
		conn.NewError(http.StatusBadRequest, 400, "invalid data", "")
		return
	}
	if Offset == "" {
		conn.NewError(http.StatusBadRequest, 400, "invalid data", "")
		return
	}
	offset, err := strconv.Atoi(Offset)
	if err != nil {
		conn.NewError(http.StatusBadRequest, 400, "invalid data", "")
		return
	}
	if post_id == "" {
		conn.NewError(http.StatusBadRequest, 400, "invalid data", "")
		return
	}
	p_id, err := strconv.Atoi(post_id)
	if err != nil {
		conn.NewError(http.StatusBadRequest, 400, "invalid data", "")
		return
	}

	query := `SELECT id, post_internal_id, user_internal_id, content, created_at 
              FROM comments WHERE `

	if fromN > 0 {
		query += `id <= ? AND `
	}

	query += `post_internal_id = ? ORDER BY created_at DESC LIMIT 10 OFFSET ?`

	rows, err := forumDB.Query(query, fromN, p_id, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			conn.NewError(http.StatusNotFound, 404, "no post with such id", "")
			return
		}
		conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
		return
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.ItemID, &comment.PostID, &comment.Publisher.Id, &comment.Content, &comment.CreationTime); err != nil {
			log.Warn(err)
			conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
			return
		}
		comment.Likes, comment.Dislikes, comment.Reaction = handlers.GetReactions(comment.ItemID, 2, user_id, forumDB)
		comment.Publisher, err = posts.GetPublicUser(comment.Publisher.Id, forumDB)
		if err != nil {
			log.Warn(err)
			conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
			return
		}

		comments = append(comments, comment)
	}

	conn.Respond(comments)
}
