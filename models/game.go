package models

import "time"

type Game struct {
	ID       int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string    `gorm:"type:varchar(100);not null" json:"name"`
	Style    string    `gorm:"type:varchar(100);not null" json:"style"`
	Location string    `gorm:"type:varchar(100);not null" json:"location"`
	DateTime time.Time `gorm:"type:datetime;not null" json:"date_time""`
	UserId   int64     `gorm:"not null" json:"user_id"`
	User     User      `gorm:"foreignKey:UserId" json:"user"`
}
