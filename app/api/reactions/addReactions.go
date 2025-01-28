package reactions

import (
	"database/sql"
	"encoding/json"
	"forum/app/handlers"
	"forum/app/modules"
	"forum/app/modules/errors"
	"net/http"
)

func AddReaction(conn *modules.Connection, forumDB *sql.DB) {

	validReactions := map[string]bool{
		"like":    true,
		"dislike": true,
	}

	var request struct {
		ItemID       string `json:"item_id"`
		ReactionType string `json:"reaction_type"`
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
	userID, httpErr := handlers.GetUserIDByToken(cookie.Value, forumDB)
	if httpErr != nil {
		conn.Error(httpErr)
		return
	}
	err = handlers.AddOrUpdateReaction(request.ItemID, userID, request.ReactionType, forumDB)
	if err != nil {
		conn.NewError(http.StatusInternalServerError, errors.CodeInternalServerError, "Internal Server Error", "The server encountered an error, please try again at later time.")
		return
	}

	conn.Resp.WriteHeader(http.StatusOK)
	json.NewEncoder(conn.Resp).Encode(map[string]string{"message": "Reaction added/updated successfully"})

}
