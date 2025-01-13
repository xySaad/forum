package db

import (
	"database/sql"
	"fmt"
	"log"
)

type Reaction struct {
	UserID       string `json:"user_id"`
	ReactionType string `json:"reaction_type"`
	Timestamp    string `json:"timestamp"`
}

// getReactions fetches reactions for either a post or a comment based on itemID
func GetReactions(itemID string) ([]Reaction, error) {
	if itemID == "" {
		return nil, fmt.Errorf("itemID must be provided")
	}

	// Connect to the database
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query reactions for both posts and comments using the same itemID
	rows, err := db.Query(`
		SELECT user_id, reaction_type, created_at 
		FROM reactions 
		WHERE post_id = ? OR comment_id = ?`, itemID, itemID)

	if err != nil {
		return nil, fmt.Errorf("could not fetch reactions: %w", err)
	}
	defer rows.Close()

	//prepare a slice to store the results
	var reactions []Reaction
	for rows.Next() {
		var reaction Reaction
		if err := rows.Scan(&reaction.UserID, &reaction.ReactionType, &reaction.Timestamp); err != nil {
			return nil, fmt.Errorf("could not scan row: %w", err)
		}
		reactions = append(reactions, reaction)
	}

	return reactions, nil
}

func AddOrUpdateReaction(itemID, userID, reactionType string) error {
	// Connect to the database
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the reaction already exists for the given post/comment and user
	var existingReactionID int
	err = db.QueryRow(`
		SELECT id FROM reactions 
		WHERE (post_id = ? OR comment_id = ?) 
		AND user_id = ?`, itemID, itemID, userID).Scan(&existingReactionID)

	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("could not check for existing reaction: %w", err)
	}

	// If the reaction already exists, update it
	if existingReactionID > 0 {
		_, err = db.Exec(`
			UPDATE reactions 
			SET reaction_type = ? 
			WHERE id = ?`, reactionType, existingReactionID)
		if err != nil {
			return fmt.Errorf("could not update reaction: %w", err)
		}
	} else {
		// If no existing reaction, insert a new one
		_, err = db.Exec(`
			INSERT INTO reactions (user_id, post_id, comment_id, reaction_type) 
			VALUES (?, ?, ?, ?)`, userID, itemID, itemID, reactionType)

		if err != nil {
			return fmt.Errorf("could not insert reaction: %w", err)
		}
	}

	// Update the reaction count in the reactionCount table
	_, err = db.Exec(`
		INSERT INTO reactionCount (post_id, comment_id, reaction_type, count) 
		VALUES (?, ?, ?, 1) 
		ON CONFLICT(post_id, comment_id, reaction_type) 
		DO UPDATE SET count = count + 1`, itemID, itemID, reactionType)

	if err != nil {
		return fmt.Errorf("could not update reaction count: %w", err)
	}

	return nil
}

// RemoveReaction removes a reaction from a post or comment
func RemoveReaction(itemID, userID string) error {
	// Connect to the database
	db, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the reaction exists for the given post/comment and user
	var reactionID int
	err = db.QueryRow(`
		SELECT id FROM reactions 
		WHERE (post_id = ? OR comment_id = ?) 
		AND user_id = ?`, itemID, itemID, userID).Scan(&reactionID)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("reaction not found for the given user and item")
		}
		return fmt.Errorf("could not check for existing reaction: %w", err)
	}

	// Delete the reaction from the reactions table
	_, err = db.Exec(`
		DELETE FROM reactions 
		WHERE id = ?`, reactionID)

	if err != nil {
		return fmt.Errorf("could not delete reaction: %w", err)
	}

	// Decrement the reaction count in the reactionCount table
	_, err = db.Exec(`
		UPDATE reactionCount 
		SET count = count - 1 
		WHERE (post_id = ? OR comment_id = ?) 
		AND reaction_type = (SELECT reaction_type FROM reactions WHERE id = ?)`,
		itemID, itemID, reactionID)

	if err != nil {
		return fmt.Errorf("could not update reaction count: %w", err)
	}

	// Remove the reaction count if it reaches zero
	_, err = db.Exec(`
		DELETE FROM reactionCount 
		WHERE (post_id = ? OR comment_id = ?) 
		AND count = 0`, itemID, itemID)

	if err != nil {
		return fmt.Errorf("could not delete reaction count: %w", err)
	}

	return nil
}
