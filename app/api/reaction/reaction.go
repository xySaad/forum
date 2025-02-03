package reactions

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"forum/app/handlers"
	"forum/app/modules"
	"forum/app/modules/errors"
)

func HandleReactions(conn *modules.Connection, db *sql.DB) {
	cookie, err := conn.Req.Cookie("token")
	if err != nil || cookie.Value == "" {
		conn.Error(errors.HttpUnauthorized)
		return
	}
	uid, Herr := handlers.GetUserIDByToken(cookie.Value, db)
	if Herr != nil || cookie.Value == "" {
		conn.Error(Herr)
		return
	}
	reationType := conn.Req.URL.Query().Get("reaction")
	item_id := conn.Req.URL.Query().Get("item_id")
	if item_id == "" {
		conn.NewError(http.StatusBadRequest, 400, "missing data", "")
		return
	}
	item_idSP := strings.Split(item_id, "-")
	if len(item_id) != 2 {
		conn.NewError(http.StatusBadRequest, 400, "invalid item id", "")
		return
	}
	if n, err := strconv.Atoi(item_idSP[1]); err != nil || n <= 0 {
		conn.NewError(http.StatusBadRequest, 400, "invalid item id", "id should be a number greater than 0")
		return
	}
	if item_idSP[0] != "comment" && item_idSP[0] != "post" {
		conn.NewError(http.StatusBadRequest, 400, "invalid item id", "item van only be comment or post")
		return
	}
	if reationType != "like" && reationType != "dislike" {
		conn.NewError(http.StatusBadRequest, 400, "invalid reaction", "only like and dislke are suppported")
		return
	}
	herr := UpdateReaction(reationType, uid, item_id, db)
	if herr != nil {
		conn.Error(herr)
		return
	}
}

func UpdateReaction(reatcion string, uid string, itemID string, db *sql.DB) *errors.HttpError {
	if reatcion == "" {
		query := `DELETE FROM reactions WHERE user_id = ? AND item_id = ?`
		_, err := db.Exec(query, uid, itemID)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil
			}
			return errors.HttpInternalServerError
		}
		return nil
	}
	query := `SELECT * FROM reactions WHERE user_id = ? AND item_id = ?`
	_, err := db.Exec(query, uid, itemID)
	if err != nil {
		if err == sql.ErrNoRows {
			query=`INSERT INTO reactions (item_id, user_id, reaction_type) VALUES (?, ?, ?)`
		} else {
			return errors.HttpInternalServerError
		}
	} else {
		query = `UPDATE reactions SET reaction_type = ? WHERE user_id = ? AND item_id = ?`
	}
	_, err = db.Exec(query, uid, itemID)
	if err != nil {
		return errors.HttpInternalServerError
	}
	return nil
}
