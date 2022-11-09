package models

import (
	"govueadmin/framework/request"
	"govueadmin/microservices/users/config"
	"time"

	"github.com/go-playground/validator/v10"
)

type Role struct {
	Uid       string    `json:"uid" gorm:"primaryKey;type:varchar(255) NOT NULL"`
	Name      string    `json:"name" gorm:"type:varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL" binding:"required"`
	IsActive  string    `json:"is_active" gorm:"type:enum('0', '1') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '1'"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-" gorm:"-"`
}

func (role *Role) ValidationMessages(err validator.ValidationErrors) map[string][]string {
	Fields := map[string]string{
		"Role.Name": "name",
	}

	Messages := map[string]string{
		"Roles.Name.required": "Please provide the name",
	}

	return request.ValidationMessages(&err, &Fields, &Messages)
}

func (role *Role) GetRoles() (roles []Role, err error) {
	if err = config.DB.Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

func (role *Role) Create() error {
	if err := config.DB.Create(role).Error; err != nil {
		return err
	}

	return nil
}

func (role *Role) Update(uid string) error {
	if err := config.DB.Where("Uid = ?", uid).Save(role).Error; err != nil {
		return err
	}
	return nil
}

func (role *Role) Delete(uid string) error {
	if err := config.DB.Where("Uid = ?", uid).Delete(role).Error; err != nil {
		return err
	}
	return nil
}
