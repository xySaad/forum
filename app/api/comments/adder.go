package comments

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"forum/app/handlers"
	"forum/app/modules"
)

func AddComment(conn *modules.Connection, forumDB *sql.DB) error {
	var comment Comment
	err := json.NewDecoder(conn.Req.Body).Decode(&comment)
	if err != nil {
		return errors.New("invalid data format")
	}
	cookie, err := conn.Req.Cookie("token")
	if err != nil || cookie.Value == "" {
		return errors.New("unotorized")
	}
	uId, err := handlers.GetUserIDByToken(cookie.Value, forumDB)
	if err != nil {
		return err
	}

	query := `INSERT INTO comments (post_id, user_id, content) VALUES (?, ?, ?)`
	_, err = forumDB.Exec(query, comment.Post_id, uId, comment.Content)
	if err != nil {
		return errors.New("internal pointer variable")
	}
	fmt.Printf("comment: %v\n", comment)
	return nil
}
