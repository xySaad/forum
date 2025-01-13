package reactions

import (
	"encoding/json"
	"forum/app/database"
	"forum/app/modules"
	"net/http"
)

// RemoveReaction handles the removal of a reaction from a post or comment
func RemoveReaction(conn *modules.Connection) {
	var request struct {
		UserID string `json:"user_id"`
		ItemID string `json:"item_id"` // item_id can refer to either post_id or comment_id
	}

	// Decode the JSON request body
	err := json.NewDecoder(conn.Req.Body).Decode(&request)
	if err != nil {
		http.Error(conn.Resp, "400 - Bad Request: Invalid JSON", 400)
		return
	}

	// Ensure that all required fields are provided
	if request.UserID == "" || request.ItemID == "" {
		http.Error(conn.Resp, "400 - Bad Request: Missing fields", 400)
		return
	}

	// Call the RemoveReaction function from the database package to remove the reaction
	err = db.RemoveReaction(request.ItemID, request.UserID)
	if err != nil {
		http.Error(conn.Resp, "500 - Internal Server Error", 500)
		return
	}

	// Respond with success
	conn.Resp.WriteHeader(http.StatusOK)
	conn.Resp.Write([]byte("Reaction removed successfully"))
}
