package main

import (
	"Go_Project/Forum/entity"
	"Go_Project/Forum/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, `usage: server username password`)
		os.Exit(-1)
	}
	if err := entity.Init(os.Args[1], os.Args[2]); err != nil {
		os.Exit(-1)
	}
	// setup gin
	r := gin.Default()
	// enable logger
	r.Use(gin.Logger())

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "success"})
	})
	r.GET("/community/page/get/:id", func(context *gin.Context) {
		topicId := context.Param("id")
		data := handler.GetPageInfo(topicId)
		context.JSON(200, data)
	})
	r.POST("/community/post/do", func(context *gin.Context) {
		userId, _ := context.GetPostForm("user_id")
		topicId, _ := context.GetPostForm("topic_id")
		content, _ := context.GetPostForm("content")
		data := handler.PublishPost(userId, topicId, content)
		context.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		return
	}
}
