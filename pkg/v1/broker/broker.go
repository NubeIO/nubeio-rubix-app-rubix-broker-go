package broker

import (
	"fmt"
	mqtt "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/listeners"
)

type Broker struct {
	Port     int    `json:"port"`
	Auth     bool   `json:"auth"`
	Password string `json:"password"`
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
	var err error
	if inst.getAuth() {
		pass := inst.getPassword()
		err = server.AddListener(tcp, &listeners.Config{
			Auth: &Auth{
				Users: map[string]string{
					"admin": pass,
				},
			},
		})
	} else {
		err = server.AddListener(tcp, nil)
	}
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

type Auth struct {
	Users         map[string]string   // A map of usernames (key) with passwords (value).
	AllowedTopics map[string][]string // A map of usernames and topics
}

// Authenticate returns true if a username and password are acceptable.
func (a *Auth) Authenticate(user, password []byte) bool {
	// If the user exists in the auth users map, and the password is correct,
	// then they can connect to the server. In the real world, this could be a database
	// or cached users lookup.
	if pass, ok := a.Users[string(user)]; ok && pass == string(password) {
		return true
	}
	return false
}

// ACL returns true if a user has access permissions to read or write on a topic.
func (a *Auth) ACL(user []byte, topic string, write bool) bool {
	// An example ACL - if the user has an entry in the auth allow list, then they are
	// subject to ACL restrictions. Only let them use a topic if it's available for their
	// user.
	if topics, ok := a.AllowedTopics[string(user)]; ok {
		for _, t := range topics {
			// In the real world you might allow all topics prefixed with a user's username,
			// or similar multi-topic filters.
			if t == topic {
				return true
			}
		}
		return false
	}
	// Otherwise, allow all topics.
	return true
}
