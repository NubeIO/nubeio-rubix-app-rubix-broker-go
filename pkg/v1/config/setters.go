package config

import "fmt"

func (inst *Configuration) SetPath(path string) *Configuration {
	if path == "" {
		if inst.Production { //if path is nil and in production then use this path
			path = fmt.Sprintf("data/rubix-broker/config/config.json")
		} else { // else use dev path
			path = fmt.Sprintf("./config/config.json")
		} // or use user path
	}
	inst.Path = path
	return inst
}
