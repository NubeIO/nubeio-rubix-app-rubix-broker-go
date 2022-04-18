package config

import (
	"encoding/json"
	"github.com/NubeIO/nubeio-rubix-app-rubix-broker-go/logger"
	"os"
)

var (
	log = logger.New()
)

type Configuration struct {
	Path       string
	Production bool
	Listen     struct {
		Address  string
		Port     int
		UseAuth  bool
		Password string
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
