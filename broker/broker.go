package broker

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-app-rubix-broker-go/config"
	"github.com/NubeIO/nubeio-rubix-app-rubix-broker-go/logger"
	mqtt "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/listeners"
	"github.com/mochi-co/mqtt/server/persistence/bolt"
	"go.etcd.io/bbolt"
	"path"
	"time"
)

var (
	log = logger.New()
)

func StartBroker(conf *config.Configuration) error {
	configPath := path.Join(conf.GetAbsConfigDir(), "config.yml")
	dataPath := path.Join(conf.GetAbsDataDir(), conf.Storage.DB)
	log.Info("starting app with config_path: ", configPath, ", data_path: ", dataPath,
		", port: ", conf.Server.Port, ", prod: ", conf.Prod, ", auth: ", conf.Credential.Auth,
		", enable_persistence: ", *conf.Storage.EnablePersistence)
	server := mqtt.New()
	var err error
	if *conf.Storage.EnablePersistence {
		err = server.AddStore(bolt.New(dataPath, &bbolt.Options{
			Timeout: 500 * time.Millisecond,
		}))
	}
	port := fmt.Sprintf(":%d", conf.Server.Port)
	// Create a TCP listener on a standard port.
	tcp := listeners.NewTCP("t1", port)
	// Add the listener to the server with default options (nil).
	if conf.Credential.Auth {
		username := conf.Credential.Username
		password := conf.Credential.Password
		err = server.AddListener(tcp, &listeners.Config{
			Auth: &Auth{
				Users: map[string]string{
					username: password,
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
