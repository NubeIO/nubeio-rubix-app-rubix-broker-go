package broker

func (inst *Broker) SetPort(port int) *Broker {
	inst.Port = port
	return inst
}

func (inst *Broker) SetAuth(auth bool) *Broker {
	inst.Auth = auth
	return inst
}

func (inst *Broker) SetPassword(auth string) *Broker {
	inst.Password = auth
	return inst
}
