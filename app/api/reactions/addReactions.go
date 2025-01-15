package reactions

import (
	"encoding/json"
	db "forum/app/database"
	"forum/app/modules"
	"forum/app/modules/errors"
	"net/http"
)

func AddReaction(conn *modules.Connection) {

	validReactions := map[string]bool{
		"like":    true,
		"dislike": true,
	}

	var request struct {
		ItemID       string `json:"item_id"`       // item_id can refer to either post_id or comment_id
		ReactionType string `json:"reaction_type"` // like, love, etc.
	}

	err := json.NewDecoder(conn.Req.Body).Decode(&request)
	if err != nil {
		http.Error(conn.Resp, "400 - invalid request body", 400)
		return
	}

	if request.ItemID == "" || request.ReactionType == "" {
		conn.NewError(http.StatusBadRequest, errors.CodeInvalidOrMissingData, "Empty Post ID/Comment ID", "")
		return
	}

	if !validReactions[request.ReactionType] {
		conn.NewError(http.StatusBadRequest, errors.CodeInvalidOrMissingData, "Invalid Reaction Type", "")
		return
	}

	cookie, err := conn.Req.Cookie("token")
	if err != nil || cookie.Value == "" {
		conn.NewError(http.StatusUnauthorized, errors.CodeUnauthorized, "Missing or invalid authentication token", "")
		return
	}
	userID, err := db.GetUserIDByToken(cookie.Value)
	if err != nil {
		conn.NewError(http.StatusUnauthorized, errors.CodeUnauthorized, "Invalid or expired authentication token", "")
		return
	}
	err = db.AddOrUpdateReaction(request.ItemID, userID, request.ReactionType)
	if err != nil {
		conn.NewError(http.StatusInternalServerError, errors.CodeInternalServerError, "Internal Server Error", "The server encountered an error, please try again at later time.")
		return
	}

	conn.Resp.WriteHeader(http.StatusOK)
	json.NewEncoder(conn.Resp).Encode(map[string]string{"message": "Reaction added/updated successfully"})

}
