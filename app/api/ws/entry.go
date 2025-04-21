package ws

import (
	"database/sql"
	"forum/app/modules"
	"forum/app/modules/log"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Entry(conn *modules.Connection, forumDB *sql.DB) {
	if !conn.IsAuthenticated(forumDB) {
		return
	}

	wsConn, err := upgrader.Upgrade(conn.Resp, conn.Req, nil)
	if err != nil {
		log.Error(err)
		return
	}
	defer wsConn.Close()
	defer notifyStatusChange(conn.User.Id, "offline")
	addActiveUser(conn.User.Id, wsConn)

	for {
		var msg wsDirectMessage
		err := wsConn.ReadJSON(&msg)
		if err == nil {
			handleWsMessage(msg, conn.User.Id)
			continue
		}

		log.Error(err)
		_, ok := err.(*websocket.CloseError)
		if ok {
			return
		}
	}
}
