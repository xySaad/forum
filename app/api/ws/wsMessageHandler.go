package ws

import (
	"forum/app/modules/log"
	"forum/app/modules/snowflake"
)
type mesage struct {
	Id       string `json:"id"`
	Receiver string `json:"receiver"`
	Msg      string `json:"msg"`
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
