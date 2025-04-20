package profile

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"forum/app/modules"
	"forum/app/modules/log"
)

func GetUserData(conn *modules.Connection, forumDB *sql.DB) {
	if !conn.IsAuthenticated(forumDB) {
		return
	}

	type User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	var userD User
	qurr := `SELECT username, email FROM users WHERE id = ?`
	err := forumDB.QueryRow(qurr, conn.User).Scan(&userD.Username, &userD.Email)
	if err != nil {
		log.Error(err)
		http.Error(conn.Resp, "Database query error", http.StatusInternalServerError)
		return
	}
	conn.Resp.Header().Set("Content-Type", "application/json")
	conn.Resp.WriteHeader(http.StatusOK)

	jsond, err := json.Marshal(userD)
	if err != nil {
		http.Error(conn.Resp, "Error in JSON encoding", http.StatusInternalServerError)
		return
	}
	conn.Resp.Write(jsond)
}
