package ws

import (
	"database/sql"
	"time"

	"forum/app/modules"
	"forum/app/modules/log"
	"forum/app/modules/snowflake"

	"github.com/gorilla/websocket"
)

type wsConnection struct {
	*websocket.Conn
	*modules.Connection
}

func (conn *wsConnection) sendMessageTo(db *sql.DB, msg modules.Message) error {
	msg.Id = snowflake.Generate()
	msg.CreationTime = time.Now().Format(time.DateTime)
	const query = "INSERT INTO message (id, receiver, sender, content, created_at) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, msg.Id, msg.Chat, conn.User.Id, msg.Value, msg.CreationTime)
	if err != nil {
		log.Error("Error inserting DM into database:", err)
		return err
	}
	userConns := activeUsers[msg.Chat]
	if msg.Chat != conn.User.Id {
		userConns = append(userConns, activeUsers[conn.User.Id]...)
	}

	for _, conn := range userConns {
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Error(err)
		}
	}

	return nil
}

func notifyStatusChange(userId snowflake.SnowflakeID, status string) {
	msg := modules.Message{
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
