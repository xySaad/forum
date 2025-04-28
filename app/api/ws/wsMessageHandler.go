package ws

import (
	"fmt"

	"forum/app/modules/log"
	"forum/app/modules/snowflake"
)

type mesage struct {
	Type     string `json:"type"`
	Id       string `json:"id"`
	Receiver string `json:"receiver"`
	Msg      string `json:"msg"`
}
type wsNotifyMsg struct {
	Type   string                `json:"type"`
	Id     snowflake.SnowflakeID `json:"id"`
	Status string                `json:"status"`
}

func sendMessages(msg mesage) {
	 user := activeUsers[msg.Receiver]
		for _, conn := range user {
			err := conn.WriteJSON(msg)
			if err != nil {
				log.Error(err)
			}
		}

}

func notifyStatusChange(userId snowflake.SnowflakeID, status string) {
	msg := wsNotifyMsg{
		Type:   "status",
		Id:     userId,
		Status: status,
	}
	
	for _, WsConnections := range activeUsers {
		fmt.Print(WsConnections)
		for _, conn := range WsConnections {
			err := conn.WriteJSON(msg)
			if err != nil {
				log.Error(err)
			}
		}
	}
}
