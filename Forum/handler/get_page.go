package handler

import (
	"Go_Project/Forum/service"
	"strconv"
)

func PublishPost(userIdStr, topicIdStr, content string) *PageData {
	userId, _ := strconv.Atoi(userIdStr)
	topicId, _ := strconv.Atoi(topicIdStr)
	postId, err := service.PublishPost(topicId, userId, content)
	if err != nil {
		return &PageData{
			Status: -1,
			Msg:    err.Error(),
			Data:   nil,
		}
	}
	return &PageData{
		Status: 0,
		Msg:    "success",
		Data:   map[string]int{"post_id": postId},
	}
}
