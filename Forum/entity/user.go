package entity

import (
	"gorm.io/gorm"
	"sync"
	"time"
)

type User struct {
	ID         int       `gorm:"column:id"`
	Name       string    `gorm:"column:Name"`
	Password   string    `gorm:"column:password"`
	CreateTime time.Time `gorm:"column:create_time"`
	ModifyTime time.Time `gorm:"column:modify_time"`
}

type UserStream struct {
}

var userStream *UserStream
var userOnce sync.Once // design pattern of singleton

func NewUserStreamInstance() *UserStream {
	userOnce.Do(func() {
		userStream = &UserStream{}
	})
	return userStream
}

func (*UserStream) GetUserById(id int) (*User, error) {
	var user User
	err := db.Where("id = ?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (*UserStream) GetUsersById(ids []int) (map[int]*User, error) {
	var users []*User
	err := db.Where("id in (?)", ids).Find(&users).Error
	if err != nil {
		return nil, err
	}
	userMap := make(map[int]*User)
	for _, user := range users {
		userMap[user.ID] = user
	}
	return userMap, nil
}
