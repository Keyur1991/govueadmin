package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host string
	Port string
	User string
	Password string
	Database string
}

func DBUrl(config *Config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
}

func Connect(config *Config) (*gorm.DB, error) {
	dsn := DBUrl(config)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
