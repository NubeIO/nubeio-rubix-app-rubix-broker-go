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

func (inst *Configuration) LoadConfig() (*Configuration, error) {
	if inst.getPath() == "" { //if path is nil load dev path
		inst.SetPath("")
	}
	file, err := os.Open(inst.getPath())
	if err != nil {
		log.Errorln("nubeio.broker.go-LoadConfig() can't open config file: ", inst.Path)
		log.Errorln("nubeio.broker.go-LoadConfig() can't open config err: ", err)
		return nil, err
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
		log.Errorln("nubeio.broker.go-LoadConfig() can't decode config JSON: ", err)
		return nil, err
	}
	return Config, err

}
