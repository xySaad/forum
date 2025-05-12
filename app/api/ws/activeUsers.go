package ws

import (
	"forum/app/modules"
	"forum/app/modules/snowflake"
	"slices"
	"sync"
)

type activeUser struct {
	conns      []*wsConnection
	typingFrom []snowflake.SnowflakeID
}

var activeUsers = map[snowflake.SnowflakeID][]*wsConnection{}
var mux sync.Mutex

func addActiveUser(conn *wsConnection) {
	mux.Lock()
	defer mux.Unlock()
	activeUsers[conn.User.Id] = append(activeUsers[conn.User.Id], conn)
	notifyStatusChange(conn.User.Id, "online")
	for _, chattingWith := range modules.Typingto[conn.User.Id] {
		go notifyTypingStatus(modules.Message{
			Type:  WsMessageType_STATUS,
			Id:    chattingWith,
			Chat:  conn.User.Id,
			Value: "typing",
		  })
	
}
}
func deleteActiveUser(conn *wsConnection) {
	notifyTypingStatus(modules.Message{
		Type:  WsMessageType_STATUS,
		Id:    conn.User.Id,
		Chat:  conn.chattingWith,
		Value: "afk",
	})
	mux.Lock()
	defer mux.Unlock()
	userConns, exist := activeUsers[conn.User.Id]
	if !exist {
		return
	}
	connIdx := slices.Index(userConns, conn)
	activeUsers[conn.User.Id] = slices.Delete(userConns, connIdx, connIdx+1)
	if len(activeUsers[conn.User.Id]) == 0 {
		delete(activeUsers, conn.User.Id)
		notifyStatusChange(conn.User.Id, "offline")
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
