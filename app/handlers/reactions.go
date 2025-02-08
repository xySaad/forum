package handlers

import (
	"database/sql"

	"forum/app/modules/errors"
	"forum/app/modules/log"
)

func GetUserIDByToken(token string, forumDB *sql.DB) (userID int, httpErr *errors.HttpError) {
	err := forumDB.QueryRow("SELECT internal_id FROM users WHERE token = ?", token).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.HttpUnauthorized
		}
		log.Error(err)
		return 0, errors.HttpInternalServerError
	}
	return
}

// getReactions fetches reactions for either a post or a comment based on itemID
func GetReactions(itemID int, item_type int, user_id string, forumDB *sql.DB) (likes, dislikes, reaction int) {
	query := `SELECT user_internal_id ,reaction_type FROM item_reactions WHERE item_internal_id =? AND item_type = ?`
	rows, err := forumDB.Query(query, itemID, item_type)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Warn(err)
		}
		return
	}
	for rows.Next() {
		cReaction := 0
		userdID := ""
		if err := rows.Scan(userdID, cReaction); err != nil {
			log.Warn(err)
			return
		}
		if userdID == user_id {
			reaction = cReaction
		}
		switch cReaction {
		case 1:
			likes++
		case 2:
			dislikes++
		default:
			log.Warn("unexpected reaction")
		}
	}
	return
}

func AddOrUpdateReaction(userID, item_type int, itemID string, reactionID int, forumDB *sql.DB) error {
	query := `SELECT * FROM item_reactions WHERE item_internal_id = ? AND user_internal_id = ? AND item_t`
	_, err := forumDB.Query(query, itemID, userID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Warn(err)
			return err
		}
		query = `INSERT INTO item_reactions (item_internal_id,user_internal_id,item_type,reaction_type) VALUES (?,?,?,?)`
		_, err := forumDB.Exec(query, itemID, userID, item_type, reactionID)
		return err
	}
	query = `UPDATE item_reactions SET  reaction_type = ? WHERE item_internal_id = ? AND user_internal_id = ? AND item_type = ?`
	_, err = forumDB.Exec(query, reactionID, itemID, userID, item_type)

	return err
}

func RemoveReaction(userID, item_type int, itemID string, forumDB *sql.DB) error {
	query := `DELETE FROM item_reactions WHERE user_internal_id = ? AND item_internal_id =? AND item_type = ?`
	_, err := forumDB.Exec(query, userID, itemID, item_type)
	if err != nil {
		log.Warn(err)
		return err
	}
	return nil
}
