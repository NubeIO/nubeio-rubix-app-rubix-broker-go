# nubeio-rubix-app-rubix-broker-go


## running

run with config
```
go run app.go -g ./
```

run without config
```
go run app.go -port 1882
```

run with auth
```
go run app.go -auth=true -password admin
```


# default port
1883


## Libs used
Small MQTT broker in go

uses https://github.com/mochi-co/mqtt


also tested and found works 

https://github.com/fhmq/hmq

example of hmq

```go
func (i *Instance) enableBroker() {
	port := "1883"
	if i.config.Port != "" {
		port = i.config.Port
	}
	HttpPort := "8099"
	if i.config.HttpPort != "" {
		HttpPort = i.config.HttpPort
	}
	os.Args = []string{"-port", port, "-httpport", HttpPort}
	config, err := broker.ConfigureConfig(os.Args)
	if err != nil {
		log.Error("configure broker config error: ", err)
	}
	b, err := broker.NewBroker(config)
	if err != nil {
		log.Error("New Broker error: ", err)
	}
	b.Start()
	s := waitForSignal()
	log.Println("signal received, broker closed.", s)
}

func waitForSignal() os.Signal {
	signalChan := make(chan os.Signal, 1)
	defer close(signalChan)
	signal.Notify(signalChan, os.Kill, os.Interrupt)
	s := <-signalChan
	signal.Stop(signalChan)
	return s
}

```
