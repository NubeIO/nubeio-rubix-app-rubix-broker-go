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
	dataDir := flag.String("d", "data", "Data Directory")
	configDir := flag.String("c", "config", "Config Directory")
	port := flag.Int("p", 1883, "MQTT TLS port")
	prod := flag.Bool("prod", false, "Deployment Mode")
	auth := flag.Bool("auth", false, "enable mqtt auth")
	password := flag.String("password", "N00BMQTT", "auth password")
	enablePersistence := true
	db := "mqtt.db"
	flag.Parse()
	absConfigDir := path.Join(*globalDir, *configDir)
	configPath := path.Join(absConfigDir, "config.json")
	c := config.New().IsProduction(*prod).SetPath(configPath).LoadConfig()
	if c != nil {
		port = &c.Listen.Port
		auth = &c.Listen.UseAuth
		password = &c.Listen.Password
		enablePersistence = c.Listen.EnablePersistence
		db = c.Listen.DB
	}
	absDataDir := path.Join(*globalDir, *dataDir)
	dataPath := path.Join(absDataDir, db)
	if err := os.MkdirAll(absConfigDir, 0755); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(absDataDir, 0755); err != nil {
		panic(err)
	}
	log.Info("starting app with config_path: ", configPath, ", data_path: ", dataPath, ", port: ", *port, ", prod: ", *prod, ", auth: ", auth, ", enable_persistence: ", enablePersistence)
	// Create the new MQTT Server
	err := broker.New().SetPort(*port).SetAuth(*auth).SetPassword(*password).SetPersistence(enablePersistence).SetAbsoluteDbPath(dataPath).StartBroker()
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
