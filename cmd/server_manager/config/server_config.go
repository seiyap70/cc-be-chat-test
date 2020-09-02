package config

// config for each server
type ChatServerConfig struct {
	AccessUrl string `json:"access_url"`
}

// config
type ServerConfig struct {
	LogDir string `json:"logDir"`
	ListenPort	string `json:"listenPort"`
	Servers []ChatServerConfig `json:"servers"`
}
