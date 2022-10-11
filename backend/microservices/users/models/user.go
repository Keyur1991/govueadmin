package models

import (
	"fmt"
	"govueadmin/framework/request"
	"govueadmin/microservices/users/config"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type User struct {
	Uid        string      `json:"uid" gorm:"primaryKey;type:varchar(255) NOT NULL"`
	FirstName  string      `json:"first_name" gorm:"type:varchar(255) NOT NULL" binding:"required"`
	LastName   string      `json:"last_name" gorm:"type:varchar(255) NOT NULL" binding:"required"`
	Email      string      `json:"email" gorm:"type:varchar(255) NOT NULL" binding:"required_without=Id,email,UniqueField"`
	Password   string      `json:"-" gorm:"type:varchar(255) NOT NULL" binding:"required_without=Id"`
	Phone      string      `json:"phone" gorm:"type:varchar(20);default:null"`
	Address    string      `json:"address" gorm:"type:text;default:null"`
	IsAdmin    string      `json:"is_admin" gorm:"type:enum('0','1');default:'0'"`
	IsVerified string      `json:"is_verified" gorm:"type:enum('0', '1'); default:'0'"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	DeletedAt  interface{} `json:"-" gorm:"-"`
}

func (user *User) IsImage(fl validator.FieldLevel) bool {
	return true
}

func (user *User) ImageMaxSizeAllowed(fl validator.FieldLevel) bool {
	return true
}

func (user *User) UniqueField(fl validator.FieldLevel) bool {
	db := config.DB.Debug().Where("email = ?", (*user).Email)

	fmt.Println((*user))
	if (*user).Uid != "" {
		db = db.Where("uid != ?", (*user).Uid)
	}

	result := db.Find(&User{})

	return result.RowsAffected == 0
}

func (user *User) PwdPattern(fl validator.FieldLevel) bool {
	m, _ := regexp.Match(`^$`, []byte((*user).Password))
	fmt.Println("Password matched ", m)
	return m
}

func (user *User) ValidationMessages(err validator.ValidationErrors) map[string][]string {
	Fields := map[string]string{
		"User.FirstName": "first name",
		"User.LastName":  "last name",
		"User.Email":     "email",
		"User.Password":  "password",
	}

	Messages := map[string]string{
		"User.FirstName.required":        "Please provide the first name",
		"User.LastName.required":         "Please provide the last name",
		"User.Email.required":            "Please provide the email",
		"User.Email.required_without":    "Please provide the email",
		"User.Email.email":               "Please provide valid email address",
		"User.Email.UniqueField":         "The account already exist with the email address",
		"User.Password.required_without": "Please provide the password.",
	}

	return request.ValidationMessages(&err, &Fields, &Messages)
}

func (user *User) GetUsers() (users []User, err error) {
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (user *User) Create() (err error) {
	uid, _ := uuid.NewRandom()

	user.Uid = uid.String()

	if err := config.DB.Debug().Create(user).Error; err != nil {
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

func (user *User) GetByEmail(email string) (err error) {
	if err := config.DB.Debug().Where("email = ?", email).First(user).Error; err != nil {
		return err
	}

	return nil
}

func (user *User) GetByIdOrEmail(idOrEmail string) (err error) {
	if err := config.DB.Debug().Where("email = ?", idOrEmail).Or("uid = ?", idOrEmail).First(user).Error; err != nil {
		return err
	}

	return nil
}

func (user *User) Update(id string) (err error) {
	if err := config.DB.Debug().Where("uid = ?", id).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) Delete(id string) (err error) {
	if err := config.DB.Where("uid = ?", id).Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) GetByHashEmail(hash string) error {
	if err := config.DB.Debug().Where("MD5(email) = ?", hash).First(user).Error; err != nil {
		return err
	}
	return nil
}
