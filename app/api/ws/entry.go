package ws

import (
	"database/sql"
	"fmt"

	"forum/app/modules"
	"forum/app/modules/log"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Mssage struct {
	Id       int
	Receiver string
	Msg      string
}

func Message(message mesage, forumDB *sql.DB) {
	//var message Msg
	//err := json.NewDecoder(conn.Req.Body).Decode(&message)
	//
	fmt.Println("Received message:", message)

	_, err := forumDB.Exec("INSERT INTO message (id, receiver, msg) VALUES (?, ?, ?)", message.Id, message.Receiver, message.Msg)
	if err != nil {
		fmt.Println("Error inserting into database:", err)
		return
	}

	// fmt.Println("Message inserted successfully.")
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
	defer deleteActiveUser(conn.User.Id, wsConn)
	addActiveUser(conn.User.Id, wsConn)

	for {
		var msg mesage
		err := wsConn.ReadJSON(&msg)
		if err == nil {
			Message(msg, forumDB)
			sendMessages(msg)
			continue
		}

		log.Error(err)
		_, ok := err.(*websocket.CloseError)
		if ok {
			return
		}
	}
}

func FetchMessages(msg mesage, forumDB *sql.DB) []Mssage {
	query := `SELECT id, receiver, msg FROM message WHERE id = ? AND receiver = ? OR id = ? AND receiver = ?`
	rows, err := forumDB.Query(query, msg.Id, msg.Receiver, msg.Id, msg.Receiver)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var messages []Mssage
	for rows.Next() {
		var msg Mssage
		if err := rows.Scan(&msg.Id, &msg.Receiver, &msg.Msg); err != nil {
			fmt.Println(err)
		}
		fmt.Println(msg)
		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(messages)
	return messages
}
