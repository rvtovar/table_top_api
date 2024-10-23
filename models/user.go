package models

import (
	"errors"
	"table_top_api/utils"
	"time"
)

type User struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar(100);unqiueIndex;not null" json:"username"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:varchar(100);not null" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
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

	return nil
}
