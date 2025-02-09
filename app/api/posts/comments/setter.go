package comments

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"
	"forum/app/modules/snowflake"

	"github.com/mattn/go-sqlite3"
)

func AddComment(conn *modules.Connection, forumDB *sql.DB) {
	if !conn.IsAuthenticated(forumDB) {
		return
	}

	var comment modules.Comment
	postId := conn.Path[2]
	err := json.NewDecoder(conn.Req.Body).Decode(&comment)
	if err != nil {
		conn.Error(errors.BadRequestError("invalid format"))
		return
	}
	if comment.Content == "" {
		conn.Error(errors.BadRequestError("missing content"))
		return
	}

	query := `INSERT INTO comments (id, post_id, user_id, content) VALUES (?, ?, ?, ?)`
	_, err = forumDB.Exec(query, snowflake.Default.Generate(), postId, conn.UserId, comment.Content)
	if err != nil {
		if sqlErr, ok := err.(sqlite3.Error); ok && sqlErr.ExtendedCode == sqlite3.ErrConstraintForeignKey {
			conn.Error(errors.BadRequestError("invalid post id"))
			return
		}
		log.Error(err)
		conn.Error(errors.HttpInternalServerError)
		return
	}
	conn.Resp.Write([]byte{'o', 'k'})
}

func UpdateComment(conn *modules.Connection, forumdb *sql.DB) {
	if conn.IsAuthenticated(forumdb) {
		return
	}

	var newcomment modules.Comment
	err := json.NewDecoder(conn.Req.Body).Decode(&newcomment)
	if err != nil {
		conn.NewError(http.StatusBadRequest, 400, "ivalid format", "")
		return
	}

	query := `UPDATE comments SET content=?, updated_at = CURRENT_TIMESTAMP WHERE id= ?`
	_, err = forumdb.Exec(query, newcomment.Content, newcomment.Id)
	if err != nil {
		conn.Error(errors.HttpInternalServerError)
	}
}
