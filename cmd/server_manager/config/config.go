package config

import (
	"errors"
	consistent_hash "github.com/lithammer/go-jump-consistent-hash"
)

type Loader interface {
	Load(v interface{}) error
}

func newFileConfig(path string) *FileConfig {
	return &FileConfig{path: path}
}

var Configuration ServerConfig

func InitConfig(configPath string) error {
	var config ServerConfig
	var configLoader Loader = newFileConfig(configPath)
	if err := configLoader.Load(&config); err != nil {
		return err
	}
	// sever list cannot be empty
	if len(config.Servers) == 0 {
		return errors.New("chat server list empty")
	}
	Configuration = config
	return nil
}

func GetChatServer(chatRoom string) (string, error) {
	// cal the hash of chat room name, index the server for client connect
	h := consistent_hash.HashString(chatRoom, int32(len(Configuration.Servers)), consistent_hash.NewCRC64())
	return Configuration.Servers[h].AccessUrl, nil
}
