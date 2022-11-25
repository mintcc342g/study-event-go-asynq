package dto

import "time"

type Announcement struct {
	ID       uint64    `json:"id"`
	From     string    `json:"from"`
	Message  string    `json:"message"`
	Seconds  int64     `json:"seconds"`
	Deadline time.Time `json:"deadline"`
}
