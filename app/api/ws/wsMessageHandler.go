package ws

import (
	"forum/app/modules/log"
	"forum/app/modules/snowflake"
)

type wsDirectMessage struct {
	Type     string `json:"type"`
	Receiver string `json:"receiver"`
	Text     string `json:"text"`
}

func handleWsMessage(msg wsDirectMessage, userId snowflake.SnowflakeID) {
	for _, conn := range activeUsers[userId] {
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Error(err)
		}
	}
}

type wsNotifyMsg struct {
	Type   string                `json:"type"`
	Id     snowflake.SnowflakeID `json:"id"`
	Status string                `json:"status"`
}

func notifyStatusChange(userId snowflake.SnowflakeID, status string) {
	msg := wsNotifyMsg{
		Type:   "status",
		Id:     userId,
		Status: status,
	}

	for _, WsConnections := range activeUsers {
		for _, conn := range WsConnections {
			err := conn.WriteJSON(msg)
			if err != nil {
				log.Error(err)
			}
		}
	}
}
