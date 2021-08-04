package models

import (
	"forRoma/pkg/custom_types"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Article struct {
	UUID      uuid.UUID             `json:"uuid" db:"uuid"`
	User      *User                 `json:"user"`
	Title     string                `json:"title" db:"title"`
	Text      string                `json:"text" db:"text"`
	Comments  []*Comment            `json:"comments"`
	CreateAt  time.Time             `json:"create_at" db:"create_at"`
	UpdatedAt time.Time             `json:"updated_at" db:"updated_at"`
	DeletedAt custom_types.NullTime `json:"deleted_at" db:"deleted_at"`
}
