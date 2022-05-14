package handler

import (
	"Go_Project/Forum/service"
	"strconv"
)

type PageData struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func GetPageInfo(topicIdStr string) *PageData {
	topicId, err := strconv.Atoi(topicIdStr)
	if err != nil {
		return &PageData{-1, err.Error(), nil}
	}
	pageInfo, err := service.GetPageInfo(topicId)
	if err != nil {
		return &PageData{-1, err.Error(), nil}
	}
	return &PageData{0, "success", pageInfo}
}
