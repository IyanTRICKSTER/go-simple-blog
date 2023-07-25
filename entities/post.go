package entities

import "time"

type Post struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `json:"userID"`
	Title     string `gorm:"not null" json:"title"`
	Slug      string `gorm:"not null" json:"slug"`
	Content   string `gorm:"not null;type:text;longtext" json:"content"`
	Image     string `gorm:"not null;type:text;text" json:"image"`
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli" json:"updatedAt"`
	DeletedAt *int64 `json:"deletedAt"`
}

func AssignDeleteTime() *int64 {
	now := time.Now().UnixMilli()
	return &now
}
