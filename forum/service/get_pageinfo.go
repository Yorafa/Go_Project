package service

import (
	"Go_Project/forum/entity"
	"errors"
	"fmt"
)

type PostInfo struct {
	Post *entity.Article
	User *entity.User
}

type TopicInfo struct {
	Tag  *entity.Topic
	User *entity.User
}

type PageInfo struct {
	Tag      *TopicInfo
	Articles []*PostInfo
}

type PageStream struct {
	topicId  int
	pageInfo *PageInfo

	topic   *entity.Topic
	posts   []*entity.Article
	userMap map[int]*entity.User
}

func GetPageInfo(id int) (*PageInfo, error) {
	return NewQueryPageInfo(id).Init()
}

func NewQueryPageInfo(id int) *PageStream {
	return &PageStream{
		topicId: id,
	}
}

func (p *PageStream) isValid() error {
	if p.topicId <= 0 {
		return errors.New("topic Id is invalid")
	}
	return nil
}

func (p *PageStream) findTopic() error {
	topic, err := entity.NewTopicStreamInstance().GetTopicById(p.topicId)
	if err != nil {
		return err
	}
	p.topic = topic
	return nil
}

func (p *PageStream) findPost() error {
	posts, err := entity.NewArticleStreamInstance().GetPostByParentId(p.topicId)
	if err != nil {
		return err
	}
	p.posts = posts
	return nil
}

func (p *PageStream) findUser() error {
	var users []int
	for _, post := range p.posts {
		users = append(users, post.UserId)
	}
	userMap, err := entity.NewUserStreamInstance().GetUsersById(users)
	if err != nil {
		return err
	}
	p.userMap = userMap
	return nil
}

func (p *PageStream) packPageInfo() error {
	userMap := p.userMap
	topicUser, ok := userMap[p.topic.UserId]
	if !ok {
		return errors.New("has no topic user info")
	}
	var posts []*PostInfo
	for _, post := range p.posts {
		postUser, ok := userMap[post.UserId]
		if !ok {
			return errors.New("has no post user info" + fmt.Sprint(post.UserId))
		}
		posts = append(posts, &PostInfo{post, postUser})
	}
	p.pageInfo = &PageInfo{&TopicInfo{p.topic, topicUser}, posts}
	return nil
}

func (p *PageStream) Init() (*PageInfo, error) {
	if err := p.isValid(); err != nil {
		return nil, err
	}
	if err := p.findTopic(); err != nil {
		return nil, err
	}
	if err := p.findPost(); err != nil {
		return nil, err
	}
	if err := p.findUser(); err != nil {
		return nil, err
	}
	if err := p.packPageInfo(); err != nil {
		return nil, err
	}
	return p.pageInfo, nil
}
