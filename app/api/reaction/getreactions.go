package reactions

import (
	"database/sql"

	"forum/app/modules/log"
)

func GetReactions(itemId string, uuid string, db *sql.DB) (DislikesCount, LikesCount, UserReaction int) {
	query := `SELECT reaction_type user_id FROM reactions WHERe item_id=?`
	rows, err := db.Query(query, itemId)
	if err != nil {
		log.Warn(err)
		return
	}
	defer rows.Close()
	r_type := ""
	uid := ""
	for rows.Next() {
		err = rows.Scan(&r_type, &uid)
		if err != nil {
			log.Warn(err)
			return
		}
		if r_type=="like" {
			LikesCount++
		}else if r_type=="dislike" {
			DislikesCount++
		}else{
			log.Warn("unexpected r_type at item_id=" + itemId)

		}
		if uid == uuid && uuid != "" {
		if r_type == "like" {
				UserReaction = 1
		} else if r_type == "dislike" {
				UserReaction = -1
		}
	}
}
	return
}
