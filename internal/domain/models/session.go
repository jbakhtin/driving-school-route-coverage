package models

import (
	"encoding/json"
	"time"
)

type Session struct {
	SessionId string
	UserId    int
	TTL       time.Duration
}

func (e *Session) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}

	return marshal
}
