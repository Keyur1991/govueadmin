package models

import (
	"govueadmin/microservices/auth/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Uid        string      `json:"uid" gorm:"primaryKey;type:varchar(255) NOT NULL"`
	FirstName  string      `json:"first_name" gorm:"type:varchar(255) NOT NULL"`
	LastName   string      `json:"last_name" gorm:"type:varchar(255) NOT NULL"`
	Email      string      `json:"email" gorm:"type:varchar(255) NOT NULL"`
	Password   string      `json:"-" gorm:"type:varchar(255) NOT NULL"`
	Phone      string      `json:"phone" gorm:"type:varchar(20);default:null"`
	Address    string      `json:"address" gorm:"type:text;default:null"`
	IsAdmin    string      `json:"is_admin" gorm:"type:enum('0','1');default:'0'"`
	IsVerified string      `json:"is_verified" gorm:"type:enum('0', '1'); default:'0'"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	DeletedAt  interface{} `json:"-" gorm:"-"`
}

func (user *User) GetByEmail(email string) (err error) {
	if err := config.DB.Debug().Where("email = ?", email).First(user).Error; err != nil {
		return err
	}

	return nil
}

func (user *User) Find(id string) (err error) {
	if err := config.DB.Where("uid = ?", id).First(user).Error; err != nil {
		return err
	}

	return nil
}
