package cmd_handler

import (
	"github.com/seiyap70/cc-be-chat-test/pkg/constant"
	"github.com/seiyap70/cc-be-chat-test/pkg/dto"
	"fmt"
	"time"
)

type joinRoomHandler struct {
}

func (handler joinRoomHandler) GetName() string {
	return constant.CmdJoinRoom
}

func (handler joinRoomHandler) Handle(c *client, request *requestMsg) {
	joinReq, ok := request.Data.(*joinRoom)
	if !ok {
		_ = c.WriteResponseData(request.Id, request.Cmd, dto.ErrorCodeInvalidParam, nil)
		return
	}
	if c.hub != nil {
		_ = c.WriteResponseData(request.Id, request.Cmd, dto.ErrorCodeInvalidStatus, nil)
		return
	}
	c.name = joinReq.UserName
	c.registerAt = time.Now()
	_roomManager.joinRoom(c, joinReq.RoomName)
	response := joinRoomResponse{}
	_ = c.WriteResponseData(request.Id, request.Cmd, dto.ErrorCodeOK, response)
}

type leftRoomHandler struct {
}

func (handler leftRoomHandler) GetName() string {
	return constant.CmdLeftRoom
}

func (handler leftRoomHandler) Handle(c *client, request *requestMsg) {
	_, ok := request.Data.(*leftRoom)
	if !ok {
		_ = c.WriteResponseData(request.Id, request.Cmd, dto.ErrorCodeInvalidParam, nil)
		return
	}
	if c.hub == nil {
		_ = c.WriteResponseData(request.Id, request.Cmd, dto.ErrorCodeInvalidStatus, nil)
		return
	}
	c.hub.unregister <- c
	response := leftRoomResponse{}
	_ = c.WriteResponseData(request.Id, request.Cmd, dto.ErrorCodeOK, response)

}

type sendChatHandler struct {
}

func (handler sendChatHandler) GetName() string {
	return constant.CmdSendChat
}

func (handler sendChatHandler) Handle(c *client, request *requestMsg) {
	sendMsg, ok := request.Data.(*sendChatMessage)
	if !ok {
		_ = c.WriteResponseData(request.Id, request.Cmd, dto.ErrorCodeInvalidParam, nil)
		return
	}
	if c.hub == nil {
		_ = c.WriteResponseData(request.Id, request.Cmd, dto.ErrorCodeInvalidStatus, nil)
		return
	}
	replacedMsg := _profanityWords.asteriskWords(sendMsg.text)

	c.hub.broadcast <- chatRoomMessage{
		origin:       sendMsg.text,
		replacedText: replacedMsg,
		from:         c.name,
	}
}

type popularWordHandler struct {
}

func (handler popularWordHandler) GetName() string {
	return constant.CmdPopularWord
}

func (handler popularWordHandler) Handle(c *client, request *requestMsg) {
	_, ok := request.Data.(*popularWord)
	if !ok {
		_ = c.WriteResponseData(request.Id, request.Cmd, dto.ErrorCodeInvalidParam, nil)
		return
	}
	if c.hub == nil {
		_ = c.WriteResponseData(request.Id, request.Cmd, dto.ErrorCodeInvalidStatus, nil)
		return
	}
	word, err := _roomManager.popularWord()
	if err != nil {
		_ = c.WriteResponseData(request.Id, request.Cmd, dto.ErrorCodeServerError, nil)
		return
	}
	response := popularWordResponse{word: word}
	_ = c.WriteResponseData(request.Id, request.Cmd, dto.ErrorCodeOK, response)
}

type statsUserHandler struct {
}

func (handler statsUserHandler) GetName() string {
	return constant.CmdStatsUser
}

func formatStatsDuration(duration time.Duration) string {
	diffSec := int32(duration.Seconds())
	daySecs := int32(24 * 60 * 60)
	hourSecs := int32(60 * 60)
	day := diffSec / (24 * 60 * 60)
	hour := (diffSec - day*daySecs) / hourSecs
	min := (diffSec - day*daySecs - hour*hourSecs) / 60
	sec := diffSec % 60

	statsStr := fmt.Sprintf("%02dd %02dh %02dm %02ds", day, hour, min, sec)
	return statsStr
}
func (handler statsUserHandler) Handle(c *client, request *requestMsg) {
	statsData, ok := request.Data.(*statsUser)
	if !ok {
		_ = c.WriteResponseData(request.Id, request.Cmd, dto.ErrorCodeInvalidParam, nil)
		return
	}
	if c.hub == nil {
		_ = c.WriteResponseData(request.Id, request.Cmd, dto.ErrorCodeInvalidStatus, nil)
		return
	}
	client := _roomManager.GetClientByName(statsData.UserName)
	if client == nil {
		_ = c.WriteResponseData(request.Id, request.Cmd, dto.ErrorCodeClientNotFound, nil)
	}
	diff := time.Now().Sub(client.registerAt)
	statsStr := formatStatsDuration(diff)
	stats := statsResponse{stats: statsStr}

	_ = c.WriteResponseData(request.Id, request.Cmd, dto.ErrorCodeOK, stats)
}
