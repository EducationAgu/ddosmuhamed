package entity

import "time"

type Session struct {
	UserId    int64
	Token     string
	ExpiresAt time.Time
}
