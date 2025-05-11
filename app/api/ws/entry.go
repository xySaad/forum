package ws

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
	// outer:
	// 	for {
	// 		var msg modules.Message
	// 		err := wsConn.ReadJSON(&msg)
	// 		if err != nil {
	// 			switch err.(type) {
	// 			case *net.OpError, *websocket.CloseError:
	// 				break outer
	// 			default:
	// 				continue
	// 			}
	// 		}
	// 		msg.Sender = conn.User.Id
	// 		switch msg.Type {
	// 		case modules.MessageType_DM:
	// 			temp := modules.OutgoingDM{
	// 				Chat:    msg.Chat,
	// 				Content: msg.Value,
	// 			}
	// 			err = wsConn.sendMessageTo(forumDB, temp)
	// 			if err != nil {
	// 				wsConn.WriteJSON(map[string]string{
	// 					"type":  "error",
	// 					"value": err.Error(),
	// 				})
	// 				log.Error(err)
	// 			}
	// 		case modules.MessageType_STATUS:
	// 			msg.Id = msg.Sender
	// 			wsConn.chattingWith = msg.Chat
	// 			wsConn.notifyTypingStatus(msg.Chat, msg.Value)
	// 		}
	// 	}

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
				log.Error(err)
				continue
			}
		}
		message, exist := modules.Incoming_Message_Types[rawMsg.Type]
		if !exist {
			fmt.Println("message type doesn't exist")
			continue
		}
		rawData, err := rawMsg.Data.MarshalJSON()
		if err != nil {
			//send error message
			log.Error(err)
		}
		json.Unmarshal(rawData, &message)

		switch msg := message.(type) {
		case modules.IncomingDM:
			err = wsConn.sendMessageTo(forumDB, msg)
			if err != nil {
				//send error message
				log.Error(err)
			}
		case modules.IncomingStatus:
			wsConn.chattingWith = msg.Chat
			wsConn.notifyTypingStatus(msg.Chat, msg.Status)
		default:
			fmt.Println(rawMsg.Type)
		}

	}

}
