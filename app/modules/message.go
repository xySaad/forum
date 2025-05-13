package modules

import (
	"encoding/json"
	"fmt"
	"forum/app/modules/snowflake"
	"time"
)

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

type Message interface {
	OutgoingDM | OutgoingStatus | User | Action
	MessageType() string
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

type OutgoingMessage[T Message] struct {
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

func (o OutgoingDM) MessageType() string { return MessageType_DM }

type OutgoingStatus struct {
	Id     snowflake.SnowflakeID `json:"id"`
	Status string                `json:"status"`
}

func (o OutgoingStatus) MessageType() string { return MessageType_STATUS }

type Action struct {
	Action string `json:"action"`
	Reason string `json:"reason"`
}

func (o Action) MessageType() string { return MessageType_Action }

func NewMessage[T Message](data *T) OutgoingMessage[T] {
	return OutgoingMessage[T]{
		Type: (*data).MessageType(),
		Data: data,
	}
}
