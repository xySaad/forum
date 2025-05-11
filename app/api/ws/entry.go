package ws

import (
	"database/sql"
	"encoding/json"
	"net"

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
		var msg modules.Message
		err := wsConn.ReadJSON(&msg)
		if err != nil {
			switch err.(type) {
			case *net.OpError, *websocket.CloseError:
				break outer
			default:
				continue
			}
		}
		msg.Sender = conn.User.Id
		switch msg.Type {
		case modules.MessageType_DM:
			temp := modules.OutgoingDM{
				Chat:    msg.Chat,
				Content: msg.Value,
			}
			err = wsConn.sendMessageTo(forumDB, temp)
			if err != nil {
				wsConn.WriteJSON(map[string]string{
					"type":  "error",
					"value": err.Error(),
				})
				log.Error(err)
			}
		case modules.MessageType_STATUS:
			msg.Id = msg.Sender
			wsConn.chattingWith = msg.Chat
			wsConn.notifyTypingStatus(msg.Chat, msg.Value)
		}
	}

	//duplicate proccess for debugging
debug:
	for {
		var rawMsg modules.IncomingMessage
		err := wsConn.ReadJSON(&rawMsg)
		if err != nil {
			switch err.(type) {
			case *net.OpError, *websocket.CloseError:
				break debug
			default:
				continue
			}
		}
		msg, exist := modules.Incoming_Message_Types[rawMsg.Type]
		if !exist {
			continue
		}
		rawData, err := rawMsg.Data.MarshalJSON()
		json.Unmarshal(rawData, &msg)

		switch v := msg.(type) {
		case modules.IncomingDM:
			temp := modules.OutgoingDM{
				Chat:    msg.Chat,
				Content: msg.Value,
			}
			err = wsConn.sendMessageTo(forumDB, temp)
			if err != nil {
				wsConn.WriteJSON(map[string]string{
					"type":  "error",
					"value": err.Error(),
				})
				log.Error(err)
			}
		case modules.IncomingStatus:
			msg.Id = conn.User.Id
			wsConn.chattingWith = msg.Chat
			wsConn.notifyTypingStatus(msg.Chat, msg.Value)
		}
	}

}
