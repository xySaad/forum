package comments

import (
	"database/sql"
	"net/http"

	"forum/app/handlers"
	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"
)

func GetPostComments(conn *modules.Connection, forumDB *sql.DB) {
	queries := conn.Req.URL.Query()
	postId := conn.Path[2]
	lastId := queries.Get("lastId")
	params := []any{postId}
	query := "SELECT id,post_id,user_id,content,created_at FROM comments WHERE post_id = ? "
	if lastId != "" {
		params = append(params, lastId)
		query += "AND id > ?"
	}

	query += "ORDER BY id DESC LIMIT 10;"
	rows, err := forumDB.Query(query, params...)
	if err != nil {
		if err == sql.ErrNoRows {
			conn.Error(errors.HttpNotFound)
			return
		}
		conn.NewError(http.StatusInternalServerError, 500, "internal server error", "")
		return
	}
	defer rows.Close()
	var comments []modules.Comment
	conn.GetUserId(forumDB)
	for rows.Next() {
		var comment modules.Comment
		err = rows.Scan(&comment.Id, &comment.PostId, &comment.Publisher.Id, &comment.Content, &comment.CreationTime)
		if err != nil {
			conn.Error(errors.HttpInternalServerError)
			return
		}
		comment.Likes, comment.Dislikes, comment.Reaction = handlers.GetReactions(comment.Id, 2, conn.UserId, forumDB)
		err := comment.Publisher.GetPublicUser(forumDB)
		if err != nil {
			conn.Error(errors.HttpInternalServerError)
			return
		}
		comments = append(comments, comment)
	}
	log.Debug(query, params)
	conn.Respond(comments)
}
