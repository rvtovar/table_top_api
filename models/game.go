package models

import "time"

type Game struct {
	ID            int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string         `gorm:"type:varchar(100);not null" json:"name"`
	Style         string         `gorm:"type:varchar(100);not null" json:"style"`
	Location      string         `gorm:"type:varchar(100);not null" json:"location"`
	DateTime      time.Time      `gorm:"type:datetime;not null" json:"date_time""`
	UserId        int64          `gorm:"not null" json:"user_id"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	Registrations []Registration `gorm:"foreignKey:GameID" json:"registrations"`
}

func (g *Game) Register(userId int64) error {
	registration := Registration{
		GameID: g.ID,
		UserID: userId,
	}
	result := DB.Create(&registration)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (g *Game) CancelRegistration(id int64) error {
	result := DB.Where("game_id = ? AND user_id = ?", g.ID, id).Delete(&Registration{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
