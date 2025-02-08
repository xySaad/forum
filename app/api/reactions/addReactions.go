package reactions

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"forum/app/handlers"
	"forum/app/modules"
	"forum/app/modules/errors"
)

func AddReaction(conn *modules.Connection, forumDB *sql.DB) {
	if !conn.IsAuthenticated(forumDB) {
		return
	}

	validReactions := map[string]int{
		"like":    1,
		"dislike": 2,
	}

	var request struct {
		ItemID       string `json:"item_id"`
		ReactionType string `json:"reaction_type"`
		Item_type    int `json:"item_type"`
	}

	err := json.NewDecoder(conn.Req.Body).Decode(&request)
	if err != nil {
		http.Error(conn.Resp, "400 - invalid request body", 400)
		return
	}

	if request.ItemID == "" || request.ReactionType == "" || request.Item_type == 0 {
		conn.NewError(http.StatusBadRequest, errors.CodeInvalidOrMissingData, "Empty Post ID/Comment ID", "")
		return
	}
	reactionID, exist := validReactions[request.ReactionType]
	if !exist {
		conn.NewError(http.StatusBadRequest, errors.CodeInvalidOrMissingData, "Invalid Reaction Type", "")
		return
	}

	err = handlers.AddOrUpdateReaction(conn.InternalUserId, request.Item_type, request.ItemID, reactionID, forumDB)
	if err != nil {
		conn.NewError(http.StatusInternalServerError, errors.CodeInternalServerError, "Internal Server Error", "The server encountered an error, please try again at later time.")
		return
	}

	conn.Resp.WriteHeader(http.StatusOK)
	json.NewEncoder(conn.Resp).Encode(map[string]string{"message": "Reaction added/updated successfully"})
}
