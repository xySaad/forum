package ws

import (
	"forum/app/modules"
	"forum/app/modules/snowflake"
	"slices"
	"sync"
)

var activeUsers = map[snowflake.SnowflakeID][]*wsConnection{}
var mux sync.Mutex

func addActiveUser(userId snowflake.SnowflakeID, conn *wsConnection) {
	mux.Lock()
	defer mux.Unlock()
	activeUsers[userId] = append(activeUsers[userId], conn)
	notifyStatusChange(userId, "online")
}
func deleteActiveUser(userId snowflake.SnowflakeID, conn *wsConnection) {
	mux.Lock()
	defer mux.Unlock()
	userConns, exist := activeUsers[userId]
	if !exist {
		return
	}
	connIdx := slices.Index(userConns, conn)
	activeUsers[userId] = slices.Delete(userConns, connIdx, connIdx+1)
	if len(activeUsers[userId]) == 0 {
		delete(activeUsers, userId)
		notifyStatusChange(userId, "offline")
	}
}
func ExpireAll(userId snowflake.SnowflakeID) {
	mux.Lock()
	defer mux.Unlock()
	for _, conn := range activeUsers[userId] {
		conn.WriteJSON(modules.Message{Type: "logout"})
		conn.Close()
	}
	delete(activeUsers, userId)
	notifyStatusChange(userId, "offline")
}

func IsActive(userId snowflake.SnowflakeID) bool {
	mux.Lock()
	defer mux.Unlock()
	_, exist := activeUsers[userId]
	return exist
}
