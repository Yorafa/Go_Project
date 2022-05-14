package entity

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(name string, password string) error {
	var err error
	dsn := name + ":" + password + "@tcp(127.0.0.1:3306)/community?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
