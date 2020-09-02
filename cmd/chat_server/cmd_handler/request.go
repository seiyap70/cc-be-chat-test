package cmd_handler

import (
	"github.com/seiyap70/cc-be-chat-test/pkg/dto"
	"encoding/json"
	"reflect"
)

type requestParseType struct {
	Cmd  string          `json:"cmd"`
	Data json.RawMessage `json:"data"`
	Id   string          `json:"id"`
}

type requestMsg struct {
	Cmd  string      `json:"cmd"`
	Data interface{} `json:"data"`
	Id   string      `json:"id"`
}

type response struct {
	Cmd  string      `json:"cmd"`
	Data interface{} `json:"data"`
	Id   string      `json:"id"`
	Code int         `json:"code"`
}

type joinRoom struct {
	UserName string `json:"userName"`
	RoomName string `json:"roomName"`
}

type joinRoomResponse struct {
}

type leftRoom struct {
}
type leftRoomResponse struct {
}

type sendChatMessage struct {
	text string `json:"text"`
}

type textMessage struct {
	text string `json:"text"`
	from string `json:"from"`
}
type broadcastMessage struct {
	messages []textMessage
}

type popularWord struct {
}
type popularWordResponse struct {
	word string
}
type statsUser struct {
	UserName string `json:"userName"`
}
type statsResponse struct {
	stats string `json:"stats"`
}

func ParseRequest(msg []byte) (*requestMsg, error) {
	var parse = new(requestParseType)
	if err := json.Unmarshal(msg, &parse); err != nil {
		return nil, err
	}
	var cmdReg cmdRegister
	var ok bool
	if cmdReg, ok = _cmdTypeMap[parse.Cmd]; !ok {
		return nil, dto.ErrorNotFoundCmd
	}
	val := reflect.New(cmdReg.requestType).Interface()
	if err := json.Unmarshal(parse.Data, &val); err != nil {
		return nil, err
	}
	var retRequest = &requestMsg{
		Cmd:  parse.Cmd,
		Id:   parse.Id,
		Data: val,
	}
	return retRequest, nil
}
