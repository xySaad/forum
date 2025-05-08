package ws

import (
	"database/sql"
	"net"

	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"

	"github.com/gorilla/websocket"
)

const WsMessageType_DM = "DM"
const WsMessageType_tying = "typing_status"

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
outer:
	for {
		msg := modules.Message{Sender: conn.User.Id}
		err := wsConn.ReadJSON(&msg)
		if err != nil {
			log.Debug(err)
			switch err.(type) {
			case *net.OpError, *websocket.CloseError:
				break outer
			default:
				continue
			}
		}

		if msg.Type == WsMessageType_DM {
			err = wsConn.sendMessageTo(forumDB, msg)
			if err != nil {
				wsConn.WriteJSON(map[string]string{
					"type":  "error",
					"value": err.Error(),
				})
				log.Error(err)
			}
		}
	}
}

func FetchMessages(conn *modules.Connection, forumDB *sql.DB) {
	if !conn.IsAuthenticated(forumDB) {
		return
	}
	if len(conn.Path) < 3 {
		conn.Error(errors.HttpNotFound)
		return
	}
	chatId := conn.Path[2]
	lastId := conn.Req.URL.Query().Get("lastId")
	query := modules.QUERY_GET_MESSAGE
	args := []any{conn.User.Id, chatId, chatId, conn.User.Id}
	if lastId != "" {
		query += "AND id < ? "
		args = append(args, lastId)
	}
	query += "ORDER BY id DESC LIMIT 10"
	rows, err := forumDB.Query(query, args...)
	if err != nil {
		conn.Error(errors.HttpInternalServerError)
		log.Debug(err)
		return
	}
	defer rows.Close()

	var messages []modules.Message
	for rows.Next() {
		msg := modules.Message{
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
