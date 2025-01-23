package reactions

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
	userID, err := handlers.GetUserIDByToken(cookie.Value, forumDB)
	if err != nil {
		conn.NewError(http.StatusUnauthorized, errors.CodeUnauthorized, "Invalid or expired authentication token", "")
		return
	}
	err = handlers.AddOrUpdateReaction(request.ItemID, userID, request.ReactionType, forumDB)
	if err != nil {
		fmt.Println(err)
		conn.NewError(http.StatusInternalServerError, errors.CodeInternalServerError, "Internal Server Error", "The server encountered an error, please try again at later time.")
		return
	}

	conn.Resp.WriteHeader(http.StatusOK)
	json.NewEncoder(conn.Resp).Encode(map[string]string{"message": "Reaction added/updated successfully"})

}
