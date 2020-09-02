package cmd_handler

import (
	"github.com/seiyap70/cc-be-chat-test/pkg/constant"
	"reflect"
)

type cmdRegister struct {
	cmd         string
	requestType reflect.Type
	cmdHandler  CommandHandler
}

var _cmdTypeMap map[string]cmdRegister

func init() {
	_cmdTypeMap = make(map[string]cmdRegister, 10)
	_cmdTypeMap[constant.CmdJoinRoom] = cmdRegister{
		cmd:         constant.CmdJoinRoom,
		requestType: reflect.TypeOf(joinRoom{}),
		cmdHandler:  joinRoomHandler{},
	}
	_cmdTypeMap[constant.CmdLeftRoom] = cmdRegister{
		cmd:         constant.CmdLeftRoom,
		requestType: reflect.TypeOf(leftRoom{}),
		cmdHandler:  leftRoomHandler{},
	}
	_cmdTypeMap[constant.CmdSendChat] = cmdRegister{
		cmd:         constant.CmdSendChat,
		requestType: reflect.TypeOf(leftRoom{}),
		cmdHandler:  sendChatHandler{},
	}
	_cmdTypeMap[constant.CmdPopularWord] = cmdRegister{
		cmd:         constant.CmdPopularWord,
		requestType: reflect.TypeOf(leftRoom{}),
		cmdHandler:  popularWordHandler{},
	}
	_cmdTypeMap[constant.CmdStatsUser] = cmdRegister{
		cmd:         constant.CmdStatsUser,
		requestType: reflect.TypeOf(leftRoom{}),
		cmdHandler:  statsUserHandler{},
	}
}
