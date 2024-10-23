package models

import "time"

type Registration struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	GameID    int64     `gorm:"not null" json:"game_id"`
	UserID    int64     `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	//User      UserInfo  `gorm:"foreignKey:UserID;references:ID" json:"user"`
}
