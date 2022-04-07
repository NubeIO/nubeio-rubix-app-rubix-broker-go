package main

import (
	"flag"
	"github.com/NubeIO/nubeio-rubix-app-rubix-broker-go/pkg/v1/broker"
	"github.com/NubeIO/nubeio-rubix-app-rubix-broker-go/pkg/v1/config"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/types"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"strconv"
)

func main() {

	pathArg := flag.String("config", "", "config file path")
	portArg := flag.String("port", "1883", "mqtt port")
	isProductionArg := flag.String("prod", "true", "set if app is in production mode")
	flag.Parse()
	isProduction, err := strconv.ParseBool(*isProductionArg)
	log.Println("nubeio.broker.go-main() START-APP PATH", *pathArg, "PORT", *portArg, "isProductionArg", isProduction)
	c, err := config.New().IsProduction(isProduction).SetPath(*pathArg).LoadConfig()
	if err != nil {
		log.Errorln(err)
	}
	port := types.ToInt(*portArg)
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
