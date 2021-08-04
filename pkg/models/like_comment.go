package models

import (
	uuid "github.com/satori/go.uuid"
)

type LikeComment struct {
	UUID    uuid.UUID `json:"uuid" db:"uuid"`
	Comment *Comment  `json:"article"`
	User    *User     `json:"user"`
}
