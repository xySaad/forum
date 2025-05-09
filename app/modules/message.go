package modules

import (
	"forum/app/modules/snowflake"
	"time"
)

type Message struct {
	Type         string                `json:"type"`
	Id           snowflake.SnowflakeID `json:"id"`
	Sender       snowflake.SnowflakeID `json:"sender,omitempty"`
	Chat         snowflake.SnowflakeID `json:"chat,omitempty"`
	Value        string                `json:"value"`
	CreationTime time.Time             `json:"creationTime,omitempty"`
}

type MessageNewUser struct {
	Type  string `json:"type"`
	Value *User  `json:"value"`
}

const QUERY_GET_MESSAGE = "SELECT id, sender, receiver, content, created_at FROM message WHERE ((sender = ? AND receiver = ?) OR (sender = ? AND receiver = ?)) "
