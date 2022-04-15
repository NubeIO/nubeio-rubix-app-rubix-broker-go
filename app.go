package main

import (
	"flag"
	"github.com/NubeIO/nubeio-rubix-app-rubix-broker-go/logger"
	"github.com/NubeIO/nubeio-rubix-app-rubix-broker-go/pkg/v1/broker"
	"github.com/NubeIO/nubeio-rubix-app-rubix-broker-go/pkg/v1/config"
	"os"
	"os/signal"
	"path"
)

var (
	log = logger.New()
)

func main() {
	globalDir := flag.String("g", "./", "Global Directory")
	configDir := flag.String("c", "config", "Config Directory")
	port := flag.Int("p", 1883, "MQTT TLS port")
	prod := flag.Bool("prod", false, "Deployment Mode")
	flag.Parse()
	configPath := path.Join(*globalDir, *configDir, "config.json")
	c := config.New().IsProduction(*prod).SetPath(configPath).LoadConfig()
	if c != nil {
		port = &c.Listen.Port
	}
	log.Info("starting app with configPath: ", configPath, ", port: ", *port, ", prod: ", *prod)
	// Create the new MQTT Server.
	err := broker.New().SetPort(*port).StartBroker()
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
