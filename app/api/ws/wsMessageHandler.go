package ws

import (
	"database/sql"
	"fmt"
	"time"

	"forum/app/modules"
	"forum/app/modules/log"
	"forum/app/modules/snowflake"

	"github.com/gorilla/websocket"
)

type message struct {
	Type         string                `json:"type"`
	Id           snowflake.SnowflakeID `json:"id"`
	Sender       snowflake.SnowflakeID `json:"sender,omitempty"`
	Chat         snowflake.SnowflakeID `json:"chat,omitempty"`
	Value        string                `json:"value"`
	CreationTime string                `json:"creationTime,omitempty"`
}

type wsConnection struct {
	*websocket.Conn
	*modules.Connection
}

func (conn *wsConnection) sendMessageTo(db *sql.DB, msg message) error {
	msg.Id = snowflake.Generate()
	msg.CreationTime = time.Now().Format(time.DateTime)
	const query = "INSERT INTO message (id, receiver, sender, content, created_at) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, msg.Id, msg.Chat, conn.User.Id, msg.Value, msg.CreationTime)
	if err != nil {
		log.Error("Error inserting DM into database:", err)
		return err
	}
	fmt.Println(msg.Chat)
	userConns := activeUsers[msg.Chat]

	for _, conn := range append(userConns, activeUsers[conn.User.Id]...) {
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Error(err)
		}
	}

	return nil
}

func notifyStatusChange(userId snowflake.SnowflakeID, status string) {
	msg := message{
		Type:  "status",
		Id:    userId,
		Value: status,
	}

	for _, WsConnections := range activeUsers {
		for _, conn := range WsConnections {
			err := conn.WriteJSON(msg)
			if err != nil {
				log.Error(err)
			}
		}
	}
}
