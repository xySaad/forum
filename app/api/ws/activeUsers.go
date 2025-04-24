package ws

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"

	"forum/app/modules/snowflake"

	"github.com/gorilla/websocket"
)

var (
	activeUsers = map[snowflake.SnowflakeID][]*websocket.Conn{}
	mux         sync.Mutex
)

func addActiveUser(userId snowflake.SnowflakeID, conn *websocket.Conn) {
	mux.Lock()
	activeUsers[userId] = append(activeUsers[userId], conn)
	notifyStatusChange(userId, "online")
	mux.Unlock()
}

func (msg wsDirectMessage) CreatMsgInDb(db *sql.DB, resp http.ResponseWriter) error {
	_, err := db.Exec("INSERT INTO users (id, receiver, msg) VALUES (? ,? ,?)", msg.Type, msg.Receiver, msg.Text)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
