package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/seiyap70/cc-be-chat-test/cmd/server_manager/config"
	"github.com/seiyap70/cc-be-chat-test/cmd/server_manager/controller"
	"github.com/seiyap70/cc-be-chat-test/pkg/logger"
)

var configPath = flag.String("configPath", "./config.json", "the config file path")

func main() {
	if err := config.InitConfig(*configPath); err != nil {
		fmt.Printf("load config failed: %s", err)
		return
	}
	if err := logger.Init(config.Configuration.LogDir); err != nil {
		fmt.Printf("init log failed: %s", err)
		return
	}

	router := gin.Default()

	router.GET("/server/:chatroom", controller.QueryChatServerInfo)

	// listen client request
	router.Run(config.Configuration.ListenPort)
}
