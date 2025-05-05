package ws

import (
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
	userConns := activeUsers[userId]
	connIdx := slices.Index(userConns, conn)
	activeUsers[userId] = slices.Delete(userConns, connIdx, connIdx+1)
	if len(activeUsers[userId]) == 0 {
		delete(activeUsers, userId)
		notifyStatusChange(userId, "offline")
	}
}

func IsActive(userId snowflake.SnowflakeID) bool {
	_, exist := activeUsers[userId]
	return exist
}
