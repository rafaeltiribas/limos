package models

import "time"

type User struct {
	UserID       int64     `json:"user_id"`
	RegisterDate time.Time `json:"start_date"`
	LastUseDate  time.Time `json:"last_date"`
}
