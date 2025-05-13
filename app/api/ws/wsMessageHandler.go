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
	ChattingWith snowflake.SnowflakeID
}

func (conn *wsConnection) sendMessageTo(db *sql.DB, inMsg *modules.IncomingDM) error {
	mux.Lock()
	defer mux.Unlock()
	msg := modules.OutgoingDM{
		Id:           snowflake.Generate(),
		Sender:       conn.User.Id,
		Chat:         inMsg.Chat,
		Content:      inMsg.Content,
		CreationTime: time.Now(),
	}

	const query = "INSERT INTO message (id, receiver, sender, content, created_at) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, msg.Id, msg.Chat, conn.User.Id, msg.Content, msg.CreationTime)
	if err != nil {
		log.Error("Error inserting DM into database:", err)
		return err
	}
	userConns := activeUsers[msg.Chat]
	if msg.Chat != conn.User.Id {
		userConns = append(userConns, activeUsers[conn.User.Id]...)
	}

	for _, conn := range userConns {
		err := conn.WriteJSON(modules.NewMessage(&msg))
		if err != nil {
			log.Error(err)
		}
	}

	return nil
}

type Notifyable interface {
	modules.Message
	modules.OutgoingStatus | modules.User
}

func Notify[T Notifyable](msg modules.OutgoingMessage[T], shouldLock bool) {
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

func (ownConn *wsConnection) notifyTypingStatus(to snowflake.SnowflakeID, status string) {
	msg := modules.OutgoingStatus{
		Id:     ownConn.User.Id,
		Status: status,
	}

	mux.Lock()
	defer mux.Unlock()
	userConns := activeUsers[to]
	for _, conn := range userConns {
		if conn == ownConn {
			continue
		}
		err := conn.WriteJSON(modules.NewMessage(&msg))
		if err != nil {
			log.Error(err)
		}
	}
}

func notifyStatusChange(userId snowflake.SnowflakeID, status string) {
	msg := modules.OutgoingStatus{
		Id:     userId,
		Status: status,
	}

	Notify(modules.NewMessage(&msg), false)
}
