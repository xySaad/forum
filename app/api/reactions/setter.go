package reactions

import (
	"database/sql"
	"net/http"

	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"

	"github.com/mattn/go-sqlite3"
)

func AddReaction(conn *modules.Connection, forumDB *sql.DB) {
	if len(conn.Path) < 5 {
		conn.Error(errors.HttpNotFound)
		return
	}

	if !conn.IsAuthenticated(forumDB) {
		return
	}
	itemType := conn.Path[2]
	itemId := conn.Path[3]
	reactionType := conn.Path[4]
	if itemType != "posts" && itemType != "comments" {
		conn.Error(errors.BadRequestError("invalid item"))
		return
	}
	exists := false
	sqlQuery := "select 1 from " + itemType + " where id = ?;"
	err := forumDB.QueryRow(sqlQuery, itemId).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			conn.Error(errors.BadRequestError("invalid item id"))
			return
		}
		log.Error(err)
		conn.Error(errors.HttpInternalServerError)
		return
	}

	query := `INSERT INTO item_reactions (item_type, item_id, user_id, reaction_id)
	VALUES (((SELECT id from items where name = ?)), ?, ?, (SELECT id from reactions where name = ?))
	ON CONFLICT(user_id, item_id, item_type) DO UPDATE SET reaction_id = excluded.reaction_id;`

	_, err = forumDB.Exec(query, itemType[:len(itemType)-1], itemId, conn.User.Id, reactionType)
	if err != nil {
		log.Error(err)
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqliteErr.ExtendedCode == sqlite3.ErrConstraintNotNull {
				conn.Error(errors.BadRequestError("invalid reaction type"))
				return
			}
		}
		conn.Error(errors.HttpInternalServerError)
		return
	}

	conn.Resp.Header().Set("Content-Type", "application/json")
	conn.Resp.Write([]byte(`{"message": "Reaction added/updated successfully"}`))
}

func RemoveReaction(conn *modules.Connection, forumDB *sql.DB) {
	if len(conn.Path) < 5 {
		conn.Error(errors.HttpNotFound)
		return
	}
	if !conn.IsAuthenticated(forumDB) {
		return
	}
	itemType := conn.Path[2]
	itemId := conn.Path[3]
	if itemType != "posts" && itemType != "comments" {
		conn.Error(errors.BadRequestError("invalid item"))
		return
	}

	sqlQuery := "DELETE FROM item_reactions WHERE user_id = ? AND item_id =?"
	result, err := forumDB.Exec(sqlQuery, conn.User.Id, itemId)
	if err != nil {
		conn.Error(errors.HttpInternalServerError)
		return
	}
	n, err := result.RowsAffected()
	if err != nil {
		log.Error(err)
	}
	if n == 0 {
		conn.Error(errors.BadRequestError("invalid item id"))
		return
	}
	conn.Resp.WriteHeader(http.StatusOK)
	conn.Resp.Write([]byte("Reaction removed successfully"))
}
