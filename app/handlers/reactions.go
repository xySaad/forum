package handlers

import (
	"database/sql"

	"forum/app/modules/log"
)

// getReactions fetches reactions for either a post or a comment based on itemID
func GetReactions(itemID int, item_type int, userId int, forumDB *sql.DB) (likes, dislikes, reaction int) {
	query := `SELECT user_id ,reaction_type FROM item_reactions WHERE item_internal_id =? AND item_type = ?`
	rows, err := forumDB.Query(query, itemID, item_type)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Error(err)
		}
		return
	}
	for rows.Next() {
		cReaction := 0
		var userdID int
		if err := rows.Scan(userdID, cReaction); err != nil {
			log.Warn(err)
			return
		}
		if userdID == userId {
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
