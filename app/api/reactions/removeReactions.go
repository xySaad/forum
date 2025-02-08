package reactions

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"forum/app/handlers"
	"forum/app/modules"
)

func RemoveReaction(conn *modules.Connection, forumDB *sql.DB) {
	if !conn.IsAuthenticated(forumDB) {
		return
	}
	var request struct {
		ItemID       string `json:"item_id"`
		ReactionType string `json:"reaction_type"`
		Item_type    int    `json:"item_type"`
	}

	err := json.NewDecoder(conn.Req.Body).Decode(&request)
	if err != nil {
		http.Error(conn.Resp, "400 - Bad Request: Invalid JSON", 400)
		return
	}

	if request.ReactionType == "" || request.ItemID == "" || request.Item_type == 0 {
		http.Error(conn.Resp, "400 - Bad Request: Missing fields", 400)
		return
	}

	err = handlers.RemoveReaction(conn.InternalUserId, request.Item_type, request.ItemID, forumDB)
	if err != nil {
		http.Error(conn.Resp, "500 - Internal Server Error", 500)
		return
	}

	conn.Resp.WriteHeader(http.StatusOK)
	conn.Resp.Write([]byte("Reaction removed successfully"))
}
