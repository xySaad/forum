package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"forum/app/handlers"
	"forum/app/modules"
	"forum/app/modules/errors"
)

func GetUserData(conn *modules.Connection, forumDB *sql.DB) {
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

	type User struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	var userD User
	qurr := `SELECT username, email FROM users WHERE internal_id = ?`
	fmt.Println(userID)
	err = forumDB.QueryRow(qurr, userID).Scan(&userD.Username, &userD.Email)
	if err != nil {
		fmt.Println(err)
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
