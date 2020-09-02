package cmd_handler

type Command interface {
	GetName() string
}

type CommandHandler interface {
	Command
	Handle(c *client, request *requestMsg)
}

type CommandHandlerManager struct {
	CommandHandlers []CommandHandler
}

func DispatchCmdRequest(c *client, request *requestMsg) {
	if handler, ok := _cmdTypeMap[request.Cmd]; ok {
		handler.cmdHandler.Handle(c, request)
	}
	//return dto.ErrorNotFoundCmd
}
