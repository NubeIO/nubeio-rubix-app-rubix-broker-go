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

func (inst *Broker) getPersistence() bool {
	return inst.EnablePersistence
}
