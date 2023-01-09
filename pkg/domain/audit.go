package domain

import (
	"errors"
	"time"
)

const (
	EntityUser = "USER"
	EntityTest = "TEST"

	ActionCreate   = "CREATE"
	ActionGet      = "GET"
	ActionUpdate   = "UPDATE"
	ActionDelete   = "DELETE"
	ActionRegister = "REGISTER"
	ActionLogin    = "LOGIN"
)

var (
	entities = map[string]LogRequest_Entities{
		EntityUser: LogRequest_USER,
		EntityTest: LogRequest_TEST,
	}

	actions = map[string]LogRequest_Actions{
		ActionCreate:   LogRequest_CREATE,
		ActionUpdate:   LogRequest_UPDATE,
		ActionGet:      LogRequest_GET,
		ActionDelete:   LogRequest_DELETE,
		ActionRegister: LogRequest_REGISTER,
		ActionLogin:    LogRequest_LOGIN,
	}
	ErrorInvalidEntity = errors.New("invalid entity")
	ErrorInvalidAction = errors.New("invalid action")
)

type LogItem struct {
	Entity    string    `bson:"entity"`
	Action    string    `bson:"action"`
	EntityID  int64     `bson:"entity_id"`
	Timestamp time.Time `bson:"timestamp"`
}

// ToPbEntity - converter
func ToPbEntity(entity string) (LogRequest_Entities, error) {
	val, ex := entities[entity]
	if !ex {
		return 0, ErrorInvalidEntity
	}

	return val, nil
}

// ToPbAction - converter
func ToPbAction(action string) (LogRequest_Actions, error) {
	val, ex := actions[action]
	if !ex {
		return 0, ErrorInvalidAction
	}

	return val, nil
}
