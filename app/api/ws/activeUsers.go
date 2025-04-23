package ws

import (
	"forum/app/modules/snowflake"
	"sync"

	"github.com/gorilla/websocket"
)

var activeUsers = map[snowflake.SnowflakeID][]*websocket.Conn{}
var mux sync.Mutex

func addActiveUser(userId snowflake.SnowflakeID, conn *websocket.Conn) {
	mux.Lock()
	defer mux.Unlock()
	activeUsers[userId] = append(activeUsers[userId], conn)
	notifyStatusChange(userId, "online")
}
func deleteActiveUser(userId snowflake.SnowflakeID) {
	mux.Lock()
	defer mux.Unlock()
	delete(activeUsers, userId)
	notifyStatusChange(userId, "offline")
}

func IsActive(userId snowflake.SnowflakeID) bool {
	_, exist := activeUsers[userId]
	return exist
}
