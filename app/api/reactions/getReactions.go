package reactions

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/app/handlers"
	"forum/app/modules"
	"net/http"
)

func GetReaction(conn *modules.Connection, forumDB *sql.DB) {
	reactions, err := handlers.GetReactions(conn.Path[2], forumDB)
	if err != nil {
		fmt.Println(err)
		http.Error(conn.Resp, "500 - internal server error", 500)
		return
	}

	conn.Resp.Header().Set("Content-Type", "application/json")
	conn.Resp.WriteHeader(http.StatusOK)
	fmt.Println(reactions)
	json.NewEncoder(conn.Resp).Encode(reactions)

}
