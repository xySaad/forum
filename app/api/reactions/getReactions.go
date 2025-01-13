package reactions

import (
	"encoding/json"
	db "forum/app/database"
	"forum/app/modules"
	"net/http"
)

func GetReaction(conn *modules.Connection, itemID string) {
	reactions, err := db.GetReactions(itemID)
	if err != nil {
		http.Error(conn.Resp, "500 - internal server error", 500)
		return
	}

	conn.Resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(conn.Resp).Encode(reactions)
}
