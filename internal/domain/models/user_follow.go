package models

import "time"

type Follow struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"not null;index:idx_follow,unique"`
	FollowerID uint      `gorm:"not null;index:idx_follow,unique"`
	CreatedAt  time.Time `gorm:"index"`
}
