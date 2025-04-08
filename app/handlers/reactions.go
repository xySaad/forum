package handlers

import (
	"database/sql"

	"forum/app/modules/log"
)

func GetReactions(itemID string, itemType int, userID int, forumDB *sql.DB) (likes, dislikes int, reaction string) {
	err := forumDB.QueryRow(`
        SELECT 
            SUM(CASE WHEN reaction_id = 1 THEN 1 ELSE 0 END),
            SUM(CASE WHEN reaction_id = 2 THEN 1 ELSE 0 END)
        FROM item_reactions
        WHERE item_id = ? AND item_type = ?`,
		itemID, itemType).Scan(&likes, &dislikes)

	if err != nil && err != sql.ErrNoRows {
		log.Error("Error counting reactions:", err)
		return
	}

	if userID != 0 {
		err = forumDB.QueryRow(`SELECT r.name FROM item_reactions
			JOIN reactions r ON reaction_id = r.id
			WHERE item_id = ? AND item_type = ? AND user_id = ?`,
			itemID, itemType, userID).Scan(&reaction)

		if err != nil && err != sql.ErrNoRows {
			log.Error(err)
		}
	}

	return
}
