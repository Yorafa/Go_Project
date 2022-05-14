package entity

import (
	"gorm.io/gorm"
	"sync"
	"time"
)

type Topic struct {
	Id         int       `gorm:"column:id"`
	Title      string    `gorm:"column:title"`
	Content    string    `gorm:"column:content"`
	UserId     int       `gorm:"column:user_id"`
	CreateTime time.Time `gorm:"column:create_time"`
}

type TopicStream struct {
}

var topicStream *TopicStream
var topicOnce sync.Once // design pattern of singleton

func NewTopicStreamInstance() *TopicStream {
	topicOnce.Do(func() {
		topicStream = &TopicStream{}
	})
	return topicStream
}

func (*TopicStream) GetTopicById(id int) (*Topic, error) {
	var topic Topic
	err := db.Where("id = ?", id).Find(&topic).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &topic, nil
}
