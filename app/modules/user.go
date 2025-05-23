package modules

import (
	"database/sql"
	"forum/app/modules/snowflake"
)

type User struct {
	Username       string                `json:"username"`
	Id             snowflake.SnowflakeID `json:"id"`
	ProfilePicture *string               `json:"profilePicture"`
	Status         string                `json:"status"`
	LastMessage    *OutgoingDM           `json:"lastMessage,omitempty"`
}

func (o User) MessageType() string { return MessageType_NEW_USER }

func (u *User) GetPublicUser(db *sql.DB) (err error) {
	qreury := `SELECT id,username,profile_picture FROM users WHERE id=?`
	err = db.QueryRow(qreury, u.Id).Scan(&u.Id, &u.Username, &u.ProfilePicture)
	return
}
