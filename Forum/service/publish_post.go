package service

import (
	"Go_Project/Forum/entity"
	"errors"
	"time"
)

type PostStream struct {
	userId  int
	content string

	topicId int
	postId  int
}

func PublishPost(topicId int, userId int, content string) (int, error) {
	return NewPublishPostStream(topicId, userId, content).Init()
}

func NewPublishPostStream(topicId, userId int, content string) *PostStream {
	return &PostStream{userId: userId, content: content, topicId: topicId}
}

func (p *PostStream) isValid() error {
	if p.userId <= 0 {
		return errors.New("user Id is invalid")
	}
	if p.topicId <= 0 {
		return errors.New("topic Id is invalid")
	}
	return nil
}

func (p *PostStream) post() error {
	article := &entity.Article{
		TopicId:    p.topicId,
		UserId:     p.userId,
		Content:    p.content,
		CreateTime: time.Now(),
	}
	if err := entity.NewArticleStreamInstance().CreatePost(article); err != nil {
		return err
	}
	p.postId = article.Id
	return nil
}

func (p *PostStream) Init() (int, error) {
	if err := p.isValid(); err != nil {
		return 0, err
	}
	if err := p.post(); err != nil {
		return 0, err
	}
	return p.postId, nil
}
