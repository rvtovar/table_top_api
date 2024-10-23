package models

import (
	"errors"
	"table_top_api/utils"
	"time"
)

type User struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"username"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:varchar(100);not null" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type UserInfo struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

func (u *User) ValidateCreds() error {
	var user User
	result := DB.Where("email = ?", u.Email).First(&user)
	if result.Error != nil {
		return result.Error
	}

	pwdIsValid := utils.CheckPasswordHash(u.Password, user.Password)
	if !pwdIsValid {
		return errors.New("Credentials are Invalid")
	}
	u.Email = user.Email
	u.ID = user.ID
	u.Username = user.Username
	return nil
}
