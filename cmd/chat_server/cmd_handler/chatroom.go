package cmd_handler

import (
	"github.com/seiyap70/cc-be-chat-test/pkg/constant"
	"github.com/seiyap70/cc-be-chat-test/pkg/dto"
	"strings"
)

const maxSavedBroadMessageCount = 50

type chatRoomMessage struct {
	origin       string
	replacedText string
	from         string
}
type chatRoom struct {
	// Registered clients.
	clients map[*client]bool

	// Inbound messages from the clients.
	broadcast chan chatRoomMessage

	// Register requests from the clients.
	register chan *client

	// Unregister requests from clients.
	unregister chan *client

	//
	lastChatMessage []chatRoomMessage
}

func newHub() *chatRoom {
	return &chatRoom{
		broadcast:       make(chan chatRoomMessage),
		register:        make(chan *client),
		unregister:      make(chan *client),
		clients:         make(map[*client]bool),
		lastChatMessage: []chatRoomMessage{},
	}
}

func (h *chatRoom) saveBroadcastMessage(message chatRoomMessage) {
	h.lastChatMessage = append(h.lastChatMessage, message)
	if len(h.lastChatMessage) > maxSavedBroadMessageCount {
		copy(h.lastChatMessage, h.lastChatMessage[len(h.lastChatMessage)-maxSavedBroadMessageCount:])
	}
	words := strings.Fields(message.replacedText)
	_roomManager.SavePopularWordScore(words)
}

func (h *chatRoom) sendLastMessage(client *client) {
	msgCnt := len(h.lastChatMessage)
	if msgCnt > maxSavedBroadMessageCount {
		msgCnt = maxSavedBroadMessageCount
	}
	sendMsg := broadcastMessage{
		messages: make([]textMessage, msgCnt, msgCnt),
	}

	for i := 0; i < msgCnt; i++ {
		sendMsg.messages[i] = textMessage{
			from: h.lastChatMessage[i].from,
			text: h.lastChatMessage[i].replacedText,
		}
	}
	_ = client.WriteResponseData(dto.RequestIdFromServer,
		constant.CmdBroadcastMessage,
		dto.ErrorCodeOK,
		sendMsg,
	)
}
func (h *chatRoom) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			h.sendLastMessage(client)
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
			}
		case message := <-h.broadcast:
			// cal word statistic data
			h.saveBroadcastMessage(message)
			// broad cast
			sendMsg := broadcastMessage{
				messages: []textMessage{{from: message.from, text: message.replacedText}},
			}
			for client := range h.clients {
				_ = client.WriteResponseData(dto.RequestIdFromServer,
					constant.CmdBroadcastMessage,
					dto.ErrorCodeOK,
					sendMsg,
				)
			}
		}
	}
}
