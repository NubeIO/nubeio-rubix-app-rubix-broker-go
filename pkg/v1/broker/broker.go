package broker

import (
	"fmt"
	mqtt "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/listeners"
)

type Broker struct {
	Port int `json:"port"`
}

// New returns a new instance of the nube common apis
func New() *Broker {
	bc := &Broker{}
	return bc
}

func (inst *Broker) StartBroker() error {
	server := mqtt.New()
	port := fmt.Sprintf(":%d", inst.getPort())
	// Create a TCP listener on a standard port.
	tcp := listeners.NewTCP("t1", port)
	// Add the listener to the server with default options (nil).
	err := server.AddListener(tcp, nil)
	if err != nil {
		return err
	}
	// Start the broker. Serve() is blocking
	err = server.Serve()
	if err != nil {
		return err
	}
	return err
}
