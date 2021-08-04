package models

import (
	"forRoma/pkg/custom_types"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Comment struct {
	UUID      uuid.UUID             `json:"uuid" db:"uuid"`
	User      *User                 `json:"user"`
	Text      string                `json:"text" db:"text"`
	CreateAt  time.Time             `json:"create_at" db:"create_at"`
	UpdatedAt time.Time             `json:"updated_at" db:"updated_at"`
	DeletedAt custom_types.NullTime `json:"deleted_at" db:"deleted_at"`
}
