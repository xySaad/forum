package modules

import (
	"encoding/json"
	"forum/app/modules/snowflake"
	"reflect"
	"time"
)

func typeStr[T any]() string {
	t := reflect.TypeFor[T]()
	return t.String()
}

const QUERY_GET_MESSAGE = "SELECT id, sender, receiver, content, created_at FROM message WHERE ((sender = ? AND receiver = ?) OR (sender = ? AND receiver = ?)) "
const (
	STATUS_ONLINE  = "ONLINE"
	STATUS_OFFLINE = "OFFLINE"
)
const ACTION_LOGOUT = "LOGOUT"
const (
	MessageType_DM       = "DM"
	MessageType_STATUS   = "STATUS"
	MessageType_Action   = "ACTION"
	MessageType_NEW_USER = "NEW_USER"
)

type OutgoingMessageData interface {
	OutgoingDM | OutgoingStatus | User | Action
}

var Outgoing_Message_Types = map[string]string{
	typeStr[OutgoingDM]():     MessageType_DM,
	typeStr[OutgoingStatus](): MessageType_STATUS,
	typeStr[User]():           MessageType_NEW_USER,
	typeStr[Action]():         MessageType_Action,
}

var Incoming_Message_Types = map[string]any{
	MessageType_DM:     IncomingDM{},
	MessageType_STATUS: IncomingStatus{},
}

type Action string

var LogoutMessage Action = ACTION_LOGOUT

type IncomingMessageData interface {
	IncomingDM | IncomingStatus
}
type IncomingMessage struct {
	Type string `json:"type"`
	Data json.RawMessage
}
type OutgoingMessage[T OutgoingMessageData] struct {
	Type string `json:"type"`
	Data *T
}

type IncomingDM struct {
	Chat    snowflake.SnowflakeID
	Content string
}
type OutgoingDM struct {
	Id           snowflake.SnowflakeID `json:"id"`
	Sender       snowflake.SnowflakeID `json:"sender"`
	Chat         snowflake.SnowflakeID `json:"chat"`
	Content      string                `json:"content"`
	CreationTime time.Time             `json:"creationTime"`
}

type IncomingStatus struct {
	Chat   snowflake.SnowflakeID
	Status string
}
type OutgoingStatus struct {
	Id     snowflake.SnowflakeID `json:"id"`
	Status string                `json:"status"`
}

func NewMessage[T OutgoingMessageData](data *T) OutgoingMessage[T] {
	return OutgoingMessage[T]{
		Type: Outgoing_Message_Types[typeStr[T]()],
		Data: data,
	}
}

// temporary
type Message struct {
	Type         string                `json:"type"`
	Id           snowflake.SnowflakeID `json:"id"`
	Sender       snowflake.SnowflakeID `json:"sender,omitempty"`
	Chat         snowflake.SnowflakeID `json:"chat,omitempty"`
	Value        string                `json:"value"`
	CreationTime time.Time             `json:"creationTime,omitempty"`
}
