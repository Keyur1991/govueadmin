package models

import (
	"app/microservices/auth/config"
	"time"
)

type AuthToken struct {
	UserId    string    `json:"user_id" gorm:"type:varchar(255) NOT NULL"`
	TokenId   string    `json:"token_id" gorm:"type:varchar(255) NOT NULL"`
	CreatedAt time.Time `json:"created_at"`
}

func (authToken *AuthToken) Create() error {
	if err := config.DB.Debug().Create(authToken).Error; err != nil {
		return err
	}
	return nil
}

func (authToken *AuthToken) Delete(uid string, token_id string) error {
	if err := config.DB.Where("user_id = ?", uid).Where("token_id", token_id).Delete(authToken).Error; err != nil {
		return err
	}
	return nil
}

func (authToken *AuthToken) Find(uid string, token_id string) error {
	if err := config.DB.Where("user_id = ?", uid).Where("token_id", token_id).First(authToken).Error; err != nil {
		return err
	}

	return nil
}
