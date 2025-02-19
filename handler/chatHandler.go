package handlers

import (
	"chatApp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JoinHandler(ctx *gin.Context, chatRoom *service.ChatRoom) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
		return
	}

	_, err := chatRoom.Join(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Client joined"})
}

func LeaveHandler(ctx *gin.Context, chatRoom *service.ChatRoom) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
		return
	}

	err := chatRoom.Leave(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Client left"})
}

func SendMessageHandler(ctx *gin.Context, chatRoom *service.ChatRoom) {
	id := ctx.Query("id")
	msg := ctx.Query("message")
	if id == "" || msg == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing id or message"})
		return
	}

	err := chatRoom.SendMessage(id, msg)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Message sent"})
}

func GetMessagesHandler(ctx *gin.Context, chatRoom *service.ChatRoom) {
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing id"})
		return
	}

	msg, err := chatRoom.GetMessages(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": msg})
}
