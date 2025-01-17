package reactions

import (
	"encoding/json"
	"fmt"
	db "forum/app/database"
	"forum/app/modules"
	"net/http"
)

func GetReaction(conn *modules.Connection, itemID string) {
	reactions, err := db.GetReactions(itemID)
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
