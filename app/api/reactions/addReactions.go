package reactions

import (
	"encoding/json"
	db "forum/app/database"
	"forum/app/modules"
	"forum/app/modules/errors"
	"net/http"
)

func AddReaction(conn *modules.Connection) {
	var request struct {
		UserID       string `json:"user_id"`
		ItemID       string `json:"item_id"`       // item_id can refer to either post_id or comment_id
		ReactionType string `json:"reaction_type"` // like, love, etc.
	}

	err := json.NewDecoder(conn.Req.Body).Decode(&request)
	if err != nil {
		http.Error(conn.Resp, "400 - invalid request body", 400)
		return
	}

	if request.UserID == "" || request.ItemID == "" || request.ReactionType == "" {
		conn.NewError(http.StatusBadRequest, errors.CodeInvalidOrMissingData, "Empty Post ID/Comment ID", "")
		return
	}

	err = db.AddOrUpdateReaction(request.ItemID, request.UserID, request.ReactionType)
	if err != nil {
		conn.NewError(http.StatusInternalServerError, errors.CodeInternalServerError, "Internal Server Error", "The server encountered an error, please try again at later time.")
		return
	}

	// need to handle how the display here

}
