package comments

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"forum/app/modules"
	"forum/app/modules/errors"
)

func UpdateComment(conn *modules.Connection, forumdb *sql.DB) {
	var newcomment Comment
	err := json.NewDecoder(conn.Req.Body).Decode(&newcomment)
	if err != nil {
		conn.NewError(http.StatusBadRequest, 400, "ivalid format", "")
		return
	}

	query := `UPDATE comments SET content=?, updated_at = CURRENT_TIMESTAMP WHERE id= ?`
	_, err = forumdb.Exec(query, newcomment.Content, newcomment.ItemID)
	if err != nil {
		conn.Error(errors.HttpInternalServerError)
	}
}
