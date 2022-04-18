package broker

func (inst *Broker) getPort() (port int) {
	return inst.Port
}

func (inst *Broker) getAuth() (auth bool) {
	return inst.Auth
}

func (inst *Broker) getPassword() (auth string) {
	return inst.Password
}
