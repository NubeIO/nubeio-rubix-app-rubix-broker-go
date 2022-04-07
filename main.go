package main

import (
	"github.com/NubeIO/nubeio-rubix-app-rubix-broker-go/pkg/v1/broker"
	"github.com/NubeIO/nubeio-rubix-app-rubix-broker-go/pkg/v1/config"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main() {

	c, err := config.New().SetPath("./config/config.json").LoadConfig()
	if err != nil {
		log.Errorln(err)
	}
	port := 1999
	if c != nil {
		port = c.Listen.Port
	}
	// Create the new MQTT Server.
	err = broker.New().SetPort(port, true).StartBroker()
	if err != nil {
		log.Errorln(err)
		return
	}
	s := keepRunning()
	log.Println("nubeio.broker.go-main() signal to close, broker closed.", s)
}

func keepRunning() os.Signal {
	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)
	signal.Notify(signalChan, os.Kill, os.Interrupt)
	s := <-signalChan
	signal.Stop(signalChan)
	return s
}
