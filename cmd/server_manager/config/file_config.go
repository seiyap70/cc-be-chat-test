package config

import (
	"encoding/json"
	"io/ioutil"
)

type FileConfig struct {
	path string
}


func (fileConfig *FileConfig) Load(v interface{}) error {
	// load file content
	data, err := ioutil.ReadFile(fileConfig.path)
	if err != nil {
		return err
	}

	// json unmarshal
	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}