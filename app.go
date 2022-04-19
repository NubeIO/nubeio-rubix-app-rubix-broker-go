package main

import (
	"github.com/NubeIO/nubeio-rubix-app-rubix-broker-go/broker"
	"github.com/NubeIO/nubeio-rubix-app-rubix-broker-go/config"
	"github.com/NubeIO/nubeio-rubix-app-rubix-broker-go/logger"
	"os"
	"os/signal"
)

var (
	log = logger.New()
)

func main() {
	conf := config.CreateApp()
	if err := os.MkdirAll(conf.GetAbsConfigDir(), 0755); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(conf.GetAbsDataDir(), 0755); err != nil {
		panic(err)
	}

	// Create the new MQTT Server
	err := broker.StartBroker(conf)
	if err != nil {
		log.Error(err)
		return
	}
	s := keepRunning()
	log.Info("signal to close, broker closed by: ", s)
}

func keepRunning() os.Signal {
	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)
	signal.Notify(signalChan, os.Kill, os.Interrupt)
	s := <-signalChan
	signal.Stop(signalChan)
	return s
}
