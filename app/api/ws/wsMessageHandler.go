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
	chattingWith snowflake.SnowflakeID
}

func (conn *wsConnection) sendMessageTo(db *sql.DB, msg modules.Message) error {
	mux.Lock()
	defer mux.Unlock()
	msg.Id = snowflake.Generate()
	msg.CreationTime = time.Now()
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

func Notify[T modules.Message | modules.MessageNewUser](msg T, shouldLock bool) {
	if shouldLock {
		mux.Lock()
		defer mux.Unlock()
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

func notifyTypingStatus(msg modules.Message) {
	mux.Lock()
	defer mux.Unlock()
	userConns := activeUsers[msg.Chat]
	for _, conn := range userConns {
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Error(err)
		}
	}
}

func notifyStatusChange(userId snowflake.SnowflakeID, status string) {
	msg := modules.Message{
		Type:  "status",
		Id:    userId,
		Value: status,
	}

	Notify(msg, false)
}
