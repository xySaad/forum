package reactions

import (
	"database/sql"
	"encoding/json"
	"forum/app/handlers"
	"forum/app/modules"
	"forum/app/modules/errors"
	"net/http"
)

func RemoveReaction(conn *modules.Connection, forumDB *sql.DB) {
	var request struct {
		ItemID       string `json:"item_id"`
		ReactionType string `json:"reaction_type"` // item_id can refer to either post_id or comment_id
	}

	err := json.NewDecoder(conn.Req.Body).Decode(&request)
	if err != nil {
		http.Error(conn.Resp, "400 - Bad Request: Invalid JSON", 400)
		return
	}
	cookie, err := conn.Req.Cookie("token")
	if err != nil || cookie.Value == "" {
		conn.NewError(http.StatusUnauthorized, errors.CodeUnauthorized, "Missing or invalid authentication token", "")
		return
	}
	userID, httpErr := handlers.GetUserIDByToken(cookie.Value, forumDB)
	if httpErr != nil {
		conn.Error(httpErr)
		return
	}
	if request.ReactionType == "" || request.ItemID == "" || userID == "" {
		http.Error(conn.Resp, "400 - Bad Request: Missing fields", 400)
		return
	}

	err = handlers.RemoveReaction(request.ItemID, userID, forumDB)
	if err != nil {
		http.Error(conn.Resp, "500 - Internal Server Error", 500)
		return
	}

	conn.Resp.WriteHeader(http.StatusOK)
	conn.Resp.Write([]byte("Reaction removed successfully"))
}
