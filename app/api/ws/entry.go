package ws

import (
	"database/sql"
	"fmt"
	"net"
	"reflect"

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

	gorillaWsConn, err := upgrader.Upgrade(conn.Resp, conn.Req, nil)
	if err != nil {
		log.Error(err)
		return
	}
	wsConn := &wsConnection{gorillaWsConn, conn, 0}
	defer wsConn.Close()
	defer deleteActiveUser(wsConn)
	addActiveUser(wsConn)
	// TODO: prevent user from sending custom status
outer:
	for {
		var rawMsg modules.IncomingMessage
		err := wsConn.ReadJSON(&rawMsg)
		if err != nil {
			switch err.(type) {
			case *net.OpError, *websocket.CloseError:
				break outer
			default:
				log.Error(err)
				continue
			}
		}
		v, err := rawMsg.Decode()
		if err != nil {
			log.Error(err)
			continue
		}
		switch msg := v.(type) {
		case *modules.IncomingDM:
			err = wsConn.sendMessageTo(forumDB, msg)
			if err != nil {
				log.Fatal(err)
			}
		case *modules.IncomingStatus:
			if msg.Status == "typing" {
				wsConn.ChattingWith = msg.Chat
			} else {
				wsConn.ChattingWith = 0
			}
			wsConn.notifyTypingStatus(msg.Chat, msg.Status)
		default:
			fmt.Println("invalid asserted type", reflect.TypeOf(v))
		}
	}
}
