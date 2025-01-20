package comments

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
)

func UpdateComent(data io.ReadCloser) error {
	var newcomment Comment
	err := json.NewDecoder(data).Decode(&newcomment)
	if err != nil {
		return errors.New("invalid format")
	}
	db, err := sql.Open("sqlite3", "forum.db")
	if err != nil {
		return errors.New("internal pointer variable")
	}
	query := `UPDATE comments SET content=?, updated_at = CURRENT_TIMESTAMP   WHERE item_id= ?`
	_, err = db.Exec(query, newcomment.Content,newcomment.Item_id)
	if err != nil {
		return errors.New("internal pointer variable")
	}
	return nil
}
