package main

import (
	"github.com/seiyap70/cc-be-chat-test/cmd/chat_server/cmd_handler"
	"github.com/seiyap70/cc-be-chat-test/pkg/logger"
	"flag"
	"fmt"
	"net/http"
)

var addr = flag.String("addr", ":8080", "service listen address")
var logDir = flag.String("log-dir", "logs", "log dir address")
var wsBaseUrl = flag.String("ws-base-url", "chat", "base url for ws")
var redisConnectUrl = flag.String("redis-host", "localhost:6379", "host for redis")

func main() {
	flag.Parse()
	// init logger
	if err := logger.Init(*logDir); err != nil {
		fmt.Printf("init log failed: %s", err)
		return
	}
	if err := cmd_handler.InitRoomManager(*redisConnectUrl); err != nil {
		logger.Fatal("InitRoomManager: ", err)
	}

	http.HandleFunc(*wsBaseUrl, func(w http.ResponseWriter, r *http.Request) {
		cmd_handler.ServeWs(w, r)
	})

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		logger.Fatal("ListenAndServe: ", err)
	}

}
