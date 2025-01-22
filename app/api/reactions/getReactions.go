package reactions

import (
	"database/sql"
	"encoding/json"
	"fmt"
	db "forum/app/database"
	"forum/app/modules"
	"net/http"
)

func GetReaction(conn *modules.Connection, itemID string, forumDB *sql.DB) {
	reactions, err := db.GetReactions(itemID, forumDB)
	if err != nil {
		fmt.Println(err)
		http.Error(conn.Resp, "500 - internal server error", 500)
		return
	}

	conn.Resp.Header().Set("Content-Type", "application/json")
	conn.Resp.WriteHeader(http.StatusOK)
	fmt.Println(reactions)
	json.NewEncoder(conn.Resp).Encode(reactions)

}
