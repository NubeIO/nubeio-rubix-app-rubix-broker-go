package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
)

type Configuration struct {
	Path       string
	Production bool
	Listen     struct {
		Address string
		Port    int
	}
}

func New() *Configuration {
	c := &Configuration{}
	return c
}

func (inst *Configuration) LoadConfig() *Configuration {
	file, err := os.Open(inst.getPath())
	if err != nil {
		log.Warn("running without config file")
		return nil
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
		}
	}(file)
	decoder := json.NewDecoder(file)
	Config := &Configuration{}
	err = decoder.Decode(Config)
	if err != nil {
		log.Error("can't decode config JSON: ", err)
		return nil
	}
	return Config
}
