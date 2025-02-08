package handlers

import (
	"database/sql"
	"fmt"

	"forum/app/modules/errors"
	"forum/app/modules/log"
)

type Reaction struct {
	UserID       string `json:"user_id"`
	ReactionType string `json:"reaction_type"`
	Timestamp    string `json:"timestamp"`
}

type ReactionCounter struct {
	Item_id      string
	ReactionType string `json:"reaction_type"`
	Count        int
}

func GetUserIDByToken(token string, forumDB *sql.DB) (userID int, httpErr *errors.HttpError) {
	err := forumDB.QueryRow("SELECT internal_id FROM users WHERE token = ?", token).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.HttpUnauthorized
		}
		log.Error(err)
		return 0, errors.HttpInternalServerError
	}
	return
}

// getReactions fetches reactions for either a post or a comment based on itemID
func GetReactions(itemID string, forumDB *sql.DB) ([]ReactionCounter, error) {
	if itemID == "" {
		return nil, fmt.Errorf("itemID must be provided")
	}
	// Query reactions for both posts and comments using the same itemID
	rows, err := forumDB.Query(`
		SELECT item_id, reaction_type, count 
		FROM reactionCount 
		WHERE item_id = ?`, itemID)
	if err != nil {
		return nil, fmt.Errorf("could not fetch reactions: %w", err)
	}
	defer rows.Close()

	// prepare a slice to store the results
	var reactions []ReactionCounter
	for rows.Next() {
		var reaction ReactionCounter
		if err := rows.Scan(&reaction.Item_id, &reaction.ReactionType, &reaction.Count); err != nil {
			return nil, fmt.Errorf("could not scan row: %w", err)
		}
		reactions = append(reactions, reaction)
	}

	return reactions, nil
}

func AddOrUpdateReaction(userID, item_type int, itemID string, reactionID int, forumDB *sql.DB) error {
	query := `SELECT * FROM item_reactions WHERE item_internal_id = ? AND user_internal_id = ? AND item_t`
	_, err := forumDB.Query(query, itemID, userID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Warn(err)
			return err
		}
		query = `INSERT INTO item_reactions (item_internal_id,user_internal_id,item_type,reaction_type) VALUES (?,?,?,?)`
		_, err := forumDB.Exec(query, itemID, userID, item_type, reactionID)
		return err
	}
	query = `UPDATE item_reactions SET  reaction_type = ? WHERE item_internal_id = ? AND user_internal_id = ? AND item_type = ?`
	_, err = forumDB.Exec(query, reactionID, itemID, userID, item_type)

	return err
}

func RemoveReaction(userID, item_type int, itemID string, forumDB *sql.DB) error {
	query := `DELETE FROM item_reactions WHERE user_internal_id = ? AND item_internal_id =? AND item_type = ?`
	_, err := forumDB.Exec(query, userID, itemID, item_type)
	if err != nil {
		log.Warn(err)
		return err
	}
	return nil
}
