package modules

import (
	"encoding/json"
	"fmt"
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
	STATUS_ONLINE  = "online"
	STATUS_OFFLINE = "offline"
)
const (
	ACTION_LOGOUT           = "logout"
	LOGOUT_REASON_NEW_LOGIN = "new_login"
)
const (
	MessageType_DM       = "DM"
	MessageType_STATUS   = "status"
	MessageType_Action   = "action"
	MessageType_NEW_USER = "new_user"
)

type Action struct {
	Action string `json:"action"`
	Reason string `json:"reason"`
}

var Outgoing_Message_Types = map[string]string{
	typeStr[OutgoingDM]():     MessageType_DM,
	typeStr[OutgoingStatus](): MessageType_STATUS,
	typeStr[User]():           MessageType_NEW_USER,
	typeStr[Action]():         MessageType_Action,
}

var incomingMessageConstructors = map[string]any{
	MessageType_DM:     &IncomingDM{},
	MessageType_STATUS: &IncomingStatus{},
}

type IncomingDM struct {
	Chat    snowflake.SnowflakeID
	Content string
}
type IncomingStatus struct {
	Chat   snowflake.SnowflakeID
	Status string
}

type IncomingMessage struct {
	Type string `json:"type"`
	Data json.RawMessage
}

func (im *IncomingMessage) Zero() any {
	ptr, ok := incomingMessageConstructors[im.Type]
	if !ok {
		return nil
	}
	return ptr
}

func (im *IncomingMessage) Decode() (any, error) {
	target := im.Zero()
	if target == nil {
		return nil, fmt.Errorf("unknown message type: %s", im.Type)
	}

	err := json.Unmarshal(im.Data, target)
	if err != nil {
		return nil, err
	}

	return target, nil
}

type OutgoingMessageData interface {
	OutgoingDM | OutgoingStatus | User | Action
}

type OutgoingMessage[T OutgoingMessageData] struct {
	Type string `json:"type"`
	Data *T     `json:"data"`
}

type OutgoingDM struct {
	Id           snowflake.SnowflakeID `json:"id"`
	Sender       snowflake.SnowflakeID `json:"sender"`
	Chat         snowflake.SnowflakeID `json:"chat"`
	Content      string                `json:"content"`
	CreationTime time.Time             `json:"creationTime"`
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
