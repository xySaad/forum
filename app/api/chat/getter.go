package chat

import (
	"database/sql"
	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"
)

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

	var messages []modules.OutgoingMessage[modules.OutgoingDM]
	for rows.Next() {
		var msg modules.OutgoingDM
		if err := rows.Scan(&msg.Id, &msg.Sender, &msg.Chat, &msg.Content, &msg.CreationTime); err != nil {
			log.Error(err)
			continue
		}
		messages = append(messages, modules.NewMessage(&msg))
	}

	conn.Respond(messages)
}
