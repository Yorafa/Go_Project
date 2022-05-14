package entity

import (
	"gorm.io/gorm"
	"sync"
	"time"
)

type Article struct {
	Id         int       `gorm:"column:id"`
	UserId     int       `gorm:"column:user_id"`
	TopicId    int       `gorm:"topic_id"`
	Content    string    `gorm:"column:content"`
	CreateTime time.Time `gorm:"column:create_time"`
}

type ArticleStream struct {
}

var articleStream *ArticleStream
var articleOnce sync.Once // design pattern of singleton

func NewArticleStreamInstance() *ArticleStream {
	articleOnce.Do(func() {
		articleStream = &ArticleStream{}
	})
	return articleStream
}

func (*ArticleStream) GetArticleById(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ?", id).Find(&article).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (*ArticleStream) GetPostByParentId(parentId int) ([]*Article, error) {
	var posts []*Article
	err := db.Where("parent_id = ?", parentId).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (*ArticleStream) CreatePost(article *Article) error {
	if err := db.Create(article).Error; err != nil {
		return err
	}
	return nil
}
