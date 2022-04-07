package broker

import (
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/networking/freeport"
	log "github.com/sirupsen/logrus"
)

/*
SetPort
	-port mqtt port
	-bumpPort if true and selected port is use it will bump the port number (port += port)
*/
func (inst *Broker) SetPort(port int, bumpPort bool) *Broker {
	if port == 0 {
		inst.Port = 1888
	} else {
		inst.Port = port
	}
	if bumpPort {
		port, err := freeport.FindFreePort(inst.Port)
		if err != nil {
			log.Errorln(err)
		} else {
			inst.Port = port
		}
	}
	return inst
}
