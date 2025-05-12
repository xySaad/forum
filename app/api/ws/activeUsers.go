package ws

import (
	"forum/app/modules"
	"forum/app/modules/snowflake"
	"slices"
	"sync"
)

var Typingto = map[snowflake.SnowflakeID][]snowflake.SnowflakeID{}
var typingMux sync.Mutex

func updateTyping(msg *modules.IncomingStatus, id snowflake.SnowflakeID) {
	typingMux.Lock()
	typingTo := Typingto[msg.Chat]
	if msg.Status == "typing" {
		Typingto[msg.Chat] = append(typingTo, id)
	} else {
		idIdx := slices.Index(typingTo, id)
		Typingto[msg.Chat] = slices.Delete(typingTo, idIdx, idIdx+1)
	}
	typingMux.Unlock()
}

var activeUsers = map[snowflake.SnowflakeID][]*wsConnection{}
var mux sync.Mutex

func addActiveUser(conn *wsConnection) {
	mux.Lock()
	defer mux.Unlock()
	activeUsers[conn.User.Id] = append(activeUsers[conn.User.Id], conn)
	notifyStatusChange(conn.User.Id, "online")
}
func deleteActiveUser(conn *wsConnection) {
	conn.notifyTypingStatus(conn.chattingWith, "afk")
	updateTyping(&modules.IncomingStatus{
		Status: "afk",
		Chat:   conn.chattingWith,
	}, conn.User.Id)

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
	logoutAction := modules.Action{
		Action: modules.ACTION_LOGOUT,
		Reason: modules.LOGOUT_REASON_NEW_LOGIN,
	}
	for _, conn := range activeUsers[userId] {
		conn.WriteJSON(modules.NewMessage(&logoutAction))
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
