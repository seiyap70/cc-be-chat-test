package cmd_handler

import (
	"github.com/seiyap70/cc-be-chat-test/pkg/logger"
	"github.com/garyburd/redigo/redis"
	"sync"
)

const (
	popularWordSetKey = "char:popwords"
)

// manage all room
type roomManager struct {
	mutex sync.Mutex
	// all chatroom
	chatRooms map[string]*chatRoom

	// all online user
	allClient map[string]*client

	//
	redisCli redis.Conn
}

var _roomManager roomManager

func InitRoomManager(redisHost string) error {
	_roomManager = roomManager{
		chatRooms: make(map[string]*chatRoom, 10),
		allClient: make(map[string]*client),
	}
	c, err := redis.Dial("tcp", redisHost)
	if err != nil {
		logger.Errorf("connect to redis err %s", err.Error())
		return err
	}
	_roomManager.redisCli = c
	return nil
}

// register client to room and manager online status
func (roomManager *roomManager) joinRoom(c *client, roomName string) {
	roomManager.mutex.Lock()
	defer roomManager.mutex.Unlock()
	if chatRoom, ok := roomManager.chatRooms[roomName]; ok {
		chatRoom.register <- c
	} else {
		chatRoom = newHub()
		chatRoom.register <- c
		roomManager.chatRooms[roomName] = chatRoom
	}
	roomManager.allClient[c.name] = c
}

// clean when client offline
func (roomManager *roomManager) clientOffline(c *client) {
	roomManager.mutex.Lock()
	defer roomManager.mutex.Unlock()
	if c == nil || c.hub == nil {
		return
	}
	delete(roomManager.allClient, c.name)
}

// called when the room is empty
func (roomManager *roomManager) cleanRoom(roomName string) {
	roomManager.mutex.Lock()
	defer roomManager.mutex.Unlock()
	if _, ok := roomManager.chatRooms[roomName]; ok {
		delete(roomManager.chatRooms, roomName)
	}
}

func (roomManager *roomManager) GetClientByName(name string) *client {
	roomManager.mutex.Lock()
	defer roomManager.mutex.Unlock()
	return roomManager.allClient[name]
}

func (roomManager *roomManager) SavePopularWordScore(words []string) {
	if len(words) == 0 {
		return
	}

	argLen := 2*len(words) + 1
	var args = make([]interface{}, argLen, argLen)
	args[0] = popularWordSetKey
	var wordIndex = 0
	for i := 0; i < argLen; i += 2 {
		args[i] = 1
		args[i+1] = words[wordIndex]
		wordIndex++
	}
	roomManager.redisCli.Do("zincrby", args...)
}

func (roomManager *roomManager) popularWord() (string, error) {
	res, err := redis.Values(roomManager.redisCli.Do("zrevrange", []interface{}{popularWordSetKey, 0, 0}))
	if err != nil {
		return "", nil
	}
	if len(res) == 0 {
		return "", nil
	}
	return string(res[0].([]byte)), nil
}
