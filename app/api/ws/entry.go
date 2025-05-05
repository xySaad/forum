package ws

import (
	"database/sql"

	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"

	"github.com/gorilla/websocket"
)

const WsMessageType_DM = "DM"

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
	wsConn := &wsConnection{gorillaWsConn, conn}
	defer wsConn.Close()
	defer deleteActiveUser(conn.User.Id, wsConn)
	addActiveUser(conn.User.Id, wsConn)

	for {
		var msg message
		err := wsConn.ReadJSON(&msg)
		if err != nil {
			log.Error(err)
			if _, ok := err.(*websocket.CloseError); ok {
				break
			} else {
				continue
			}
		}

		switch msg.Type {
		case WsMessageType_DM:
			err = wsConn.sendMessageTo(forumDB, msg)
			if err != nil {
				wsConn.WriteJSON(map[string]string{
					"type":  "error",
					"value": err.Error(),
				})
				log.Error(err)
			}
		default:
			//Bad Request
		}
	}
}

func FetchMessages(conn *modules.Connection, forumDB *sql.DB) {
	if !conn.IsAuthenticated(forumDB) {
		return
	}
	chatId := conn.Path[2]
	query := `SELECT id, sender, receiver, content, created_at FROM message WHERE (sender = ? AND receiver = ?) OR (sender = ? AND receiver = ?)`
	rows, err := forumDB.Query(query, conn.User.Id, chatId, chatId, conn.User.Id)
	if err != nil {
		conn.Error(errors.HttpInternalServerError)
		log.Error(err)
		return
	}
	defer rows.Close()

	var messages []message
	for rows.Next() {
		msg := message{
			Type: WsMessageType_DM,
		}
		if err := rows.Scan(&msg.Id, &msg.Sender, &msg.Chat, &msg.Value, &msg.CreationTime); err != nil {
			log.Error(err)
			continue
		}
		messages = append(messages, msg)
	}

	conn.Respond(messages)
}
