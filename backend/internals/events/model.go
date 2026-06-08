package events

import "time"

type Event struct {
	Id          int64
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Datetime    time.Time
	Location    string
	UserId      string
}
