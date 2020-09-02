package controller

import (
	"github.com/seiyap70/cc-be-chat-test/cmd/server_manager/config"
	"github.com/seiyap70/cc-be-chat-test/pkg/dto"
	"github.com/seiyap70/cc-be-chat-test/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

// query get chat server by chat room name
func QueryChatServerInfo(context *gin.Context) {
	chatroom := context.Param("chatroom")
	accessUrl, err := config.GetChatServer(chatroom)
	if err != nil {
		logger.Warn("get chat server failed: ", err)
		context.JSON(http.StatusOK, dto.Response{
			Code:    dto.ErrorFoundChatServer.Code,
			Message: err.Error(),
		})
	}
	//
	logger.Infof("get chat server %s for chatroom %s", accessUrl, chatroom)
	context.JSON(http.StatusOK, dto.Response{
		Code:    dto.ErrorOK.Code,
		Message: dto.ErrorOK.Message,
		Result: dto.ChatServerInfo{
			AccessUrl: accessUrl,
		},
	})

}
