package models

import (
	uuid "github.com/satori/go.uuid"
)

type LikeArticle struct {
	UUID    uuid.UUID `json:"uuid" db:"uuid"`
	Article *Article  `json:"article"`
	User    *User     `json:"user"`
}
