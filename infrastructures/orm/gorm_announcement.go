package orm

import "time"

type Announcement struct {
	ID           uint64 `gorm:"primaryKey"`
	From         string
	Message      string
	Timeout      time.Duration `gorm:"default:null"`
	Deadline     time.Time     `gorm:"default:null"`
	UpdatedCount uint64        `gorm:"default:0"` // for async test
}
