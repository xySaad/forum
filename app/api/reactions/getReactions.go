package reactions

import (
	"database/sql"
	"encoding/json"
	"forum/app/handlers"
	"forum/app/modules"
	"net/http"
)

func GetReaction(conn *modules.Connection, forumDB *sql.DB) {
	reactions, err := handlers.GetReactions(conn.Path[2], forumDB)
	if err != nil {
		http.Error(conn.Resp, "500 - internal server error", 500)
		return
	}

	conn.Resp.Header().Set("Content-Type", "application/json")
	conn.Resp.WriteHeader(http.StatusOK)
	json.NewEncoder(conn.Resp).Encode(reactions)

}
