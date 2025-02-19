package main

import (
	"log"

	handler "chatApp/handler"
	"chatApp/service"

	"github.com/gin-gonic/gin"
)

func main() {
	chatRoom := service.NewChatRoom()
	go chatRoom.BroadcastMessages()
	router := gin.Default()

	router.GET("/chatApp/join", func(ctx *gin.Context) { handler.JoinHandler(ctx, chatRoom) })
	router.GET("/chatApp/leave", func(ctx *gin.Context) { handler.LeaveHandler(ctx, chatRoom) })
	router.GET("/chatApp/send", func(ctx *gin.Context) { handler.SendMessageHandler(ctx, chatRoom) })
	router.GET("/chatApp/getMessage", func(ctx *gin.Context) { handler.GetMessagesHandler(ctx, chatRoom) })
	log.Println("Chat server started on :8080")
	router.Run(":8080")
}
