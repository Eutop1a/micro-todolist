package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// 定义model，就是数据库模型
type User struct {
	gorm.Model
	UserName string `gorm:"unique"`
	Password string
}

const PasswordCost = 12

func (user *User) SetPassword(password string) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return
	}
	user.Password = string(bytes)
	return
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
